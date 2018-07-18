package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/dave/jennifer/jen"
)

var typeMap = map[string]string{
	"String":                     "string",
	"Accelerator":                "string",
	"Boolean":                    "bool",
	"Number":                     "int",
	"Buffer":                     "[]byte",
	"Function":                   "qrpc.ObjectHandle",
	"Integer":                    "int",
	"Double":                     "float64",
	"Menu":                       "qrpc.ObjectHandle",
	"MenuItem":                   "qrpc.ObjectHandle",
	"NativeImage":                "qrpc.ObjectHandle",
	"Tray":                       "qrpc.ObjectHandle",
	"Notification":               "qrpc.ObjectHandle",
	"MenuItemConstructorOptions": "*MenuItemConstructorOptions",
}

func convertType(typ interface{}, collection bool) string {
	pre := ""
	if collection {
		pre = "[]"
	}
	switch t := typ.(type) {
	case string:
		tt, ok := typeMap[t]
		if !ok {
			return pre + t
		}
		return pre + tt
	case []interface{}:
		first := t[0].(map[string]interface{})
		return convertType(first["typeName"].(string), collection)
	default:
		fmt.Println("unknown type:", typ)
		return ""
	}
}

func propTag(name string, required bool) map[string]string {
	if !required {
		return map[string]string{"msgpack": fmt.Sprintf("%s,omitempty", name)}
	}
	return map[string]string{"msgpack": name}
}

func main() {
	b, err := ioutil.ReadFile("schema/schema.json")
	if err != nil {
		panic(err)
	}
	header, err := ioutil.ReadFile("schema/lib/clientheader.go")
	if err != nil {
		panic(err)
	}
	var schema Schema
	err = json.Unmarshal(b, &schema)
	if err != nil {
		panic(err)
	}
	submodules := make(map[string]string)
	for _, mod := range schema.Modules {
		if strings.Contains(mod.Path, ".") {
			submodules[mod.Path] = mod.Name
		}
	}
	f := jen.NewFile("electrond")
	ptr := jen.Op("*")

	// setup client
	fields := []jen.Code{jen.Add(ptr).Id("qrpc.Client")}
	var setup []jen.Code
	var moduleFields []string
	for _, mod := range schema.Modules {
		var fieldName string
		if _, ok := submodules[mod.Path]; ok {
			fieldName = strings.Replace(strings.Title(strings.Replace(mod.Path, ".", " ", -1)), " ", "", -1)
		} else {
			fieldName = strings.Title(mod.Name)
		}
		fieldType := strings.Title(mod.Name) + "Module"
		fields = append(fields, jen.Id(fieldName).Add(ptr).Id(fieldType))
		setup = append(setup, jen.Id("c").Dot(fieldName).Op("=").Add(jen.Op("&")).Id(fieldType).Values(jen.Dict{
			jen.Id("Client"): jen.Id("c").Dot("Client"),
		}))
		moduleFields = append(moduleFields, fieldName)
	}
	for _, cls := range schema.Classes {
		fieldName := strings.Title(cls.Name)
		skip := false
		for _, name := range moduleFields {
			if name == fieldName {
				skip = true
				break
			}
		}
		if skip {
			continue
		}
		fieldType := fieldName + "Module"
		fields = append(fields, jen.Id(fieldName).Add(ptr).Id(fieldType))
		setup = append(setup, jen.Id("c").Dot(fieldName).Op("=").Add(jen.Op("&")).Id(fieldType).Values(jen.Dict{
			jen.Id("Client"): jen.Id("c").Dot("Client"),
		}))
	}
	f.Type().Id("Client").Struct(fields...)
	f.Func().Params(jen.Id("c").Id("*Client")).Id("setup").Params().Id("*Client").Block(append(setup, jen.Return(jen.Id("c")))...)

	//f.Type().Id("NativeImage").Struct()

	for _, str := range schema.Structs {
		var fields []jen.Code
		for _, prop := range str.Properties {
			fields = append(fields, jen.Id(strings.Title(prop.Name)).Id(convertType(prop.Type, prop.Collection)).Tag(propTag(prop.Name, prop.Required)))
		}
		if len(fields) > 0 {
			f.Type().Id(strings.Title(str.Name)).Struct(fields...)
		}
	}

	for _, mod := range schema.Modules {
		fields := []jen.Code{jen.Add(ptr).Id("qrpc.Client")}
		f.Type().Id(strings.Title(mod.Name) + "Module").Struct(fields...)
		for _, method := range mod.Methods {
			generateMethod(f, strings.Title(mod.Name), mod.Path, method)
		}
	}

	for _, cls := range schema.Classes {
		f.Type().Id(strings.Title(cls.Name)).Struct()

		for _, method := range cls.StaticMethods {
			generateMethod(f, strings.Title(cls.Name), cls.Path, method)
		}
		paramName := strings.Title(cls.Name) + "Params"
		fields = ParamFields(f, paramName, cls.ConstructorMethod.Parameters)
		if len(fields) > 0 {
			f.Type().Id(paramName).Struct(fields...)
		}

		skip := false
		for _, name := range moduleFields {
			if name == strings.Title(cls.Name) {
				skip = true
				break
			}
		}
		if skip {
			continue
		}
		fields := []jen.Code{jen.Add(ptr).Id("qrpc.Client")}
		f.Type().Id(strings.Title(cls.Name) + "Module").Struct(fields...)
	}

	generated := strings.Replace(fmt.Sprintf("%#v", f), "package electrond\n", "", 1)
	fmt.Printf("%s\n%s", header, generated)
}

func generateMethod(f *jen.File, nsName string, nsPath string, method Method) {
	methodName := strings.Title(method.Name)
	rcvrType := "*" + nsName + "Module"
	paramName := fmt.Sprintf("%s%sParams", nsName, strings.Title(method.Name))
	hasParams := len(method.Parameters) > 0
	// method function definition
	if hasParams {
		f.Func().Params(jen.Id("c").Id(rcvrType)).Id(methodName).Params(jen.Id("params").Id(paramName), jen.Id("ret").Interface()).Error().Block(
			jen.Return(jen.Id("c").Dot("Call").Call(jen.Lit(nsPath+"."+method.Name), jen.Id("params"), jen.Id("ret"))),
		)
	} else {
		f.Func().Params(jen.Id("c").Id(rcvrType)).Id(methodName).Params(jen.Id("ret").Interface()).Error().Block(
			jen.Return(jen.Id("c").Dot("Call").Call(jen.Lit(nsPath+"."+method.Name), jen.Nil(), jen.Id("ret"))),
		)
	}
	// parameter struct
	fields := ParamFields(f, paramName, method.Parameters)
	if len(fields) > 0 {
		f.Type().Id(paramName).Struct(fields...)
	}
	// optional return struct
	if method.Returns.Type == "Object" {
		var fields []jen.Code
		for _, param := range method.Returns.Parameters {
			fields = append(fields, jen.Id(strings.Title(param.Name)).Id(convertType(param.Type, param.Collection)).Tag(propTag(param.Name, param.Required)))
		}
		if len(fields) > 0 {
			f.Type().Id(fmt.Sprintf("%s%sReturn", nsName, strings.Title(method.Name))).Struct(fields...)
		}
	}
}

func ParamFields(f *jen.File, nsName string, params []Parameter) []jen.Code {
	var fields []jen.Code
	for _, param := range params {
		if param.Type == "Object" {
			var subfields []jen.Code
			for _, prop := range param.Properties {
				subfields = append(subfields, jen.Id(strings.Title(prop.Name)).Id(convertType(prop.Type, prop.Collection)).Tag(propTag(prop.Name, false)))
			}
			subName := nsName + strings.Title(param.Name)
			f.Type().Id(subName).Struct(subfields...)
			fields = append(fields, jen.Id(strings.Title(param.Name)).Id(subName).Tag(propTag(param.Name, param.Required)))
		} else {
			fields = append(fields, jen.Id(strings.Title(param.Name)).Id(convertType(param.Type, param.Collection)).Tag(propTag(param.Name, param.Required)))
		}
	}
	return fields
}

type Parameter struct {
	Name           string      `json:"name"`
	Type           interface{} `json:"type"`
	Collection     bool        `json:"collection"`
	Required       bool        `json:"required"`
	PossibleValues []struct {
		Value string `json:"value"`
	}
	Properties []struct {
		Name        string      `json:"name"`
		Description string      `json:"description,omitempty"`
		Type        interface{} `json:"type"`
		Collection  bool        `json:"collection"`
	} `json:"properties"`
}

type Method struct {
	Name        string   `json:"name"`
	Signature   string   `json:"signature"`
	Description string   `json:"description,omitempty"`
	Platforms   []string `json:"platforms,omitempty"`
	Returns     struct {
		Type        interface{} `json:"type"`
		Collection  bool        `json:"collection"`
		Description string      `json:"description"`
		Parameters  []Parameter `json:"parameters,omitempty"`
	} `json:"returns,omitempty"`
	Parameters []Parameter `json:"parameters,omitempty"`
}

type Schema struct {
	Modules []struct {
		Path        string `json:"path"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Process     struct {
			Main     bool `json:"main"`
			Renderer bool `json:"renderer"`
		} `json:"process"`
		Version    string   `json:"version"`
		Type       string   `json:"type"`
		Slug       string   `json:"slug"`
		WebsiteURL string   `json:"websiteUrl"`
		RepoURL    string   `json:"repoUrl"`
		Methods    []Method `json:"methods,omitempty"`
		Events     []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Returns     []struct {
				Name       string      `json:"name"`
				Type       interface{} `json:"type"`
				Collection bool        `json:"collection"`
				Required   bool        `json:"required"`
			} `json:"returns"`
		} `json:"events,omitempty"`
		Properties []struct {
			Name       string      `json:"name"`
			Type       interface{} `json:"type"`
			Required   bool        `json:"required"`
			Properties []struct {
				Name        string      `json:"name"`
				Signature   string      `json:"signature"`
				Platforms   []string    `json:"platforms"`
				Description string      `json:"description,omitempty"`
				Parameters  []Parameter `json:"parameters,omitempty"`
				Returns     struct {
					Type        interface{} `json:"type"`
					Collection  bool        `json:"collection"`
					Description string      `json:"description"`
				} `json:"returns,omitempty"`
				Type string `json:"type"`
			} `json:"properties"`
		} `json:"properties,omitempty"`
	} `json:"modules"`
	Classes []struct {
		Path        string `json:"path"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Process     struct {
			Main     bool `json:"main"`
			Renderer bool `json:"renderer"`
		} `json:"process"`
		Version            string   `json:"version"`
		Type               string   `json:"type"`
		Slug               string   `json:"slug"`
		WebsiteURL         string   `json:"websiteUrl"`
		RepoURL            string   `json:"repoUrl"`
		StaticMethods      []Method `json:"staticMethods,omitempty"`
		ConstructorMethod  Method   `json:"constructorMethod,omitempty"`
		InstanceName       string   `json:"instanceName"`
		InstanceMethods    []Method `json:"instanceMethods,omitempty"`
		InstanceProperties []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Type        string `json:"type"`
			Collection  bool   `json:"collection"`
		} `json:"instanceProperties,omitempty"`
		InstanceEvents []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Returns     []struct {
				Name       string      `json:"name"`
				Type       interface{} `json:"type"`
				Collection bool        `json:"collection"`
				Required   bool        `json:"required"`
			} `json:"returns"`
		} `json:"instanceEvents,omitempty"`
	} `json:"classes"`
	Structs []struct {
		Name       string      `json:"name"`
		Version    string      `json:"version"`
		Type       interface{} `json:"type"`
		Slug       string      `json:"slug"`
		WebsiteURL string      `json:"websiteUrl"`
		RepoURL    string      `json:"repoUrl"`
		Properties []struct {
			Name        string      `json:"name"`
			Type        interface{} `json:"type"`
			Collection  bool        `json:"collection"`
			Description string      `json:"description"`
			Required    bool        `json:"required"`
		} `json:"properties"`
	} `json:"structs"`
	Types []string `json:"types"`
}
