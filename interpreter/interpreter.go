package interpreter

import (
	"os"

	"strconv"

	"strings"

	"github.com/k0kubun/pp"
	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/builtin"
	"github.com/tzmfreedom/goland/compiler"
)

type Interpreter struct {
	Context *Context
	Extra   map[string]interface{}
}

func NewInterpreter(classTypeMap *ast.ClassMap) *Interpreter {
	interpreter := &Interpreter{
		Context: NewContext(),
		Extra: map[string]interface{}{
			"stdout": os.Stdout,
			"stderr": os.Stderr,
			"errors": []*builtin.TestError{},
		},
	}
	interpreter.Extra["interpreter"] = interpreter
	interpreter.Context.ClassTypes = classTypeMap
	return interpreter
}

func NewInterpreterWithBuiltin(classTypes []*ast.ClassType) *Interpreter {
	interpreter := NewInterpreter(builtin.PrimitiveClassMap())
	for _, classType := range classTypes {
		interpreter.Context.ClassTypes.Set(classType.Name, classType)
	}
	interpreter.Context.NameSpaces = builtin.GetNameSpaceStore()
	return interpreter
}

func (v *Interpreter) LoadStaticField() {
	v.Context.StaticField = NewStaticFieldMap()
	for className, classType := range v.Context.ClassTypes.Data {
		objectMap := ast.NewObjectMap()
		if classType.StaticFields != nil {
			for _, f := range classType.StaticFields.Data {
				val, _ := f.Expression.Accept(v)
				objectMap.Set(f.Name, val.(*ast.Object))
			}
		}
		v.Context.StaticField.Set("_", className, objectMap)
	}
}

func (v *Interpreter) VisitClassDeclaration(n *ast.ClassDeclaration) (interface{}, error) {
	panic("not pass")
	return ast.VisitClassDeclaration(v, n)
}

func (v *Interpreter) VisitModifier(n *ast.Modifier) (interface{}, error) {
	panic("not pass")
	return ast.VisitModifier(v, n)
}

func (v *Interpreter) VisitAnnotation(n *ast.Annotation) (interface{}, error) {
	panic("not pass")
	return ast.VisitAnnotation(v, n)
}

func (v *Interpreter) VisitInterfaceDeclaration(n *ast.InterfaceDeclaration) (interface{}, error) {
	panic("not pass")
	return ast.VisitInterfaceDeclaration(v, n)
}

func (v *Interpreter) VisitIntegerLiteral(n *ast.IntegerLiteral) (interface{}, error) {
	return builtin.NewInteger(n.Value), nil
}

func (v *Interpreter) VisitParameter(n *ast.Parameter) (interface{}, error) {
	panic("not pass")
	return ast.VisitParameter(v, n)
}

func (v *Interpreter) VisitArrayAccess(n *ast.ArrayAccess) (interface{}, error) {
	r, err := n.Receiver.Accept(v)
	if err != nil {
		return nil, err
	}
	k, err := n.Key.Accept(v)
	if err != nil {
		return nil, err
	}
	receiver := r.(*ast.Object)
	key := k.(*ast.Object)
	if receiver.ClassType.Name == "List" {
		records := receiver.Extra["records"].([]*ast.Object)
		return records[key.IntegerValue()], nil
	}

	records := receiver.Extra["values"].(map[string]*ast.Object)
	return records[key.StringValue()], nil
}

func (v *Interpreter) VisitBooleanLiteral(n *ast.BooleanLiteral) (interface{}, error) {
	return builtin.NewBoolean(n.Value), nil
}

func (v *Interpreter) VisitBreak(n *ast.Break) (interface{}, error) {
	return builtin.Break, nil
}

func (v *Interpreter) VisitContinue(n *ast.Continue) (interface{}, error) {
	return builtin.Continue, nil
}

func (v *Interpreter) VisitDml(n *ast.Dml) (interface{}, error) {
	o, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	var records []*ast.Object
	obj := o.(*ast.Object)
	// TODO: check SObject class
	if obj.ClassType.Name == "List" {
		records = obj.Extra["records"].([]*ast.Object)
	} else {
		records = []*ast.Object{obj}
	}
	sObjectType := records[0].ClassType.Name
	builtin.DatabaseDriver.Execute(n.Type, sObjectType, records, n.UpsertKey)
	return nil, nil
}

func (v *Interpreter) VisitDoubleLiteral(n *ast.DoubleLiteral) (interface{}, error) {
	return builtin.NewDouble(n.Value), nil
}

func (v *Interpreter) VisitFieldDeclaration(n *ast.FieldDeclaration) (interface{}, error) {
	panic("not pass")
	return ast.VisitFieldDeclaration(v, n)
}

func (v *Interpreter) VisitTry(n *ast.Try) (interface{}, error) {
	res, err := n.Block.Accept(v)
	if err != nil {
		return nil, err
	}
	if res != nil {
		switch obj := res.(*ast.Object); obj.ClassType {
		case builtin.ReturnType, builtin.BreakType, builtin.ContinueType:
			return obj, nil
		case builtin.RaiseType:
			// TODO: implement
			for _, catch := range n.CatchClause {
				raiseValue := obj.Value().(*ast.Object)
				if builtin.Equals(catch.Type, raiseValue.ClassType) {
					v.Context.Env.Define(catch.Identifier, raiseValue)
					_, err := catch.Accept(v)
					if err != nil {
						return nil, err
					}
					// TODO: return, raise impl
				}
			}
		}
	}
	if n.FinallyBlock != nil {
		if _, err = n.FinallyBlock.Accept(v); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (v *Interpreter) VisitCatch(n *ast.Catch) (interface{}, error) {
	return n.Block.Accept(v)
}

func (v *Interpreter) VisitFinally(n *ast.Finally) (interface{}, error) {
	return ast.VisitFinally(v, n)
}

func (v *Interpreter) VisitFor(n *ast.For) (interface{}, error) {
	return v.NewEnv(func() (interface{}, error) {
		switch control := n.Control.(type) {
		case *ast.ForControl:
			for _, forInit := range control.ForInit {
				_, err := forInit.Accept(v)
				if err != nil {
					return nil, err
				}
			}
			for {
				res, err := control.Expression.Accept(v)
				if err != nil {
					return nil, err
				}
				if res.(*ast.Object).BoolValue() {
					res, err = n.Statements.Accept(v)
					if res != nil {
						switch obj := res.(*ast.Object); obj.ClassType {
						case builtin.BreakType:
							return nil, nil
						case builtin.ContinueType:
							for _, stmt := range control.ForUpdate {
								stmt.Accept(v)
							}
							continue
						case builtin.ReturnType, builtin.RaiseType:
							return obj, nil
						}
					}
					for _, stmt := range control.ForUpdate {
						stmt.Accept(v)
					}
				} else {
					break
				}
			}
		case *ast.EnhancedForControl:
			_, err := control.Accept(v)
			if err != nil {
				return nil, err
			}
			iterator, _ := control.Expression.Accept(v)
			records := iterator.(*ast.Object).Extra["records"].([]*ast.Object)
			for _, record := range records {
				v.Context.Env.Define(control.VariableDeclaratorId, record)
				res, err := n.Statements.Accept(v)
				if err != nil {
					return nil, err
				}
				if res != nil {
					switch obj := res.(*ast.Object); obj.ClassType {
					case builtin.BreakType:
						return nil, nil
					case builtin.ContinueType:
						continue
					case builtin.ReturnType, builtin.RaiseType:
						return obj, nil
					}
				}
			}
		}
		return nil, nil
	})
}

func (v *Interpreter) VisitForControl(n *ast.ForControl) (interface{}, error) {
	return ast.VisitForControl(v, n)
}

func (v *Interpreter) VisitEnhancedForControl(n *ast.EnhancedForControl) (interface{}, error) {
	return ast.VisitEnhancedForControl(v, n)
}

func (v *Interpreter) VisitIf(n *ast.If) (interface{}, error) {
	res, err := n.Condition.Accept(v)
	if err != nil {
		return nil, err
	}
	if res.(*ast.Object).BoolValue() {
		return n.IfStatement.Accept(v)
	} else if n.ElseStatement != nil {
		return n.ElseStatement.Accept(v)
	}
	return nil, nil
}

func (v *Interpreter) VisitMethodDeclaration(n *ast.MethodDeclaration) (interface{}, error) {
	panic("not pass")
	return ast.VisitMethodDeclaration(v, n)
}

func (v *Interpreter) VisitMethodInvocation(n *ast.MethodInvocation) (interface{}, error) {
	Publish("method_start", v.Context, n)
	var receiver interface{}
	var m *ast.Method
	var err error

	evaluated := make([]*ast.Object, len(n.Parameters))
	for i, p := range n.Parameters {
		obj, err := p.Accept(v)
		if err != nil {
			return nil, err
		}
		evaluated[i] = obj.(*ast.Object)
	}
	switch exp := n.NameOrExpression.(type) {
	case *ast.FieldAccess:
		r, err := exp.Expression.Accept(v)
		if err != nil {
			return nil, err
		}
		receiver = r
		// TODO: extend
		_, m, err = FindInstanceMethod(receiver.(*ast.Object), exp.FieldName, evaluated, compiler.MODIFIER_ALL_OK)
		if err != nil {
			panic("not found")
		}
	case *ast.Name:
		// TODO: implement
		if exp.Value[0] == "_Debugger" {
			switch exp.Value[1] {
			case "run":
				Debugger.Debug(v.Context, n)
			case "debug":
				for _, p := range n.Parameters {
					pp.Println(p)
				}
			}
			return nil, nil
		}
		resolver := NewTypeResolver(v.Context)
		receiver, m, err = resolver.ResolveMethod(exp.Value, evaluated)
		if err != nil {
			return nil, err
		}
	}
	prevClass := v.Context.CurrentClass
	switch typedReceiver := receiver.(type) {
	case *ast.Object:
		v.Context.CurrentClass = typedReceiver.ClassType
	case *ast.ClassType:
		v.Context.CurrentClass = typedReceiver
	}
	defer func() {
		v.Context.CurrentClass = prevClass
	}()

	if m.NativeFunction != nil {
		var r interface{}
		v.Extra["node"] = n
		switch typedReceiver := receiver.(type) {
		case *ast.Object:
			r = m.NativeFunction(typedReceiver, evaluated, v.Extra)
		case *ast.ClassType:
			r = m.NativeFunction(nil, evaluated, v.Extra)
		}
		v.Extra["node"] = nil
		Publish("method_end", v.Context, n)
		return r, nil
	}
	prev := v.Context.Env
	v.Context.Env = NewEnv(nil)
	for i, param := range m.Parameters {
		v.Context.Env.Define(param.Name, evaluated[i])
	}
	switch obj := receiver.(type) {
	case *ast.Object:
		v.Context.Env.Define("this", obj)
	}
	r, err := m.Statements.Accept(v)
	Publish("method_end", v.Context, n)
	if err != nil {
		return nil, err
	}
	v.Context.Env = prev

	if r != nil {
		obj := r.(*ast.Object)
		switch obj.ClassType {
		case builtin.ReturnType:
			return obj.Value(), nil
		case builtin.RaiseType:
			return obj, nil
		}
	}
	return nil, nil
}

func (v *Interpreter) VisitNew(n *ast.New) (interface{}, error) {
	classType := n.Type
	newObj := &ast.Object{
		ClassType:      classType,
		InstanceFields: ast.NewObjectMap(),
		Extra:          map[string]interface{}{},
	}
	for _, f := range classType.InstanceFields.Data {
		if f.Expression == nil {
			newObj.InstanceFields.Set(f.Name, builtin.Null)
		} else {
			r, _ := f.Expression.Accept(v)
			newObj.InstanceFields.Set(f.Name, r.(*ast.Object))
		}
	}
	typeResolver := NewTypeResolver(v.Context)
	if classType.HasConstructor() {
		evaluated := make([]*ast.Object, len(n.Parameters))
		for i, p := range n.Parameters {
			r, err := p.Accept(v)
			if err != nil {
				return nil, err
			}
			evaluated[i] = r.(*ast.Object)
		}
		_, constructor, err := typeResolver.SearchConstructor(classType, evaluated)
		if err != nil {
			return nil, err
		}
		if constructor == nil {
			panic("constructor is not found")
		}

		if constructor.NativeFunction != nil {
			constructor.NativeFunction(newObj, evaluated, v.Extra)
		} else {
			prev := v.Context.Env
			v.Context.Env = NewEnv(nil)
			for i, param := range constructor.Parameters {
				v.Context.Env.Define(param.Name, evaluated[i])
			}
			v.Context.Env.Define("this", newObj)
			constructor.Statements.Accept(v)
			v.Context.Env = prev
		}
	}
	if classType.Name == "List" {
		newObj.Extra["records"] = []*ast.Object{}
		if n.Init != nil {
			if len(n.Init.Records) != 0 {
				records := make([]*ast.Object, len(n.Init.Records))
				for i, r := range n.Init.Records {
					initRecord, err := r.Accept(v)
					if err != nil {
						return nil, err
					}
					records[i] = initRecord.(*ast.Object)
				}
				newObj.Extra["records"] = records
			}
			// TODO: multi dimmension initialize
			if len(n.Init.Sizes) > 0 {
				size, err := n.Init.Sizes[0].Accept(v)
				if err != nil {
					return nil, err
				}
				if s, ok := size.(*ast.Object); ok && s.ClassType == builtin.IntegerType {
					records := newObj.Extra["records"].([]*ast.Object)
					recordSize := len(records)
					remain := s.IntegerValue() - recordSize
					if remain > 0 {
						for i := 0; i < remain; i++ {
							records = append(records, builtin.Null)
						}
						newObj.Extra["records"] = records
					}
				}
			}
		}
	}
	if classType.Name == "Map" {
		if n.Init == nil {
			newObj.Extra["values"] = map[string]*ast.Object{}
		} else {
			values := map[string]*ast.Object{}
			for key, value := range n.Init.Values {
				mapValue, err := value.Accept(v)
				if err != nil {
					return nil, err
				}
				mapKey, err := key.Accept(v)
				if err != nil {
					return nil, err
				}
				values[mapKey.(*ast.Object).StringValue()] = mapValue.(*ast.Object)
			}
			newObj.Extra["values"] = values
		}
	}
	return newObj, nil
}

func (v *Interpreter) VisitNullLiteral(n *ast.NullLiteral) (interface{}, error) {
	return builtin.Null, nil
}

func (v *Interpreter) VisitUnaryOperator(n *ast.UnaryOperator) (interface{}, error) {
	if n.Op == "++" {
		name := n.Expression.(*ast.Name)
		l, _ := name.Accept(v)
		exp := builtin.NewInteger(l.(*ast.Object).IntegerValue() + 1)
		// TODO: implement
		v.Context.Env.Update(name.Value[0], exp)
		return exp, nil
	}
	if n.Op == "--" {
		name := n.Expression.(*ast.Name)
		l, _ := name.Accept(v)
		exp := builtin.NewInteger(l.(*ast.Object).IntegerValue() - 1)
		// TODO: implement
		v.Context.Env.Update(name.Value[0], exp)
		return exp, nil
	}
	panic("not pass")
	return nil, nil
}

func (v *Interpreter) VisitBinaryOperator(n *ast.BinaryOperator) (interface{}, error) {
	left, err := n.Left.Accept(v)
	if err != nil {
		return nil, err
	}
	right, err := n.Right.Accept(v)
	if err != nil {
		return nil, err
	}

	lObj := left.(*ast.Object)
	lType := lObj.ClassType
	rObj := right.(*ast.Object)
	rType := rObj.ClassType

	switch n.Op {
	case "+":
		if lType == builtin.IntegerType {
			l := lObj.IntegerValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				return builtin.NewInteger(l + r), nil
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				return builtin.NewDouble(float64(l) + r), nil
			}
		} else if lType == builtin.DoubleType {
			l := lObj.DoubleValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				return builtin.NewDouble(l + float64(r)), nil
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				return builtin.NewDouble(r + l), nil
			}
		} else if lType == builtin.StringType {
			l := lObj.StringValue()
			r := rObj.StringValue()
			return builtin.NewString(l + r), nil
		}
		panic("type error")
	case "-":
		if lType == builtin.IntegerType {
			l := lObj.IntegerValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				return builtin.NewInteger(l - r), nil
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				return builtin.NewDouble(float64(l) - r), nil
			}
		} else if lType == builtin.DoubleType {
			l := lObj.DoubleValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				return builtin.NewDouble(l - float64(r)), nil
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				return builtin.NewDouble(r - l), nil
			}
		}
		panic("type error")
	case "*":
		if lType == builtin.IntegerType {
			l := lObj.IntegerValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				return builtin.NewInteger(l * r), nil
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				return builtin.NewDouble(float64(l) * r), nil
			}
		} else if lType == builtin.DoubleType {
			l := lObj.DoubleValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				return builtin.NewDouble(l * float64(r)), nil
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				return builtin.NewDouble(r * l), nil
			}
		}
		panic("type error")
	case "/":
		if lType == builtin.IntegerType {
			l := lObj.IntegerValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				return builtin.NewInteger(l / r), nil
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				return builtin.NewDouble(float64(l) / r), nil
			}
		} else if lType == builtin.DoubleType {
			l := lObj.DoubleValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				return builtin.NewDouble(l / float64(r)), nil
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				return builtin.NewDouble(r / l), nil
			}
		}
		panic("type error")
	case "%":
		l := lObj.IntegerValue()
		r := rObj.IntegerValue()
		return builtin.NewInteger(l % r), nil
	case "<":
		if lType == builtin.IntegerType {
			l := lObj.IntegerValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				return builtin.NewBoolean(l < r), nil
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				return builtin.NewBoolean(float64(l) < r), nil
			}
		} else if lType == builtin.DoubleType {
			l := lObj.DoubleValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				return builtin.NewBoolean(l < float64(r)), nil
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				return builtin.NewBoolean(r < l), nil
			}
		}
		panic("type error")
	case ">":
		if lType == builtin.IntegerType {
			l := lObj.IntegerValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				return builtin.NewBoolean(l > r), nil
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				return builtin.NewBoolean(float64(l) > r), nil
			}
		} else if lType == builtin.DoubleType {
			l := lObj.DoubleValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				return builtin.NewBoolean(l > float64(r)), nil
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				return builtin.NewBoolean(r > l), nil
			}
		}
		panic("type error")
	case "<=":
		if lType == builtin.IntegerType {
			l := lObj.IntegerValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				return builtin.NewBoolean(l <= r), nil
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				return builtin.NewBoolean(float64(l) <= r), nil
			}
		} else if lType == builtin.DoubleType {
			l := lObj.DoubleValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				return builtin.NewBoolean(l <= float64(r)), nil
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				return builtin.NewBoolean(r <= l), nil
			}
		}
		panic("type error")
	case ">=":
		if lType == builtin.IntegerType {
			l := lObj.IntegerValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				return builtin.NewBoolean(l >= r), nil
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				return builtin.NewBoolean(float64(l) >= r), nil
			}
		} else if lType == builtin.DoubleType {
			l := lObj.DoubleValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				return builtin.NewBoolean(l >= float64(r)), nil
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				return builtin.NewBoolean(r >= l), nil
			}
		}
		panic("type error")
	case "==":
		if lType == builtin.IntegerType {
			l := lObj.IntegerValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				return builtin.NewBoolean(l == r), nil
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				return builtin.NewBoolean(float64(l) == r), nil
			}
		} else if lType == builtin.DoubleType {
			l := lObj.DoubleValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				return builtin.NewBoolean(l == float64(r)), nil
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				return builtin.NewBoolean(r == l), nil
			}
		} else if lType == builtin.StringType {
			l := lObj.StringValue()
			r := rObj.StringValue()
			return builtin.NewBoolean(l == r), nil
		}
		panic("type error")
	case "!=":
		if lType == builtin.IntegerType {
			l := lObj.IntegerValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				return builtin.NewBoolean(l != r), nil
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				return builtin.NewBoolean(float64(l) != r), nil
			}
		} else if lType == builtin.DoubleType {
			l := lObj.DoubleValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				return builtin.NewBoolean(l != float64(r)), nil
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				return builtin.NewBoolean(r != l), nil
			}
		} else if lType == builtin.StringType {
			l := lObj.StringValue()
			r := rObj.StringValue()
			return builtin.NewBoolean(l != r), nil
		}
		panic("type error")
	case "=":
		// TODO: implement
		switch t := n.Left.(type) {
		case *ast.Name:
			resolver := NewTypeResolver(v.Context)
			resolver.SetVariable(t.Value, rObj)
		case *ast.FieldAccess:
			exp, err := t.Expression.Accept(v)
			if err != nil {
				return nil, err
			}
			exp.(*ast.Object).InstanceFields.Set(t.FieldName, rObj)
		case *ast.ArrayAccess:
			k, err := t.Key.Accept(v)
			if err != nil {
				return nil, err
			}
			key := k.(*ast.Object)
			r, err := t.Receiver.Accept(v)
			if err != nil {
				return nil, err
			}
			receiver := r.(*ast.Object)
			if receiver.ClassType.Name == "List" {
				receiver.Extra["records"].([]*ast.Object)[key.IntegerValue()] = rObj
			}
			if receiver.ClassType.Name == "Map" {
				receiver.Extra["values"].(map[string]*ast.Object)[key.StringValue()] = rObj
			}
			// TODO: implment set type
		}
		return rObj, nil
	case "+=":
		// TODO: implement
		var value *ast.Object
		if lType == builtin.IntegerType {
			l := lObj.IntegerValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				value = builtin.NewInteger(l + r)
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				value = builtin.NewDouble(float64(l) + r)
			}
		} else if lType == builtin.DoubleType {
			l := lObj.DoubleValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				value = builtin.NewDouble(l + float64(r))
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				value = builtin.NewDouble(r + l)
			}
		} else if lType == builtin.StringType {
			l := lObj.StringValue()
			r := rObj.StringValue()
			value = builtin.NewString(l + r)
		}

		switch t := n.Left.(type) {
		case *ast.Name:
			resolver := NewTypeResolver(v.Context)
			resolver.SetVariable(t.Value, value)
		case *ast.FieldAccess:
			exp, err := t.Expression.Accept(v)
			if err != nil {
				return nil, err
			}
			exp.(*ast.Object).InstanceFields.Set(t.FieldName, value)
		}
		return value, nil
	case "-=":
		var value *ast.Object
		if lType == builtin.IntegerType {
			l := lObj.IntegerValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				value = builtin.NewInteger(l - r)
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				value = builtin.NewDouble(float64(l) - r)
			}
		} else if lType == builtin.DoubleType {
			l := lObj.DoubleValue()
			if rType == builtin.IntegerType {
				r := rObj.IntegerValue()
				value = builtin.NewDouble(l - float64(r))
			}
			if rType == builtin.DoubleType {
				r := rObj.DoubleValue()
				value = builtin.NewDouble(r - l)
			}
		}

		switch t := n.Left.(type) {
		case *ast.Name:
			resolver := NewTypeResolver(v.Context)
			resolver.SetVariable(t.Value, value)
		case *ast.FieldAccess:
			exp, err := t.Expression.Accept(v)
			if err != nil {
				return nil, err
			}
			exp.(*ast.Object).InstanceFields.Set(t.FieldName, value)
		}
		return value, nil
	}
	return nil, nil
}

func (v *Interpreter) VisitReturn(n *ast.Return) (interface{}, error) {
	if n.Expression != nil {
		res, err := n.Expression.Accept(v)
		if err != nil {
			return nil, err
		}
		return builtin.CreateReturn(res.(*ast.Object)), nil
	}
	return builtin.CreateReturn(nil), nil
}

func (v *Interpreter) VisitThrow(n *ast.Throw) (interface{}, error) {
	res, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	return builtin.CreateRaise(res.(*ast.Object)), nil
}

func (v *Interpreter) VisitSoql(n *ast.Soql) (interface{}, error) {
	executor := &SoqlExecutor{}
	objects, err := executor.Execute(n, v)
	if err != nil {
		return nil, err
	}
	return objects, nil
}

func (v *Interpreter) VisitSosl(n *ast.Sosl) (interface{}, error) {
	return ast.VisitSosl(v, n)
}

func (v *Interpreter) VisitStringLiteral(n *ast.StringLiteral) (interface{}, error) {
	return builtin.NewString(n.Value), nil
}

func (v *Interpreter) VisitSwitch(n *ast.Switch) (interface{}, error) {
	exp, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	expObj := exp.(*ast.Object)
	for _, when := range n.WhenStatements {
		for _, cond := range when.Condition {
			c, err := cond.Accept(v)
			if err != nil {
				return nil, err
			}
			condObj := c.(*ast.Object)
			if v.Equals(expObj, condObj) {
				return when.Statements.Accept(v)
			}
		}
	}
	if n.ElseStatement != nil {
		return n.ElseStatement.Accept(v)
	}
	return nil, nil
}

func (v *Interpreter) VisitTrigger(n *ast.Trigger) (interface{}, error) {
	return ast.VisitTrigger(v, n)
}

func (v *Interpreter) VisitTriggerTiming(n *ast.TriggerTiming) (interface{}, error) {
	panic("not pass")
	return ast.VisitTriggerTiming(v, n)
}

func (v *Interpreter) VisitVariableDeclaration(n *ast.VariableDeclaration) (interface{}, error) {
	for _, declarator := range n.Declarators {
		if declarator.Expression != nil {
			val, err := declarator.Expression.Accept(v)
			if err != nil {
				panic(err)
			}
			v.Context.Env.Define(declarator.Name, val.(*ast.Object))
		} else {
			v.Context.Env.Define(declarator.Name, builtin.Null)
		}
	}
	return nil, nil
}

func (v *Interpreter) VisitVariableDeclarator(n *ast.VariableDeclarator) (interface{}, error) {
	panic("not pass")
	return nil, nil
}

func (v *Interpreter) VisitWhen(n *ast.When) (interface{}, error) {
	return ast.VisitWhen(v, n)
}

func (v *Interpreter) VisitWhenType(n *ast.WhenType) (interface{}, error) {
	return ast.VisitWhenType(v, n)
}

func (v *Interpreter) VisitWhile(n *ast.While) (interface{}, error) {
	for {
		c, err := n.Condition.Accept(v)
		if err != nil {
			return nil, err
		}
		if !c.(*ast.Object).BoolValue() {
			break
		}
		res, err := n.Statements.Accept(v)
		if res != nil {
			obj := res.(*ast.Object)
			switch obj.ClassType {
			case builtin.ReturnType, builtin.RaiseType:
				return obj, nil
			case builtin.BreakType:
				return nil, nil
			case builtin.ContinueType:
				continue
			}
		}
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (v *Interpreter) VisitNothingStatement(n *ast.NothingStatement) (interface{}, error) {
	return ast.VisitNothingStatement(v, n)
}

func (v *Interpreter) VisitCastExpression(n *ast.CastExpression) (interface{}, error) {
	return ast.VisitCastExpression(v, n)
}

func (v *Interpreter) VisitFieldAccess(n *ast.FieldAccess) (interface{}, error) {
	r, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	f, _ := r.(*ast.Object).InstanceFields.Get(n.FieldName)
	return f, nil
}

func (v *Interpreter) VisitType(n *ast.TypeRef) (interface{}, error) {
	return ast.VisitType(v, n)
}

func (v *Interpreter) VisitBlock(n *ast.Block) (interface{}, error) {
	prevEnv := v.Context.Env
	v.Context.Env = NewEnv(prevEnv)
	defer func() {
		v.Context.Env = prevEnv
	}()
	for _, stmt := range n.Statements {
		Publish("line", v.Context, stmt)
		res, err := stmt.Accept(v)
		if err != nil {
			return nil, err
		}
		if res != nil {
			obj := res.(*ast.Object)
			switch obj.ClassType {
			case builtin.ReturnType, builtin.BreakType, builtin.ContinueType, builtin.RaiseType:
				return obj, nil
			}
		}
	}
	return nil, nil
}

func (v *Interpreter) VisitGetterSetter(n *ast.GetterSetter) (interface{}, error) {
	return ast.VisitGetterSetter(v, n)
}

func (v *Interpreter) VisitPropertyDeclaration(n *ast.PropertyDeclaration) (interface{}, error) {
	return ast.VisitPropertyDeclaration(v, n)
}

func (v *Interpreter) VisitArrayInitializer(n *ast.ArrayInitializer) (interface{}, error) {
	return ast.VisitArrayInitializer(v, n)
}

func (v *Interpreter) VisitArrayCreator(n *ast.ArrayCreator) (interface{}, error) {
	return ast.VisitArrayCreator(v, n)
}

func (v *Interpreter) VisitSoqlBindVariable(n *ast.SoqlBindVariable) (interface{}, error) {
	return ast.VisitSoqlBindVariable(v, n)
}

func (v *Interpreter) VisitTernalyExpression(n *ast.TernalyExpression) (interface{}, error) {
	res, err := n.Condition.Accept(v)
	if err != nil {
		return nil, err
	}
	if res.(*ast.Object).Extra["value"].(bool) {
		return n.TrueExpression.Accept(v)
	}
	return n.FalseExpression.Accept(v)
}

func (v *Interpreter) VisitMapCreator(n *ast.MapCreator) (interface{}, error) {
	return ast.VisitMapCreator(v, n)
}

func (v *Interpreter) VisitSetCreator(n *ast.SetCreator) (interface{}, error) {
	return ast.VisitSetCreator(v, n)
}

func (v *Interpreter) VisitName(n *ast.Name) (interface{}, error) {
	resolver := NewTypeResolver(v.Context)
	return resolver.ResolveVariable(n.Value)
}

func (v *Interpreter) VisitConstructorDeclaration(n *ast.ConstructorDeclaration) (interface{}, error) {
	return ast.VisitConstructorDeclaration(v, n)
}

func (v *Interpreter) NewEnv(f func() (interface{}, error)) (interface{}, error) {
	prevEnv := v.Context.Env
	v.Context.Env = NewEnv(prevEnv)
	r, err := f()
	v.Context.Env = prevEnv
	return r, err
}

func (v *Interpreter) Equals(o, other *ast.Object) bool {
	if o == builtin.Null {
		return other == builtin.Null
	}
	m, ok := o.ClassType.InstanceMethods.Get("equals")
	if !ok {
		return o.Value() == other.Value()
	}
	method := m[0]
	if method.NativeFunction != nil {
		bObj := method.NativeFunction(o, []*ast.Object{other}, nil).(*ast.Object)
		return bObj.BoolValue()
	}
	prev := v.Context.Env
	v.Context.Env = NewEnv(nil)
	v.Context.Env.Define(method.Parameters[0].Name, other)
	v.Context.Env.Define("this", o)
	r, _ := method.Statements.Accept(v) // TODO
	v.Context.Env = prev
	return r.(*ast.Object).BoolValue()
}

// @return controller object, pageref object, error
func (i *Interpreter) BindAndRun(name, method string, params map[string][]string, state map[string]interface{}) (*ast.Object, *ast.Object, error) {
	classType, ok := i.Context.ClassTypes.Get(name)
	if !ok {
		panic("no exist class: " + name)
	}
	newNode := &ast.New{
		Type:       classType,
		Parameters: []ast.Node{},
		Init:       nil,
	}
	r, err := i.VisitNew(newNode)
	if err != nil {
		return nil, nil, err
	}
	controller := r.(*ast.Object)
	applyViewState(controller, state)
	for k, v := range params {
		if k == "__action" || k == "__viewstate" {
			continue
		}
		names := strings.Split(k, ".")
		value := controller
		var ok bool
		for _, name := range names[:len(names)-1] {
			value, ok = value.InstanceFields.Get(name)
			if !ok {
				panic("not found: " + name)
			}
		}
		fieldName := names[len(names)-1]
		if _, ok := value.ClassType.InstanceFields.Get(fieldName); !ok {
			panic("not found: " + k)
		}
		switch value.ClassType {
		case builtin.StringType:
			value.InstanceFields.Set(fieldName, builtin.NewString(v[0]))
		case builtin.IntegerType:
			intValue, err := strconv.Atoi(v[0])
			if err != nil {
				panic(err)
			}
			value.InstanceFields.Set(fieldName, builtin.NewInteger(intValue))
		case builtin.BooleanType:
			value.InstanceFields.Set(fieldName, builtin.NewBoolean(v[0] == "true"))
		case builtin.DoubleType:
			floatValue, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				panic(err)
			}
			value.InstanceFields.Set(fieldName, builtin.NewDouble(floatValue))
		}
	}
	if method == "" {
		return controller, nil, nil
	}
	i.Context.StaticField.Add("_", classType.Name, "_context", controller)
	retValue, err := i.VisitMethodInvocation(&ast.MethodInvocation{
		NameOrExpression: &ast.Name{Value: []string{classType.Name, "_context", method}},
		Parameters:       []ast.Node{},
	})
	if err != nil {
		return nil, nil, err
	}
	return controller, retValue.(*ast.Object), nil
}

func applyViewState(o *ast.Object, params map[string]interface{}) {
	if params == nil {
		return
	}
	for k, v := range params {
		if v == nil {
			o.InstanceFields.Set(k, builtin.Null)
			continue
		}
		switch value := v.(type) {
		case string:
			o.InstanceFields.Set(k, builtin.NewString(value))
		case int:
			o.InstanceFields.Set(k, builtin.NewInteger(value))
		case bool:
			o.InstanceFields.Set(k, builtin.NewBoolean(value))
		case float64:
			o.InstanceFields.Set(k, builtin.NewDouble(value))
		case map[string]interface{}:
			field, ok := o.InstanceFields.Get(k)
			if !ok {
				panic("field not found: " + k)
			}
			applyViewState(field, value)
		}
	}
}
