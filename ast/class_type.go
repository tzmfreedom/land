package ast

import (
	"fmt"
	"strings"
)

type ClassType struct {
	Annotations        []*Annotation
	Modifiers          []*Modifier
	Name               string
	SuperClassRef      *TypeRef
	SuperClass         *ClassType
	ImplementClasses   []*ClassType
	ImplementClassRefs []*TypeRef
	Constructors       []*Method
	InstanceFields     *FieldMap
	StaticFields       *FieldMap
	InstanceMethods    *MethodMap
	StaticMethods      *MethodMap
	InnerClasses       *ClassMap
	ToString           func(*Object) string
	Extra              map[string]interface{}
	Generics           []*ClassType
	Interface          bool
	Location           *Location
	Parent             Node
}

func (t *ClassType) IsInterface() bool {
	return t.Interface
}

func (t *ClassType) IsAbstract() bool {
	return t.Is("abstract")
}

func (t *ClassType) IsGenerics() bool {
	return t.Name == "List" ||
		t.Name == "Map" ||
		t.Name == "Set"
}

func (t *ClassType) String() string {
	if t.IsGenerics() {
		classTypes := t.Generics
		generics := make([]string, len(classTypes))
		for i, classType := range classTypes {
			generics[i] = classType.String()
		}
		return fmt.Sprintf("%s<%s>", t.Name, strings.Join(generics, ", "))
	}
	return t.Name
}

func (t *ClassType) HasConstructor() bool {
	if len(t.Constructors) > 0 {
		return true
	}
	if t.SuperClass != nil {
		return t.SuperClass.HasConstructor()
	}
	return false
}

func (t *ClassType) Is(name string) bool {
	name = strings.ToLower(name)
	for _, modifier := range t.Modifiers {
		modifierName := strings.ToLower(modifier.Name)
		if modifierName == name {
			return true
		}
	}
	return false
}

type Field struct {
	TypeRef    *TypeRef
	Type       *ClassType
	Modifiers  []*Modifier
	Name       string
	Expression Node
	Getter     Node
	Setter     Node
	Location   *Location
	Parent     Node
}

func (f *Field) IsAccessor(modifier string, checkSetter bool) bool {
	is := f.Is(modifier)
	if checkSetter {
		if f.Setter != nil && !f.Setter.(*GetterSetter).IsModifierBlank() {
			is = f.Setter.(*GetterSetter).Is(modifier)
		}
	} else {
		if f.Getter != nil && !f.Getter.(*GetterSetter).IsModifierBlank() {
			is = f.Getter.(*GetterSetter).Is(modifier)
		}
	}
	return is
}

func (f *Field) IsPublic(checkSetter bool) bool {
	return f.IsAccessor("public", checkSetter)
}

func (f *Field) IsPrivate(checkSetter bool) bool {
	return f.IsAccessor("private", checkSetter)
}

func (f *Field) IsProtected(checkSetter bool) bool {
	return f.IsAccessor("protected", checkSetter)
}

func (f *Field) AccessModifier(checkSetter bool) string {
	if f.IsPublic(checkSetter) {
		return "public"
	}
	if f.IsPrivate(checkSetter) {
		return "private"
	}
	if f.IsProtected(checkSetter) {
		return "protected"
	}
	return ""
}

func (f *Field) IsOverride() bool {
	return f.Is("override")
}

func (f *Field) IsAbstract() bool {
	return f.Is("abstract")
}

func (f *Field) IsVirtual() bool {
	return f.Is("virtual")
}

func (f *Field) Is(name string) bool {
	name = strings.ToLower(name)
	for _, modifier := range f.Modifiers {
		modifierName := strings.ToLower(modifier.Name)
		if modifierName == name {
			return true
		}
	}
	return false
}

type FieldMap struct {
	Data map[string]*Field
}

func NewFieldMap() *FieldMap {
	return &FieldMap{
		Data: map[string]*Field{},
	}
}

func (m *FieldMap) Set(k string, n *Field) {
	m.Data[strings.ToLower(k)] = n
}

func (m *FieldMap) Get(k string) (*Field, bool) {
	n, ok := m.Data[strings.ToLower(k)]
	return n, ok
}

type Method struct {
	Name          string
	Annotations   []*Annotation
	Modifiers     []*Modifier
	ReturnType    *ClassType
	ReturnTypeRef *TypeRef
	Parameters    []*Parameter
	Throws        []Node
	Statements    Node
	// func(receiver, value, options)
	NativeFunction func(*Object, []*Object, map[string]interface{}) interface{}
	Location       *Location
	Parent         Node
}

func NewMethod(decl *MethodDeclaration) *Method {
	return &Method{
		Name:          decl.Name,
		Annotations:   decl.Annotations,
		Modifiers:     decl.Modifiers,
		ReturnTypeRef: decl.ReturnType,
		Parameters:    decl.Parameters,
		Throws:        decl.Throws,
		Statements:    decl.Statements,
		Location:      decl.Location,
		Parent:        decl.Parent,
	}
}

func NewConstructor(decl *ConstructorDeclaration) *Method {
	return &Method{
		Name:        decl.GetParent().(*ClassDeclaration).Name,
		Modifiers:   decl.Modifiers,
		Annotations: decl.Annotations,
		Parameters:  decl.Parameters,
		Throws:      decl.Throws,
		Statements:  decl.Statements,
		Location:    decl.Location,
		Parent:      decl.Parent,
	}
}

func (m *Method) IsPublic() bool {
	return m.Is("public")
}

func (m *Method) IsPrivate() bool {
	return m.Is("private")
}

func (m *Method) IsProtected() bool {
	return m.Is("protected")
}

func (m *Method) IsTestMethod() bool {
	return m.Is("testMethod") || m.IsAnnotated("isTest")
}

func (m *Method) IsAnnotated(name string) bool {
	name = strings.ToLower(name)
	for _, annotation := range m.Annotations {
		if strings.ToLower(annotation.Name) == name {
			return true
		}
	}
	return false
}

func (m *Method) AccessModifier() string {
	if m.IsPublic() {
		return "public"
	}
	if m.IsPrivate() {
		return "private"
	}
	if m.IsProtected() {
		return "protected"
	}
	return ""
}

func (m *Method) IsOverride() bool {
	return m.Is("override")
}

func (m *Method) IsAbstract() bool {
	return m.Is("abstract")
}

func (m *Method) IsVirtual() bool {
	return m.Is("virtual")
}

func (m *Method) Is(name string) bool {
	name = strings.ToLower(name)
	for _, modifier := range m.Modifiers {
		modifierName := strings.ToLower(modifier.Name)
		if modifierName == name {
			return true
		}
	}
	return false
}

type MethodMap struct {
	Data map[string][]*Method
}

func NewMethodMap() *MethodMap {
	return &MethodMap{
		Data: map[string][]*Method{},
	}
}

func (m *MethodMap) Add(k string, n *Method) {
	if data, ok := m.Get(k); ok {
		data = append(data, n)
		m.Set(k, data)
	} else {
		m.Set(k, []*Method{n})
	}
}

func (m *MethodMap) Set(k string, n []*Method) {
	m.Data[strings.ToLower(k)] = n
}

func (m *MethodMap) Get(k string) ([]*Method, bool) {
	n, ok := m.Data[strings.ToLower(k)]
	return n, ok
}

func (m *MethodMap) All() [][]*Method {
	methods := make([][]*Method, len(m.Data))
	for _, v := range m.Data {
		methods = append(methods, v)
	}
	return methods
}

/**
 * ClassMap
 */
type ClassMap struct {
	Data map[string]*ClassType
}

func NewClassMap() *ClassMap {
	return &ClassMap{
		Data: map[string]*ClassType{},
	}
}

func (m *ClassMap) Set(k string, n *ClassType) {
	m.Data[strings.ToLower(k)] = n
}

func (m *ClassMap) Get(k string) (*ClassType, bool) {
	n, ok := m.Data[strings.ToLower(k)]
	return n, ok
}

func (m *ClassMap) Clear() {
	m.Data = map[string]*ClassType{}
}

func CreateClass(
	name string,
	constructors []*Method,
	instanceMethods *MethodMap,
	staticMethods *MethodMap,
) *ClassType {
	return &ClassType{
		Name:            name,
		Modifiers:       []*Modifier{PublicModifier()},
		Constructors:    constructors,
		InstanceFields:  NewFieldMap(),
		StaticFields:    NewFieldMap(),
		InstanceMethods: instanceMethods,
		StaticMethods:   staticMethods,
		InnerClasses:    NewClassMap(),
	}
}

func CreateMethod(
	name string,
	returnType *ClassType,
	parameters []*Parameter,
	nativeFunction func(*Object, []*Object, map[string]interface{}) interface{},
) *Method {
	return &Method{
		Name:           name,
		Modifiers:      []*Modifier{PublicModifier()},
		ReturnType:     returnType,
		Parameters:     parameters,
		NativeFunction: nativeFunction,
	}
}

func CreateField(
	name string,
	fieldType *ClassType,
) *Field {
	return &Field{
		Name:      name,
		Modifiers: []*Modifier{PublicModifier()},
		Type:      fieldType,
	}
}

var publicModifier = &Modifier{Name: "public"}
var privateModifier = &Modifier{Name: "private"}
var protectedModifier = &Modifier{Name: "protected"}
var globalModifier = &Modifier{Name: "global"}
var abstractModifier = &Modifier{Name: "abstract"}

func PublicModifier() *Modifier {
	return publicModifier
}
