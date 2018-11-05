package ast

type Position struct {
	FileName string
	Column   int
	Line     int
}

type ClassDeclaration struct {
	Modifiers []Modifier
	Name      string
	Position  *Position
}

type Modifier struct {
	Name     string
	Position *Position
}

type Annotation struct {
	Name       string
	Parameters []interface{}
	Position   *Position
}

type Interface struct {
	Annotations     []interface{}
	Modifiers       []Modifier
	Name            string
	SuperClass      []interface{}
	InstanceMethods []map[string]string
	StaticMethods   []map[string]string
	Position        *Position
}

type Integer struct {
	Value    int
	Position *Position
}

type Parameter struct {
	Modifiers []Modifier
	Type      string
	Name      string
	Position  *Position
}

type ArrayAccess struct {
	Receiver interface{}
	Key      interface{}
	Position *Position
}

type Boolean struct {
	Value    bool
	Position *Position
}

type Break struct {
	Position *Position
}

type Continue struct {
	Position *Position
}

type Dml struct {
	Type       string
	Expression interface{}
	UpsertKey  string
	Position   *Position
}

type Double struct {
	Value    float32
	Position *Position
}

type FieldDeclaration struct {
	Type       string
	Modifiers  []Modifier
	Expression interface{}
	Position   *Position
	Getter     interface{}
	Setter     interface{}
}

type Try struct {
	Block        interface{}
	CatchClause  interface{}
	FinallyBlock interface{}
	Position     *Position
}

type Catch struct {
	Modifiers  []Modifier
	Type       string
	Identifier string
	Block      interface{}
	Position   *Position
}

type ConstructorDeclaration struct {
}
