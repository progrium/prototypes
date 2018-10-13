package reflected

import (
	"sort"
	"testing"
)

type TestValue struct {
	StringField       string `test:""`
	IntField          int    `test:"foo"`
	BoolField         bool   `test:"bar"`
	FloatField        float64
	noArgsNoRetCalled bool
}

func (v *TestValue) NoArgsNoRet() {
	v.noArgsNoRetCalled = true
	return
}

func (v *TestValue) Echo(s string) string {
	return s
}

func (v *TestValue) MultiEcho(s1, s2 string) (string, string) {
	return s1, s2
}

func (v *TestValue) VariadicEcho(s ...string) []string {
	return s
}

func TestTypedValues(t *testing.T) {
	v := ValueOf(true)
	if v.Bool() != true {
		t.Fatalf("got %v, want %v", v.Bool(), true)
	}

	v = ValueOf(128)
	if v.Int() != 128 {
		t.Fatalf("got %v, want %v", v.Int(), 128)
	}

	v = ValueOf(0.1)
	if v.Float() != 0.1 {
		t.Fatalf("got %v, want %v", v.Float(), 0.1)
	}

	v = ValueOf("test")
	if v.String() != "test" {
		t.Fatalf("got %v, want %v", v.String(), "test")
	}

	v = ValueOf(&TestValue{})
	if _, ok := v.Interface().(*TestValue); !ok {
		t.Fatalf("unable to type assert to value of underlying type")
	}
}

func TestSetFields(t *testing.T) {
	value := &TestValue{}
	v := ValueOf(value)

	v.Set("StringField", "test")
	if value.StringField != "test" {
		t.Fatalf("got %v, want %v", value.StringField, "test")
	}

	v.Set("BoolField", true)
	if value.BoolField != true {
		t.Fatalf("got %v, want %v", value.StringField, "test")
	}

	v.Set("IntField", 64)
	if value.IntField != 64 {
		t.Fatalf("got %v, want %v", value.IntField, 64)
	}

	v.Set("FloatField", 0.1)
	if value.FloatField != 0.1 {
		t.Fatalf("got %v, want %v", value.FloatField, 0.1)
	}
}

func TestGetFields(t *testing.T) {
	value := &TestValue{
		StringField: "test",
		IntField:    64,
		BoolField:   true,
		FloatField:  0.1,
	}
	v := ValueOf(value)

	ret := v.Get("StringField")
	if ret.String() != "test" {
		t.Fatalf("got %v, want %v", ret.String(), "test")
	}

	ret = v.Get("BoolField")
	if ret.Bool() != true {
		t.Fatalf("got %v, want %v", ret.Bool(), true)
	}

	ret = v.Get("IntField")
	if ret.Int() != 64 {
		t.Fatalf("got %v, want %v", ret.Int(), 64)
	}

	ret = v.Get("FloatField")
	if ret.Float() != 0.1 {
		t.Fatalf("got %v, want %v", ret.Float(), 0.1)
	}
}

func TestMethodCalls(t *testing.T) {
	value := &TestValue{}
	v := ValueOf(value)

	v.Call("NoArgsNoRet")
	if !value.noArgsNoRetCalled {
		t.Fatalf("got %v, want %v", value.noArgsNoRetCalled, true)
	}

	ret := v.Call("Echo", "test")
	if ret[0].String() != "test" {
		t.Fatalf("got %v, want %v", ret[0].String(), "test")
	}

	ret = v.Call("MultiEcho", "test1", "test2")
	if ret[0].String() != "test1" {
		t.Fatalf("got %v, want %v", ret[0].String(), "test1")
	}
	if ret[1].String() != "test2" {
		t.Fatalf("got %v, want %v", ret[2].String(), "test2")
	}
}

func TestFunctionInvoke(t *testing.T) {
	v := ValueOf(func(s string) string {
		return s
	})
	ret := v.Invoke("test")
	if ret[0].String() != "test" {
		t.Fatalf("got %v, want %v", ret[0].String(), "test")
	}

}

func TestFunctionInvokeMethod(t *testing.T) {
	v := ValueOf(&TestValue{})
	m := v.Get("Echo")

	ret := m.Invoke("test")
	if ret[0].String() != "test" {
		t.Fatalf("got %v, want %v", ret[0].String(), "test")
	}

}

func TestGetSetMaps(t *testing.T) {
	value := map[string]interface{}{}
	v := ValueOf(value)

	v.Set("foo", "bar")
	ret := v.Get("foo")
	if ret.String() != "bar" {
		t.Fatalf("got %v, want %v", ret.String(), "bar")
	}
}

func TestGetSetFields(t *testing.T) {
	value := &TestValue{}
	v := ValueOf(value)

	v.Set("StringField", "test")
	ret := v.Get("StringField")
	if ret.String() != "test" {
		t.Fatalf("got %v, want %v", ret.String(), "test")
	}
}

func TestLen(t *testing.T) {
	v := ValueOf("test")
	if v.Len() != 4 {
		t.Fatalf("got %v, want %v", v.Len(), 4)
	}

	v = ValueOf([]string{"one", "two", "three"})
	if v.Len() != 3 {
		t.Fatalf("got %v, want %v", v.Len(), 3)
	}

	// TODO: channels, etc
}

func TestTypeFields(t *testing.T) {
	typ := ValueOf(&TestValue{}).Type()

	fields := typ.Fields()
	expect := []string{"StringField", "IntField", "BoolField", "FloatField"}
	if !eqStringSlice(fields, expect) {
		t.Fatalf("got %v, want %v", fields, expect)
	}

	got := typ.HasField("StringField")
	if !got {
		t.Fatalf("got %v, want %v", got, true)
	}
}

func TestTypeMethods(t *testing.T) {
	typ := ValueOf(&TestValue{}).Type()

	methods := typ.Methods()
	expect := []string{"Echo", "MultiEcho", "NoArgsNoRet", "VariadicEcho"}
	if !eqStringSlice(methods, expect) {
		t.Fatalf("got %v, want %v", methods, expect)
	}

	got := typ.HasMethod("Echo")
	if !got {
		t.Fatalf("got %v, want %v", got, true)
	}
}

func TestMethods(t *testing.T) {
	v := ValueOf(&TestValue{})

	methods := v.Methods()
	expect := []string{"Echo", "MultiEcho", "NoArgsNoRet", "VariadicEcho"}
	if !eqStringSlice(methods, expect) {
		t.Fatalf("got %v, want %v", methods, expect)
	}

	got := v.HasMethod("Echo")
	if !got {
		t.Fatalf("got %v, want %v", got, true)
	}
}

func TestMembers(t *testing.T) {
	v := ValueOf(&TestValue{})

	members := v.Members()
	expect := []string{"StringField", "IntField", "BoolField", "FloatField", "Echo", "MultiEcho", "NoArgsNoRet", "VariadicEcho"}
	if !eqStringSlice(members, expect) {
		t.Fatalf("got %v, want %v", members, expect)
	}

	got := v.HasMember("Echo")
	if !got {
		t.Fatalf("got %v, want %v", got, true)
	}
}

func TestKeysStruct(t *testing.T) {
	v := ValueOf(&TestValue{})

	fields := v.Keys()
	expect := []string{"StringField", "IntField", "BoolField", "FloatField"}
	if !eqStringSlice(fields, expect) {
		t.Fatalf("got %v, want %v", fields, expect)
	}

	got := v.HasKey("StringField")
	if !got {
		t.Fatalf("got %v, want %v", got, true)
	}
}

func TestKeysMap(t *testing.T) {
	v := ValueOf(map[string]interface{}{
		"one":   nil,
		"two":   nil,
		"three": nil,
	})

	keys := v.Keys()
	sort.Strings(keys)
	expect := []string{"one", "three", "two"}
	if !eqStringSlice(keys, expect) {
		t.Fatalf("got %v, want %v", keys, expect)
	}

	got := v.HasKey("three")
	if !got {
		t.Fatalf("got %v, want %v", got, true)
	}
}

func TestTypeTaggedFields(t *testing.T) {
	typ := ValueOf(&TestValue{}).Type()

	withKey := typ.FieldsTagged("test", "")
	expect := []string{"StringField", "IntField", "BoolField"}
	if !eqStringSlice(withKey, expect) {
		t.Fatalf("got %v, want %v", withKey, expect)
	}

	withKeyValue := typ.FieldsTagged("test", "foo")
	expect = []string{"IntField"}
	if !eqStringSlice(withKeyValue, expect) {
		t.Fatalf("got %v, want %v", withKeyValue, expect)
	}
}

func TestIter(t *testing.T) {
	fields := ValueOf(&TestValue{}).Type().Fields()
	rfields := ValueOf(fields)

	var f []string
	for _, v := range rfields.Iter() {
		f = append(f, v.String())
	}
	if !eqStringSlice(f, fields) {
		t.Fatalf("got %v, want %v", f, fields)
	}
}

func TestIndex(t *testing.T) {
	v := ValueOf([]string{"one", "two", "three"})

	got := v.Index(1).String()
	if got != "two" {
		t.Fatalf("got %v, want %v", got, "two")
	}

	v.SetIndex(1, "TWO")
	got = v.Index(1).String()
	if got != "TWO" {
		t.Fatalf("got %v, want %v", got, "TWO")
	}
}

func eqStringSlice(a, b []string) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
