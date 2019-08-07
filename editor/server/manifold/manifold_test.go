package manifold

import (
	"testing"
)

type SubStruct struct {
	BoolValue bool
}

type DemoComponent struct {
	StringValue string
	IntValue    int
	StructValue SubStruct

	TestBool   bool
	TestString string
}

func newDemoComponent() *DemoComponent {
	return &DemoComponent{
		StringValue: "Foo",
		IntValue:    64,
		StructValue: SubStruct{
			true,
		},
		TestBool:   false,
		TestString: "test",
	}
}

func newTree() *Node {
	n := NewNode("Root")
	n1 := NewNode("Node1")
	n1.Append(NewNode("NodeA"))
	n1.Append(NewNode("NodeB"))
	n2 := NewNode("Node2")
	n2child := NewNode("Foo")
	n2child.Append(NewNode("Bar"))
	n2.Append(n2child)
	n3 := NewNode("Node3")
	n.Append(n1)
	n.Append(n2)
	n.Append(n3)
	return n
}

func TestAppend(t *testing.T) {
	n := NewNode("Node1")
	n.Append(NewNode("Node2"))
	if n.Children[0].Name != "Node2" {
		t.Fatal("Node2 not appended")
	}
}

func TestChild(t *testing.T) {
	n := NewNode("Node1")
	n.Append(NewNode("Node2"))
	if n.Child("Node2").Name != "Node2" {
		t.Fatal("Node2 not found")
	}
	if n.Child("Node3") != nil {
		t.Fatal("Node3 unexpectedly found")
	}
}

func TestRoot(t *testing.T) {
	n1 := NewNode("Node1")
	n2 := NewNode("Node2")
	n3 := NewNode("Node3")
	n2.Append(n3)
	n1.Append(n2)
	if n3.Root() != n1 {
		t.Fatal("Node1 root not found")
	}
}

func TestFindNode(t *testing.T) {
	n := newTree()
	found := n.FindNode("/Node2/Foo/Bar")
	if found == nil {
		t.Fatal("node not found")
	}
	v := newDemoComponent()
	found.AppendComponent(v)
	found = n.FindNode("Node2/Foo")
	if found == nil {
		t.Fatal("node not found")
	}
	found = found.FindNode("Bar")
	if found == nil {
		t.Fatal("node not found")
	}
	found = found.FindNode("../../../Node1/NodeA")
	if found == nil {
		t.Fatal("node not found")
	}
	found = found.FindNode("/Node3")
	if found == nil {
		t.Fatal("node not found")
	}
	found = found.FindNode("Baz")
	if found != nil {
		t.Fatal("node unexpectedly found")
	}
	found = n.FindNode("/Node3").FindNode("/Node2/Foo/Bar/DemoComponent/StringValue")
	expect := n.FindNode("/Node2/Foo/Bar")
	if found != expect {
		t.Fatalf("got: %#v, wanted: %#v", found, expect)
	}
}

func TestComponent(t *testing.T) {
	n := NewNode("Node1")
	v := newDemoComponent()
	n.AppendComponent(v)
	vv := n.RemoveComponent("DemoComponent")
	if vv != v {
		t.Fatal("diff value")
	}
}

func TestSync(t *testing.T) {
	n := NewNode("Node1")
	v := newDemoComponent()
	n.AppendComponent(v)
	n.Sync()
	v.StringValue = "new value"
	v.StructValue.BoolValue = false
	n.Sync()
}

func TestSetValue(t *testing.T) {
	n := NewNode("Node1")
	v := newDemoComponent()
	n.AppendComponent(v)
	n.SetValue("DemoComponent/IntValue", 100)
	if v.IntValue != 100 {
		t.Fatal("IntValue not set")
	}
}

func TestFullPath(t *testing.T) {
	n := newTree()
	found := n.FindNode("/Node2/Foo/Bar")
	if found.FullPath() != "/Node2/Foo/Bar" {
		t.Fatalf("not expected path: %s", found.FullPath())
	}
}

func TestExpressions(t *testing.T) {
	r := newTree()
	n := r.Child("Node1")
	v := newDemoComponent()
	n.AppendComponent(v)
	expect := "/Node1/DemoComponent/TestString"
	n.SetExpression("DemoComponent/StringValue", expect)
	if got := n.Expression("DemoComponent/StringValue"); got != expect {
		t.Fatalf("got: %#v, want: %#v", got, expect)
	}
	got := n.Value("DemoComponent/StringValue")
	if got != "test" {
		t.Fatalf("got: %s, wanted: test", got)
	}
	v.TestString = "test2"
	n.Sync()
	got = n.Value("DemoComponent/StringValue")
	if got != "test2" {
		t.Fatalf("got: %s, wanted: test2", got)
	}
}

func TestObserver(t *testing.T) {
	n := NewNode("Node1")
	v := newDemoComponent()
	n.AppendComponent(v)
	ch := make(chan bool, 1)
	n.Observe(&NodeObserver{
		Path: "DemoComponent/StringValue",
		OnChange: func(old, new interface{}, path string) {
			ch <- true
		},
	})
	n.Component("DemoComponent").(*DemoComponent).StringValue = "newval"
	n.Sync()
	if <-ch != true {
		t.Fatal("got bad value")
	}
}
