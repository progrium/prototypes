// Package reflected provides a higher-level reflection API that makes working with
// Go values feel more like a dynamic language. The API was inspired by syscall/js
// for interacting with JavaScript objects.
package reflected

import (
	"reflect"
	"sort"
	"strings"
)

// Type is the representation of a Go type.
type Type struct {
	reflect.Type
}

// Fields returns exported field names for Type t.
// If it is not a struct, it returns an empty slice.
func (t Type) Fields() []string {
	var f []string
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		if name[0] == strings.ToUpper(name)[0] {
			f = append(f, name)
		}
	}
	return f
}

func (t Type) FieldType(f string) Type {
	field, _ := t.FieldByName(f)
	return Type{field.Type}
}

// Methods returns method names for Type t and reflect.PtrTo(t).
func (t Type) Methods() []string {
	var methods []string

	for idx := 0; idx < reflect.PtrTo(t.Type).NumMethod(); idx++ {
		methods = append(methods, reflect.PtrTo(t.Type).Method(idx).Name)
	}
	// for idx := 0; idx < t.NumMethod(); idx++ {
	// 	methods = append(methods, t.Method(idx).Name)
	// }
	return methods
}

// HasField returns whether or not a struct has a field by name f.
func (t Type) HasField(f string) bool {
	for _, field := range t.Fields() {
		if f == field {
			return true
		}
	}
	return false
}

// HasMethod returns whether or not a value has a method by name m.
func (t Type) HasMethod(m string) bool {
	for _, meth := range t.Methods() {
		if meth == m {
			return true
		}
	}
	return false
}

// FieldsTagged returns field names that have a struct tag
// including a particular key, or if value is provided it returns
// fields that include that key and value.
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

// Value is the reflection interface to a Go value.
type Value struct {
	v reflect.Value
}

// Undefined returns a Value representing a non-value similar to JavaScript undefined.
func Undefined() Value {
	return Value{}
}

// ValueOf returns a new Value initialized to the concrete value stored in the interface v. ValueOf(nil) returns the zero Value.
func ValueOf(v interface{}) Value {
	return Value{
		v: reflect.ValueOf(v),
	}
}

// New returns a Value representing a pointer to a new zero value for the specified type. That is, the returned Value's Type is reflect.PtrTo(t).
func New(t Type) Value {
	return Value{v: reflect.New(t.Type)}
}

// Kind returns v's Kind. If v is the zero Value (IsValid returns false), Kind returns Invalid.
func (v Value) Kind() reflect.Kind {
	return v.v.Kind()
}

func (v Value) IsValid() bool {
	return v.v.IsValid()
}

// IsNil reports whether its argument v is nil. The argument must be a chan, func, interface, map, pointer, or slice value; if it is not, IsNil panics. Note that IsNil is not always equivalent to a regular comparison with nil in Go. For example, if v was created by calling ValueOf with an uninitialized interface variable i, i==nil will be true but v.IsNil will panic as v will be the zero Value.
func (v Value) IsNil() bool {
	if v.v.Kind() == reflect.Invalid && !v.v.IsValid() {
		return true
	}
	return v.v.IsNil()
}

// Interface returns v's current value as an interface{}.
func (v Value) Interface() interface{} {
	return v.v.Interface()
}

// Bool returns v's underlying value. It panics if v's kind is not Bool.
func (v Value) Bool() bool {
	return v.v.Bool()
}

// Float returns v's underlying value, as a float64. It panics if v's Kind is not Float32 or Float64
func (v Value) Float() float64 {
	return v.v.Float()
}

// Int returns v's underlying value, as an int. It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64.
func (v Value) Int() int {
	return int(v.v.Int())
}

// String returns the string v's underlying value, as a string. String is a special case because of Go's String method convention.
// Unlike the other getters, it does not panic if v's Kind is not String. Instead, it returns a string of the form "<T value>"
// where T is v's type. The fmt package treats Values specially. It does not call their String method implicitly but instead prints
// the concrete values they hold.
func (v Value) String() string {
	return v.v.String()
}

// Len returns v's length. It panics if v's Kind is not Array, Chan, Map, Slice, or String.
func (v Value) Len() int {
	return v.v.Len()
}

// Index returns v's i'th element. It panics if v's Kind is not Array, Slice, or String or i is out of range.
func (v Value) Index(i int) Value {
	return Value{v: v.v.Index(i)}
}

// SetIndex sets the index i of value v to ValueOf(x).
func (v Value) SetIndex(i int, x interface{}) {
	v.v.Index(i).Set(reflect.ValueOf(x))
}

func (v Value) valueOrElem() reflect.Value {
	vv := v.v
	if v.v.Type().Kind() == reflect.Ptr {
		vv = v.v.Elem()
	}
	return vv
}

// Set sets the field or key m of value v to ValueOf(x).
func (v Value) Set(m string, x interface{}) {
	var fv reflect.Value
	if v.v.Type().Kind() == reflect.Map {
		v.v.SetMapIndex(reflect.ValueOf(m), reflect.ValueOf(x))
	} else {
		fv = v.valueOrElem().FieldByName(m)
		fv.Set(reflect.ValueOf(x))
	}

}

// Iter returns a slice of Values for the values contained in Value v if it is an Array or Slice.
// It panics if v's Kind is not Array or Slice.
func (v Value) Iter() []Value {
	kind := v.v.Type().Kind()
	// TODO: maps?
	if kind != reflect.Array && kind != reflect.Slice {
		panic("Iter called on value that is not an Array or Slice")
	}
	var vals []Value
	for i := 0; i < v.v.Len(); i++ {
		vals = append(vals, Value{v: v.v.Index(i)})
	}
	return vals
}

// Members returns the names of keys, fields, and methods of Value v.
func (v Value) Members() []string {
	members := v.Keys()
	if v.v.Type().Kind() == reflect.Map {
		return members
	}
	for _, m := range v.Methods() {
		members = append(members, m)
	}
	return members
}

// HasMember returns whether Value v has the member m.
func (v Value) HasMember(m string) bool {
	for _, member := range v.Members() {
		if m == member {
			return true
		}
	}
	return false
}

// Methods returns the names of methods on Value v.
func (v Value) Methods() []string {
	var methods []string
	for idx := 0; idx < v.v.NumMethod(); idx++ {
		methods = append(methods, v.v.Type().Method(idx).Name)
	}
	return methods
}

// HasMethod returns whether Value v has the method m.
func (v Value) HasMethod(m string) bool {
	for _, meth := range v.Methods() {
		if m == meth {
			return true
		}
	}
	return false
}

// HasKey returns whether Value v has the field or map key k.
func (v Value) HasKey(k string) bool {
	for _, key := range v.Keys() {
		if k == key {
			return true
		}
	}
	return false
}

// Keys returns the names of settable fields or map keys of Value v.
func (v Value) Keys() []string {
	if v.v.Type().Kind() == reflect.Map {
		var keys []string
		for _, key := range v.v.MapKeys() {
			k, ok := key.Interface().(string)
			if !ok {
				continue
			}
			keys = append(keys, k)
		}
		sort.Sort(sort.StringSlice(keys))
		return keys
	}
	return v.Type().Fields()
}

// Get returns the member value by name m of value v. Members include map keys, struct fields, and methods.
func (v Value) Get(m string) Value {
	var fv reflect.Value
	if v.v.Type().Kind() == reflect.Map {
		fv = v.v.MapIndex(reflect.ValueOf(m))
	} else {
		if v.HasMethod(m) {
			fv = v.v.MethodByName(m)
		} else {
			fv = v.valueOrElem().FieldByName(m)
		}
	}
	zero := reflect.Value{}
	if fv == zero {
		return Undefined()
	}
	return ValueOf(fv.Interface())
}

// Call does a call to the method m of value v with the given arguments. It panics if v has no method m.
// Call returns an empty slice if there are no return values.
func (v Value) Call(m string, args ...interface{}) []Value {
	var in []reflect.Value
	for _, arg := range args {
		in = append(in, reflect.ValueOf(arg))
	}
	var out []Value
	method := v.v.MethodByName(m)
	if !method.IsValid() {
		panic("call to undefined method: " + m)
	}
	ret := method.Call(in)
	for _, v := range ret {
		out = append(out, ValueOf(v.Interface()))
	}
	return out
}

// Invoke does a call of the value v with the given arguments. It panics if v is not a function.
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

// Type returns the Type for Value v.
func (v Value) Type() Type {
	return Type{v.valueOrElem().Type()}
}

func TypeOf(v interface{}) Type {
	return ValueOf(v).Type()
}
