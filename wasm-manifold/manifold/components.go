package manifold

import (
	"encoding/json"

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

func componentFromValue(v interface{}) Component {
	return Component{
		Name: reflected.ValueOf(v).Type().Name(),
		Ref:  v,
	}
}

type Component struct {
	Name string
	Ref  interface{}
}

type componentData struct {
	Name string
	Data json.RawMessage
}

func (c *Component) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Name": c.Name,
		"Data": c.Ref,
	}
	return json.Marshal(m)
}

func (c *Component) UnmarshalJSON(b []byte) error {
	var cd componentData
	err := json.Unmarshal(b, &cd)
	if err != nil {
		return err
	}
	c.Name = cd.Name
	com := NewComponent(c.Name)
	err = json.Unmarshal(cd.Data, com)
	if err != nil {
		return err
	}
	c.Ref = com
	return nil
}

type ComponentSet struct {
	Components []Component
	registry   *objects.Registry
}

func (c *ComponentSet) RemoveAt(idx int) interface{} {
	v := c.Components[idx]
	c.Components = append(c.Components[:idx], c.Components[idx+1:]...)
	return v.Ref
}

func (c *ComponentSet) Insert(idx int, v interface{}) {
	com := componentFromValue(v)
	c.Components = append(c.Components[:idx], append([]Component{com}, c.Components[idx:]...)...)
}

func (c *ComponentSet) Append(v interface{}) {
	com := componentFromValue(v)
	c.Components = append(c.Components, com)
}
