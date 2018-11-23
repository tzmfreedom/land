package main

import (
	"os"
)

func ExampleRun1() {
	os.Args = []string{"land", "run", "-f", "fixtures/example.cls"}
	main()
	// Output:
	// 6
	// 7
	// 9
	// hoge
	// foo/bar
	// 1.560000
}

func ExampleRun2() {
	os.Args = []string{"land", "run", "-f", "fixtures/example2.cls"}
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

func ExampleRun3() {
	os.Args = []string{"land", "run", "-f", "fixtures/example3.cls"}
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
