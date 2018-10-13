package goja

import (
	"github.com/dop251/goja"
	reflected "github.com/progrium/prototypes/go-reflected"
	vtemplate "github.com/progrium/prototypes/go-vtemplate"
)

type evaluator struct {
	*goja.Runtime
}

func Evaluator() vtemplate.Evaluator {
	return &evaluator{
		goja.New(),
	}
}

func (e *evaluator) Set(name string, value interface{}) {
	e.Runtime.Set(name, value)
}
func (e *evaluator) Unset(name string) {
	e.Runtime.Set(name, goja.Undefined())
}
func (e *evaluator) Eval(exp string) (reflected.Value, error) {
	v, err := e.RunString(exp)
	if err != nil {
		return reflected.Undefined(), err
	}
	return reflected.ValueOf(v.Export()), nil
}
