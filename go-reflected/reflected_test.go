package reflected

import "testing"

type TestValue struct {
	StringField       string
	IntField          int
	BoolField         bool
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

func TestGetSetMaps(t *testing.T) {
	value := map[string]interface{}{}
	v := ValueOf(value)

	v.Set("foo", "bar")
	ret := v.Get("foo")
	if ret.String() != "bar" {
		t.Fatalf("got %v, want %v", ret.String(), "bar")
	}
}
