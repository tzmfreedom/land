package interpreter

import (
	"io"
	"os"

	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/builtin"
	"github.com/tzmfreedom/goland/compiler"
)

type Interpreter struct {
	Context *Context
	Stdout  io.Writer
	Stderr  io.Writer
	Extra   map[string]interface{}
}

func NewInterpreter(classTypeMap *builtin.ClassMap) *Interpreter {
	interpreter := &Interpreter{
		Context: NewContext(),
		Extra: map[string]interface{}{
			"stdout": os.Stdout,
			"stderr": os.Stderr,
		},
	}
	interpreter.Context.ClassTypes = classTypeMap
	return interpreter
}

func (v *Interpreter) LoadStaticField() {
	v.Context.StaticField = NewStaticFieldMap()
	for className, classType := range v.Context.ClassTypes.Data {
		objectMap := builtin.NewObjectMap()
		if classType.StaticFields != nil {
			for _, f := range classType.StaticFields.Data {
				val, _ := f.Expression.Accept(v)
				objectMap.Set(f.Name, val.(*builtin.Object))
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
	receiver := r.(*builtin.Object)
	key := k.(*builtin.Object)
	if receiver.ClassType == builtin.ListType {
		records := receiver.Extra["records"].([]*builtin.Object)
		return records[key.IntegerValue()], nil
	}

	records := receiver.Extra["values"].(map[string]builtin.Object)
	return records[key.StringValue()], nil
}

func (v *Interpreter) VisitBooleanLiteral(n *ast.BooleanLiteral) (interface{}, error) {
	return builtin.NewBoolean(n.Value), nil
}

func (v *Interpreter) VisitBreak(n *ast.Break) (interface{}, error) {
	return &Break{}, nil
}

func (v *Interpreter) VisitContinue(n *ast.Continue) (interface{}, error) {
	return &Continue{}, nil
}

func (v *Interpreter) VisitDml(n *ast.Dml) (interface{}, error) {
	o, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	var records []*builtin.Object
	obj := o.(*builtin.Object)
	// TODO: check SObject class
	if obj.ClassType == builtin.ListType {
		records = obj.Extra["records"].([]*builtin.Object)
	} else {
		records = []*builtin.Object{obj}
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
		switch res.(type) {
		case *Raise:
			raise := res.(*Raise)
			// TODO: implement
			var c *ast.Catch
			for _, catch := range n.CatchClause {
				c = catch.(*ast.Catch)
				typeResolver := NewTypeResolver(v.Context)
				catchType, err := typeResolver.ResolveType(c.Type.(*ast.TypeRef).Name)
				if err != nil {
					return nil, err
				}
				if catchType.Equals(raise.Value.ClassType) {
					v.Context.Env.Set(c.Identifier, raise.Value)
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
			if res.(*builtin.Object).BoolValue() {
				res, err = n.Statements.Accept(v)
				if res != nil {
					switch r := res.(type) {
					case *Break:
						return nil, nil
					case *Continue:
						for _, stmt := range control.ForUpdate {
							stmt.Accept(v)
						}
						continue
					case *Return, *Raise:
						return r, nil
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
		records := iterator.(*builtin.Object).Extra["records"].([]*builtin.Object)
		for _, record := range records {
			v.Context.Env.Set(control.VariableDeclaratorId, record)
			res, err := n.Statements.Accept(v)
			if err != nil {
				return nil, err
			}
			if res != nil {
				switch r := res.(type) {
				case *Break:
					return nil, nil
				case *Continue:
					continue
				case *Return, *Raise:
					return r, nil
				}
			}
		}
	}
	return nil, nil
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
	if res.(*builtin.Object).BoolValue() {
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
	var m *builtin.Method
	var err error

	evaluated := make([]*builtin.Object, len(n.Parameters))
	for i, p := range n.Parameters {
		obj, err := p.Accept(v)
		evaluated[i] = obj.(*builtin.Object)
		if err != nil {
			return nil, err
		}
	}
	switch exp := n.NameOrExpression.(type) {
	case *ast.FieldAccess:
		r, err := exp.Expression.Accept(v)
		receiver := r.(*builtin.Object)
		if err != nil {
			return nil, err
		}
		// TODO: extend
		typeResolver := NewTypeResolver(v.Context)
		_, m, err = typeResolver.FindInstanceMethod(receiver, exp.FieldName, evaluated, compiler.MODIFIER_ALL_OK)
		if err != nil {
			panic("not found")
		}
	case *ast.Name:
		// TODO: implement
		if exp.Value[0] == "Debugger" {
			Debugger.Debug(v.Context, n)
			return nil, nil
		}
		resolver := NewTypeResolver(v.Context)
		receiver, m, err = resolver.ResolveMethod(exp.Value, evaluated)
		if err != nil {
			return nil, err
		}
	}
	if m.NativeFunction != nil {
		var r interface{}
		switch typedReceiver := receiver.(type) {
		case *builtin.Object:
			r = m.NativeFunction(typedReceiver, evaluated, v.Extra)
		case *builtin.ClassType:
			r = m.NativeFunction(nil, evaluated, v.Extra)
		}
		Publish("method_end", v.Context, n)
		return r, nil
	} else {
		prev := v.Context.Env
		v.Context.Env = NewEnv(nil)
		for i, p := range m.Parameters {
			param := p.(*ast.Parameter)
			v.Context.Env.Set(param.Name, evaluated[i])
		}
		switch obj := receiver.(type) {
		case *builtin.Object:
			v.Context.Env.Set("this", obj)
		}
		r, err := m.Statements.Accept(v)
		Publish("method_end", v.Context, n)
		if err != nil {
			return nil, err
		}
		v.Context.Env = prev

		if r != nil {
			switch ret := r.(type) {
			case *Return:
				return ret.Value, nil
			}
			return r, nil
		}
	}
	return nil, nil
}

func (v *Interpreter) VisitNew(n *ast.New) (interface{}, error) {
	resolver := NewTypeResolver(v.Context)
	typeRef := n.Type.(*ast.TypeRef)
	classType, err := resolver.ResolveType(typeRef.Name)
	if err != nil {
		return nil, err
	}
	genericType := make([]*builtin.ClassType, len(n.Parameters))
	newObj := &builtin.Object{
		ClassType:      classType,
		InstanceFields: builtin.NewObjectMap(),
		GenericType:    genericType,
		Extra:          map[string]interface{}{},
	}
	for _, f := range classType.InstanceFields.Data {
		if f.Expression == nil {
			newObj.InstanceFields.Set(f.Name, builtin.Null)
		} else {
			r, _ := f.Expression.Accept(v)
			newObj.InstanceFields.Set(f.Name, r.(*builtin.Object))
		}
	}
	if len(classType.Constructors) > 0 {
		// TODO: implement multiple constructor
		evaluated := make([]*builtin.Object, len(n.Parameters))
		for i, p := range n.Parameters {
			r, err := p.Accept(v)
			if err != nil {
				return nil, err
			}
			evaluated[i] = r.(*builtin.Object)
		}
		typeResolver := NewTypeResolver(v.Context)
		constructor := typeResolver.SearchMethod(classType, classType.Constructors, evaluated)

		if constructor.NativeFunction != nil {
			constructor.NativeFunction(newObj, evaluated, v.Extra)
		} else {
			prev := v.Context.Env
			v.Context.Env = NewEnv(nil)
			for i, p := range constructor.Parameters {
				param := p.(*ast.Parameter)
				v.Context.Env.Set(param.Name, evaluated[i])
			}
			v.Context.Env.Set("this", newObj)
			constructor.Statements.Accept(v)
			v.Context.Env = prev
		}
	}
	if classType == builtin.ListType {
		if n.Init != nil && len(n.Init.Records) != 0 {
			records := make([]*builtin.Object, len(n.Init.Records))
			for i, r := range n.Init.Records {
				initRecord, err := r.Accept(v)
				if err != nil {
					return nil, err
				}
				records[i] = initRecord.(*builtin.Object)
			}
			newObj.Extra["records"] = records
		} else {
			newObj.Extra["records"] = []*builtin.Object{}
		}
	}
	if classType == builtin.MapType {
		if n.Init == nil {
			newObj.Extra["values"] = map[string]*builtin.Object{}
		} else if len(n.Init.Values) != 0 {
			values := map[string]*builtin.Object{}
			for key, value := range n.Init.Values {
				mapValue, err := value.Accept(v)
				if err != nil {
					return nil, err
				}
				mapKey, err := key.Accept(v)
				if err != nil {
					return nil, err
				}
				values[mapKey.(*builtin.Object).StringValue()] = mapValue.(*builtin.Object)
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
		exp := builtin.NewInteger(l.(*builtin.Object).IntegerValue() + 1)
		// TODO: implement
		v.Context.Env.Set(name.Value[0], exp)
		return exp, nil
	}
	if n.Op == "--" {
		name := n.Expression.(*ast.Name)
		l, _ := name.Accept(v)
		exp := builtin.NewInteger(l.(*builtin.Object).IntegerValue() - 1)
		// TODO: implement
		v.Context.Env.Set(name.Value[0], exp)
		return exp, nil
	}
	panic("not pass")
	return nil, nil
}

func (v *Interpreter) VisitBinaryOperator(n *ast.BinaryOperator) (interface{}, error) {
	left, _ := n.Left.Accept(v)
	right, _ := n.Right.Accept(v)

	lObj := left.(*builtin.Object)
	lType := lObj.ClassType
	rObj := right.(*builtin.Object)
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
			exp.(*builtin.Object).InstanceFields.Set(t.FieldName, rObj)
		}
		return rObj, nil
	case "+=":
		// TODO: implement
		var value *builtin.Object
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
			exp.(*builtin.Object).InstanceFields.Set(t.FieldName, value)
		}
		return value, nil
	case "-=":
		var value *builtin.Object
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
			exp.(*builtin.Object).InstanceFields.Set(t.FieldName, value)
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
		return &Return{res}, nil
	}
	return &Return{}, nil
}

func (v *Interpreter) VisitThrow(n *ast.Throw) (interface{}, error) {
	res, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	return &Raise{res.(*builtin.Object)}, nil
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
	return ast.VisitSwitch(v, n)
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
		d := declarator.(*ast.VariableDeclarator)
		if d.Expression != nil {
			val, err := d.Expression.Accept(v)
			if err != nil {
				panic(err)
			}
			v.Context.Env.Set(d.Name, val.(*builtin.Object))
		} else {
			v.Context.Env.Set(d.Name, builtin.Null)
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
		if !c.(*builtin.Object).BoolValue() {
			break
		}
		res, err := n.Statements.Accept(v)
		if res != nil {
			switch r := res.(type) {
			case *Break:
				return nil, nil
			case *Continue:
				continue
			case *Return, *Raise:
				return r, nil
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
	obj, _ := r.(*builtin.Object).InstanceFields.Get(n.FieldName)
	return obj, nil
}

func (v *Interpreter) VisitType(n *ast.TypeRef) (interface{}, error) {
	return ast.VisitType(v, n)
}

func (v *Interpreter) VisitBlock(n *ast.Block) (interface{}, error) {
	for _, stmt := range n.Statements {
		Publish("line", v.Context, stmt)
		res, err := stmt.Accept(v)
		if err != nil {
			return nil, err
		}
		if res != nil {
			switch r := res.(type) {
			case *Break, *Continue, *Return, *Raise:
				return r, nil
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
	if res.(*builtin.Object).Extra["value"].(bool) {
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

var createObject = builtin.CreateObject

type Return struct {
	Value interface{}
}

type Break struct{}

type Continue struct{}

type Raise struct {
	Value *builtin.Object
}
