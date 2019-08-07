package repl

import (
	"bufio"
	"fmt"

	"io"
	"reflect"
	"time"

	"github.com/d5/tengo/compiler"
	"github.com/d5/tengo/compiler/ast"
	"github.com/d5/tengo/compiler/parser"
	"github.com/d5/tengo/compiler/source"
	"github.com/d5/tengo/compiler/token"
	"github.com/d5/tengo/objects"
	"github.com/d5/tengo/runtime"
)

type REPL struct {
	Prompt      string
	fileSet     *source.FileSet
	globals     []*objects.Object
	symbolTable *compiler.SymbolTable
}

func NewREPL(print func(interface{})) *REPL {
	repl := &REPL{
		fileSet:     source.NewFileSet(),
		globals:     make([]*objects.Object, runtime.GlobalsSize),
		symbolTable: compiler.NewSymbolTable(),
	}
	for idx, fn := range objects.Builtins {
		repl.symbolTable.DefineBuiltin(idx, fn.Name)
	}
	var obj objects.Object
	obj = &objects.BuiltinFunction{Value: func(args ...objects.Object) (objects.Object, error) {
		for _, arg := range args {
			// TODO: replace with param callback instead of fmt.Println
			if str, ok := arg.(*objects.String); ok {
				if print != nil {
					print(str.Value)
				} else {
					fmt.Println(str.Value)
				}
			} else {
				if print != nil {
					print(arg.String())
				} else {
					fmt.Println(arg.String())
				}
			}
		}

		return nil, nil
	}}
	symbol := repl.symbolTable.Define("print")
	repl.globals[symbol.Index] = &obj

	// for k, v := range globals {
	// 	vv, err := FromInterface(v)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	symbol = repl.symbolTable.Define(k)
	// 	repl.globals[symbol.Index] = &vv
	// }

	return repl
}

func (repl *REPL) Evaluate(line string, constants []objects.Object) ([]objects.Object, error) {
	file, err := parser.ParseFile(repl.fileSet.AddFile("repl", -1, len(line)), []byte(line), nil)
	if err != nil {
		return constants, err
	}
	file = addPrints(file)
	c := compiler.NewCompiler(repl.symbolTable, constants, nil, nil)
	if err := c.Compile(file); err != nil {
		return constants, err
	}
	bytecode := c.Bytecode()
	machine := runtime.NewVM(bytecode, repl.globals)
	if err := machine.Run(); err != nil {
		return constants, err
	}
	return bytecode.Constants, nil
}

func (repl *REPL) Run(in io.Reader, out io.Writer, globals map[string]interface{}) {
	stdin := bufio.NewScanner(in)
	for k, v := range globals {
		vv, err := FromInterface(v)
		if err != nil {
			panic(err)
		}
		symbol := repl.symbolTable.Define(k)
		repl.globals[symbol.Index] = &vv
	}
	var constants []objects.Object
	for {
		scanned := stdin.Scan()
		if !scanned {
			return
		}

		line := stdin.Text()

		//_, _ = fmt.Fprintf(out, "%s%s\n", repl.prompt, line)

		var err error
		constants, err = repl.Evaluate(line, constants)
		if err != nil {
			_, _ = fmt.Fprintf(out, "error: %s\n", err.Error())
			continue
		}
	}
}

func addPrints(file *ast.File) *ast.File {
	var stmts []ast.Stmt
	for _, s := range file.Stmts {
		switch s := s.(type) {
		case *ast.ExprStmt:
			stmts = append(stmts, &ast.ExprStmt{
				Expr: &ast.CallExpr{
					Func: &ast.Ident{
						Name: "print",
					},
					Args: []ast.Expr{s.Expr},
				},
			})

		case *ast.AssignStmt:
			stmts = append(stmts, s)

			stmts = append(stmts, &ast.ExprStmt{
				Expr: &ast.CallExpr{
					Func: &ast.Ident{
						Name: "print",
					},
					Args: s.LHS,
				},
			})

		default:
			stmts = append(stmts, s)
		}
	}

	return &ast.File{
		InputFile: file.InputFile,
		Stmts:     stmts,
	}
}

func FromInterface(v interface{}) (objects.Object, error) {
	switch v := v.(type) {
	case nil:
		return objects.UndefinedValue, nil
	case string:
		return &objects.String{Value: v}, nil
	case int64:
		return &objects.Int{Value: v}, nil
	case int:
		return &objects.Int{Value: int64(v)}, nil
	case bool:
		if v {
			return objects.TrueValue, nil
		}
		return objects.FalseValue, nil
	case rune:
		return &objects.Char{Value: v}, nil
	case byte:
		return &objects.Char{Value: rune(v)}, nil
	case float64:
		return &objects.Float{Value: v}, nil
	case []byte:
		return &objects.Bytes{Value: v}, nil
	case error:
		return &objects.Error{Value: &objects.String{Value: v.Error()}}, nil
	case map[string]objects.Object:
		return &objects.Map{Value: v}, nil
	case map[string]interface{}:
		kv := make(map[string]objects.Object)
		for vk, vv := range v {
			vo, err := FromInterface(vv)
			if err != nil {
				return nil, err
			}
			kv[vk] = vo
		}
		return &objects.Map{Value: kv}, nil
	case []objects.Object:
		return &objects.Array{Value: v}, nil
	case []interface{}:
		arr := make([]objects.Object, len(v), len(v))
		for i, e := range v {
			vo, err := FromInterface(e)
			if err != nil {
				return nil, err
			}

			arr[i] = vo
		}
		return &objects.Array{Value: arr}, nil
	case time.Time:
		return &objects.Time{Value: v}, nil
	case objects.Object:
		return v, nil
	default:
		rv := reflect.ValueOf(v)
		if (rv.Kind() == reflect.Ptr && rv.Elem().Kind() == reflect.Struct) || rv.Kind() == reflect.Struct {
			return &Struct{Value: rv}, nil
		}
		if rv.Kind() == reflect.Slice {
			arr := make([]objects.Object, rv.Len(), rv.Len())
			for i := 0; i < rv.Len(); i++ {
				vo, err := FromInterface(rv.Index(i).Interface())
				if err != nil {
					return nil, err
				}

				arr[i] = vo
			}
			return &objects.Array{Value: arr}, nil
		}
	}

	return nil, fmt.Errorf("unsupported value type: %T", v)
}

func ToInterface(o objects.Object) (res interface{}) {
	switch o := o.(type) {
	case *objects.Int:
		res = o.Value
	case *objects.String:
		res = o.Value
	case *objects.Float:
		res = o.Value
	case *objects.Bool:
		res = o == objects.TrueValue
	case *objects.Char:
		res = o.Value
	case *objects.Bytes:
		res = o.Value
	case *objects.Array:
		res = make([]interface{}, len(o.Value))
		for i, val := range o.Value {
			res.([]interface{})[i] = ToInterface(val)
		}
	case *objects.Map:
		res = make(map[string]interface{})
		for key, v := range o.Value {
			res.(map[string]interface{})[key] = ToInterface(v)
		}
	case objects.Object:
		return o
	}

	return
}

type Struct struct {
	Value reflect.Value
}

func (o *Struct) TypeName() string {
	return o.Value.Type().Name()
}

func (o *Struct) String() string {
	return fmt.Sprintf("%#v", o.Value.Interface())
}

func (o *Struct) BinaryOp(op token.Token, rhs objects.Object) (objects.Object, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (o *Struct) IsFalsy() bool {
	return false
}

func (o *Struct) Equals(another objects.Object) bool {
	obj, ok := another.(*Struct)
	if ok {
		return obj.Value.Interface() == o.Value.Interface()
	}
	return false
}

func (o *Struct) Copy() objects.Object {
	return o
}

func (o *Struct) IndexGet(index objects.Object) (objects.Object, error) {
	key, ok := index.(*objects.String)
	if ok {
		v := o.Value
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		field := v.FieldByName(key.Value)
		if field.IsValid() {
			return FromInterface(field.Interface())
		}
		method := o.Value.MethodByName(key.Value)
		if method.IsValid() {
			// TODO: create new Method type Object?
			return &objects.BuiltinFunction{Value: func(args ...objects.Object) (objects.Object, error) {
				var in []reflect.Value
				for _, arg := range args {
					in = append(in, reflect.ValueOf(ToInterface(arg)))
				}
				retVals := method.Call(in)
				if len(retVals) == 0 {
					return nil, nil
				}
				// assuming up to 2 return values, one being an error
				var retVal reflect.Value
				errorInterface := reflect.TypeOf((*error)(nil)).Elem()
				for _, v := range retVals {
					if v.Type().Implements(errorInterface) {
						if !v.IsNil() {
							return nil, v.Interface().(error)
						}
					} else {
						retVal = v
					}
				}
				return FromInterface(retVal.Interface())
			}}, nil
		}

	}
	return &objects.Undefined{}, nil
}
