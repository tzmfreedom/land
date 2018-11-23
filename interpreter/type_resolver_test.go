package interpreter

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/k0kubun/pp"
	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/builtin"
)

func TestResolveVariable(t *testing.T) {
	testCases := []struct {
		Input    []string
		Context  *Context
		Expected *builtin.ClassType
		Error    error
	}{}
	typeResolver := &TypeResolver{}
	for _, testCase := range testCases {
		actual, err := typeResolver.ResolveVariable(testCase.Input, testCase.Context)
		if testCase.Error != nil && testCase.Error.Error() != err.Error() {
			diff := cmp.Diff(testCase.Error.Error(), err.Error())
			t.Errorf(diff)
		}

		if ok := cmp.Equal(testCase.Expected, actual); !ok {
			diff := cmp.Diff(testCase.Expected, actual)
			pp.Print(actual)
			t.Errorf(diff)
		}
	}
}

func TestResolveType(t *testing.T) {
	testCases := []struct {
		Input    []string
		Context  *Context
		Expected *builtin.ClassType
		Error    error
	}{}
	typeResolver := &TypeResolver{}
	for _, testCase := range testCases {
		actual, err := typeResolver.ResolveType(testCase.Input, testCase.Context)
		if testCase.Error != nil && testCase.Error.Error() != err.Error() {
			diff := cmp.Diff(testCase.Error.Error(), err.Error())
			t.Errorf(diff)
		}

		if ok := cmp.Equal(testCase.Expected, actual); !ok {
			diff := cmp.Diff(testCase.Expected, actual)
			pp.Print(actual)
			t.Errorf(diff)
		}
	}
}

func TestResolveMethod(t *testing.T) {
	testCases := []struct {
		Input    []string
		Context  *Context
		Expected *ast.MethodDeclaration
		Error    error
	}{}
	typeResolver := &TypeResolver{}
	for _, testCase := range testCases {
		actual, err := typeResolver.ResolveMethod(testCase.Input, testCase.Context)
		if testCase.Error != nil && testCase.Error.Error() != err.Error() {
			diff := cmp.Diff(testCase.Error.Error(), err.Error())
			t.Errorf(diff)
		}

		if ok := cmp.Equal(testCase.Expected, actual); !ok {
			diff := cmp.Diff(testCase.Expected, actual)
			pp.Print(actual)
			t.Errorf(diff)
		}
	}
}
