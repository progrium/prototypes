package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/dave/jennifer/jen"
)

var typeMap = map[string]string{
	"String":      "string",
	"Accelerator": "string",
	"Boolean":     "bool",
	"Number":      "int",
	"Buffer":      "[]byte",
	"Function":    "qrpc.ObjectHandle",
	"Integer":     "int",
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
		return pre + first["typeName"].(string)
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
	fields := []jen.Code{jen.Add(ptr).Id("qrpc.Client")}
	var setup []jen.Code
	for _, mod := range schema.Modules {
		var name string
		if _, ok := submodules[mod.Path]; ok {
			name = strings.Replace(strings.Title(strings.Replace(mod.Path, ".", " ", -1)), " ", "", -1)
		} else {
			name = strings.Title(mod.Name)
		}
		fields = append(fields, jen.Id(name).Add(ptr).Id(strings.Title(mod.Name)))
		setup = append(setup, jen.Id("c").Dot(name).Op("=").Add(jen.Op("&")).Id(strings.Title(mod.Name)).Values(jen.Dict{
			jen.Id("Client"): jen.Id("c").Dot("Client"),
		}))
	}
	f.Type().Id("Client").Struct(fields...)
	f.Func().Params(jen.Id("c").Id("*Client")).Id("setup").Params().Id("*Client").Block(append(setup, jen.Return(jen.Id("c")))...)

	f.Type().Id("NativeImage").Struct()

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
		// for path, name := range submodules {
		// 	if strings.HasPrefix(path, mod.Name) {
		// 		fields = append(fields, jen.Id(strings.Title(name)).Add(ptr).Id(strings.Title(name)))
		// 	}
		// }
		f.Type().Id(strings.Title(mod.Name)).Struct(fields...)
		for _, method := range mod.Methods {
			methodName := strings.Title(method.Name)
			rcvrType := "*" + strings.Title(mod.Name)
			paramName := fmt.Sprintf("%s%sParams", strings.Title(mod.Name), strings.Title(method.Name))
			hasParams := len(method.Parameters) > 0
			if hasParams {
				f.Func().Params(jen.Id("c").Id(rcvrType)).Id(methodName).Params(jen.Id("params").Id(paramName), jen.Id("ret").Interface()).Error().Block(
					jen.Return(jen.Id("c").Dot("Call").Call(jen.Lit(mod.Path+"."+method.Name), jen.Id("params"), jen.Id("ret"))),
				)
			} else {
				f.Func().Params(jen.Id("c").Id(rcvrType)).Id(methodName).Params(jen.Id("ret").Interface()).Error().Block(
					jen.Return(jen.Id("c").Dot("Call").Call(jen.Lit(mod.Path+"."+method.Name), jen.Nil(), jen.Id("ret"))),
				)
			}
			var fields []jen.Code
			for _, param := range method.Parameters {
				if param.Type == "Object" {
					var subfields []jen.Code
					for _, prop := range param.Properties {
						subfields = append(subfields, jen.Id(strings.Title(prop.Name)).Id(convertType(prop.Type, prop.Collection)).Tag(propTag(prop.Name, false)))
					}
					fields = append(fields, jen.Id(strings.Title(param.Name)).Struct(subfields...).Tag(propTag(param.Name, param.Required)))
				} else {
					fields = append(fields, jen.Id(strings.Title(param.Name)).Id(convertType(param.Type, param.Collection)).Tag(propTag(param.Name, param.Required)))
				}
			}
			if len(fields) > 0 {
				f.Type().Id(paramName).Struct(fields...)
			}
			if method.Returns.Type == "Object" {
				var fields []jen.Code
				for _, param := range method.Returns.Parameters {
					fields = append(fields, jen.Id(strings.Title(param.Name)).Id(convertType(param.Type, param.Collection)).Tag(propTag(param.Name, param.Required)))
				}
				if len(fields) > 0 {
					f.Type().Id(fmt.Sprintf("%s%sReturn", strings.Title(mod.Name), strings.Title(method.Name))).Struct(fields...)
				}
			}

		}
	}
	generated := strings.Replace(fmt.Sprintf("%#v", f), "package electrond\n", "", 1)
	fmt.Printf("%s\n%s", header, generated)
}

type Parameter struct {
	Name       string      `json:"name"`
	Type       interface{} `json:"type"`
	Collection bool        `json:"collection"`
	Required   bool        `json:"required"`
	Properties []struct {
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
		Type        string `json:"type"`
		Collection  bool   `json:"collection"`
	} `json:"properties"`
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
		Version    string `json:"version"`
		Type       string `json:"type"`
		Slug       string `json:"slug"`
		WebsiteURL string `json:"websiteUrl"`
		RepoURL    string `json:"repoUrl"`
		Methods    []struct {
			Name        string   `json:"name"`
			Signature   string   `json:"signature"`
			Description string   `json:"description,omitempty"`
			Platforms   []string `json:"platforms,omitempty"`
			Returns     struct {
				Type        interface{} `json:"type"`
				Collection  bool        `json:"collection"`
				Description string      `json:"description"`
				Parameters  []struct {
					Name       string      `json:"name"`
					Type       interface{} `json:"type"`
					Collection bool        `json:"collection"`
					Required   bool        `json:"required"`
				} `json:"parameters,omitempty"`
			} `json:"returns,omitempty"`
			Parameters []struct {
				Name       string      `json:"name"`
				Type       interface{} `json:"type"`
				Collection bool        `json:"collection"`
				Required   bool        `json:"required"`
				Properties []struct {
					Name        string `json:"name"`
					Description string `json:"description,omitempty"`
					Type        string `json:"type"`
					Collection  bool   `json:"collection"`
				} `json:"properties"`
			} `json:"parameters,omitempty"`
		} `json:"methods,omitempty"`
		Events []struct {
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
				Name        string   `json:"name"`
				Signature   string   `json:"signature"`
				Platforms   []string `json:"platforms"`
				Description string   `json:"description,omitempty"`
				Parameters  []struct {
					Name           string      `json:"name"`
					Type           interface{} `json:"type"`
					Collection     bool        `json:"collection"`
					Description    string      `json:"description"`
					Required       bool        `json:"required"`
					PossibleValues []struct {
						Value string `json:"value"`
					} `json:"possibleValues"`
				} `json:"parameters,omitempty"`
				Returns struct {
					Type        interface{} `json:"type"`
					Collection  bool        `json:"collection"`
					Description string      `json:"description"`
				} `json:"returns,omitempty"`
				Type string `json:"type"`
			} `json:"properties"`
		} `json:"properties,omitempty"`
	} `json:"modules"`
	Classes []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Process     struct {
			Main     bool `json:"main"`
			Renderer bool `json:"renderer"`
		} `json:"process"`
		Version       string `json:"version"`
		Type          string `json:"type"`
		Slug          string `json:"slug"`
		WebsiteURL    string `json:"websiteUrl"`
		RepoURL       string `json:"repoUrl"`
		StaticMethods []struct {
			Name        string `json:"name"`
			Signature   string `json:"signature"`
			Description string `json:"description"`
			Parameters  []struct {
				Name       string      `json:"name"`
				Type       interface{} `json:"type"`
				Collection bool        `json:"collection"`
				Required   bool        `json:"required"`
			} `json:"parameters,omitempty"`
			Returns struct {
				Type        interface{} `json:"type"`
				Collection  bool        `json:"collection"`
				Description string      `json:"description"`
			} `json:"returns,omitempty"`
			Platforms []string `json:"platforms,omitempty"`
		} `json:"staticMethods,omitempty"`
		ConstructorMethod struct {
			Signature string `json:"signature"`
		} `json:"constructorMethod,omitempty"`
		InstanceName    string `json:"instanceName"`
		InstanceMethods []struct {
			Name        string `json:"name"`
			Signature   string `json:"signature"`
			Description string `json:"description,omitempty"`
			Parameters  []struct {
				Name       string      `json:"name"`
				Type       interface{} `json:"type"`
				Collection bool        `json:"collection"`
				Required   bool        `json:"required"`
				Properties []struct {
					Name        string        `json:"name"`
					Type        interface{}   `json:"type"`
					Collection  bool          `json:"collection"`
					Description string        `json:"description"`
					Required    bool          `json:"required"`
					Parameters  []interface{} `json:"parameters,omitempty"`
				} `json:"properties"`
			} `json:"parameters"`
			Returns struct {
				Type        interface{} `json:"type"`
				Collection  bool        `json:"collection"`
				Description string      `json:"description"`
			} `json:"returns,omitempty"`
		} `json:"instanceMethods,omitempty"`
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
		Name       string `json:"name"`
		Version    string `json:"version"`
		Type       string `json:"type"`
		Slug       string `json:"slug"`
		WebsiteURL string `json:"websiteUrl"`
		RepoURL    string `json:"repoUrl"`
		Properties []struct {
			Name        string `json:"name"`
			Type        string `json:"type"`
			Collection  bool   `json:"collection"`
			Description string `json:"description"`
			Required    bool   `json:"required"`
		} `json:"properties"`
	} `json:"structs"`
	Types []string `json:"types"`
}
