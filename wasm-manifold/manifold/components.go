package manifold

import (
	"github.com/gliderlabs/com/objects"
	reflected "github.com/progrium/prototypes/go-reflected"
)

var registeredComponents []reflected.Type

func RegisterComponent(v interface{}) {
	registeredComponents = append(registeredComponents, reflected.ValueOf(v).Type())
}

func RegisteredComponents() []string {
	var names []string
	for _, typ := range registeredComponents {
		names = append(names, typ.Name())
	}
	return names
}

func NewComponent(name string) interface{} {
	for _, typ := range registeredComponents {
		if typ.Name() == name {
			return reflected.New(typ).Interface()
		}
	}
	return nil
}

type Components struct {
	components []interface{}
	registry   *objects.Registry
}

func (c *Components) RemoveAt(idx int) interface{} {
	v := c.components[idx]
	c.components = append(c.components[:idx], c.components[idx+1:]...)
	return v
}

func (c *Components) Insert(idx int, v interface{}) {
	c.components = append(c.components[:idx], append([]interface{}{v}, c.components[idx:]...)...)
}

func (c *Components) Append(v interface{}) {
	c.components = append(c.components, v)
}
