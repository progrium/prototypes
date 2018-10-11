package reflected

import (
	"reflect"
)

type Type struct {
	reflect.Type
}

func (t Type) Fields() []string {
	var f []string
	for i := 0; i < t.NumField(); i++ {
		f = append(f, t.Field(i).Name)
	}
	return f
}

func (t Type) FieldsTagged(key, value string) []string {
	var f []string
	for i := 0; i < t.NumField(); i++ {
		v, ok := t.Field(i).Tag.Lookup(key)
		if !ok {
			continue
		}
		if value != "" && v != value {
			continue
		}
		f = append(f, t.Field(i).Name)
	}
	return f
}

type Value struct {
	v reflect.Value
}

func Undefined() Value {
	return Value{}
}

func ValueOf(v interface{}) Value {
	return Value{
		v: reflect.ValueOf(v),
	}
}

func (v Value) Interface() interface{} {
	return v.v.Interface()
}

func (v Value) Bool() bool {
	return v.v.Bool()
}

func (v Value) Float() float64 {
	return v.v.Float()
}

func (v Value) Int() int {
	return int(v.v.Int())
}

func (v Value) String() string {
	return v.v.String()
}

func (v Value) valueOrElem() reflect.Value {
	vv := v.v
	if v.v.Type().Kind() == reflect.Ptr {
		vv = v.v.Elem()
	}
	return vv
}

func (v Value) Set(f string, x interface{}) {
	var fv reflect.Value
	if v.v.Type().Kind() == reflect.Map {
		v.v.SetMapIndex(reflect.ValueOf(f), reflect.ValueOf(x))
	} else {
		fv = v.valueOrElem().FieldByName(f)
		fv.Set(reflect.ValueOf(x))
	}

}

func (v Value) Iter() []Value {
	kind := v.v.Type().Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return nil
	}
	var vals []Value
	for i := 0; i < v.v.Len(); i++ {
		vals = append(vals, Value{v: v.v.Index(i)})
	}
	return vals
}

func (v Value) Members() []string {
	members := v.Props()
	if v.v.Type().Kind() == reflect.Map {
		return members
	}
	for _, m := range v.Methods() {
		members = append(members, m)
	}
	return members
}

func (v Value) Methods() []string {
	var methods []string
	for idx := 0; idx < v.v.NumMethod(); idx++ {
		methods = append(methods, v.v.Type().Method(idx).Name)
	}
	return methods
}

func (v Value) Props() []string {
	if v.v.Type().Kind() == reflect.Map {
		var keys []string
		for _, key := range v.v.MapKeys() {
			k, ok := key.Interface().(string)
			if !ok {
				continue
			}
			keys = append(keys, k)
		}
		return keys
	}
	return v.Type().Fields()
}

func (v Value) Get(f string) Value {
	var fv reflect.Value
	if v.v.Type().Kind() == reflect.Map {
		fv = v.v.MapIndex(reflect.ValueOf(f))
	} else {
		fv = v.valueOrElem().FieldByName(f)
	}
	zero := reflect.Value{}
	if fv == zero {
		return Undefined()
	}
	// TODO: return methods?
	return ValueOf(fv.Interface())
}

func (v Value) Call(m string, args ...interface{}) []Value {
	var in []reflect.Value
	for _, arg := range args {
		in = append(in, reflect.ValueOf(arg))
	}
	var out []Value
	ret := v.v.MethodByName(m).Call(in)
	for _, v := range ret {
		out = append(out, ValueOf(v.Interface()))
	}
	return out
}

func (v Value) Invoke(args ...interface{}) []Value {
	var in []reflect.Value
	for _, arg := range args {
		in = append(in, reflect.ValueOf(arg))
	}
	var out []Value
	ret := v.v.Call(in)
	for _, v := range ret {
		out = append(out, ValueOf(v.Interface()))
	}
	return out
}

// func (v Value) New() Value

// func (v Value) Length() int
// func (v Value) Index(i int) Value
// func (v Value) InstanceOf(t Value) bool
// func (v Value) SetIndex(i int, x interface{})

func New(t Type) Value {
	return Value{v: reflect.New(t.Type)}
}

func (v Value) Type() Type {
	return Type{v.valueOrElem().Type()}
}
