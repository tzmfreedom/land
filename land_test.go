package main

import (
	"os"

	"github.com/tzmfreedom/land/ast"
)

func setup() {
	classMap = ast.NewClassMap()
}

// Arithmetic
func ExampleRun1() {
	setup()
	os.Args = []string{"land", "run", "-a", "Foo#action", "-f", "fixtures/example1.cls"}
	main()
	// Output:
	// 6
	// 7
	// 9
	// hoge
	// foo/bar
	// 1.560000
}

// Object Creation, FieldAccess
func ExampleRun2() {
	setup()
	os.Args = []string{"land", "run", "-a", "Foo#action", "-f", "fixtures/example2.cls"}
	main()
	// Output:
	// <Foo> {
	//   b: false
	//   d: 1.230000
	//   i: 100
	//   s: foo
	// }
	// 200
	// foo
	// foo
	// 100
	// false
	// foo
	// 1.230000
}

// For, While, Continue, Break, If, Else
func ExampleRun3() {
	setup()
	os.Args = []string{"land", "run", "-a", "Foo#action", "-f", "fixtures/example3.cls"}
	main()
	// Output:
	// 0
	// 1
	// 2
	// 30
	// 40
	// true
	// false
}

// For, While, Continue, Break, If, Else
func ExampleInterface() {
	setup()
	os.Args = []string{"land", "run", "-a", "Implemented#main", "-d", "fixtures/interface"}
	main()
	// Output:
	// 1.200000
	// 3.400000
}

// For, While, Continue, Break, If, Else
func ExampleAbstract() {
	setup()
	os.Args = []string{"land", "run", "-a", "Extended#main", "-d", "fixtures/abstract"}
	main()
	// Output:
	// hello
	// world
}
