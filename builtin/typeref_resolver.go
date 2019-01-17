package builtin

import (
	"fmt"
	"strings"

	"github.com/tzmfreedom/goland/ast"
)

type TypeRefResolver struct {
	NameSpaces   *NameSpaceStore
	ClassTypes   *ast.ClassMap
	CurrentClass *ast.ClassType
}

func (r *TypeRefResolver) ResolveType(names []string) (*ast.ClassType, error) {
	if len(names) == 1 {
		className := names[0]
		if class, ok := r.ClassTypes.Get(className); ok {
			return class, nil
		}
		if classTypes, ok := r.NameSpaces.Get("System"); ok {
			if class, ok := classTypes.Get(className); ok {
				return class, nil
			}
		}
		// search for UserClass.InnerClass
		if r.CurrentClass != nil {
			if class, ok := r.CurrentClass.InnerClasses.Get(className); ok {
				return class, nil
			}
		}
	} else if len(names) == 2 {
		// search for UserClass.InnerClass
		if class, ok := r.ClassTypes.Get(names[0]); ok {
			if inner, ok := class.InnerClasses.Get(names[1]); ok {
				return inner, nil
			}
		}
		// search for NameSpace.UserClass
		if classTypes, ok := r.NameSpaces.Get(names[0]); ok {
			if class, ok := classTypes.Get(names[1]); ok {
				return class, nil
			}
		}
	} else if len(names) == 3 {
		// search for NameSpace.UserClass.InnerClass
		if classTypes, ok := r.NameSpaces.Get(names[0]); ok {
			if class, ok := classTypes.Get(names[1]); ok {
				if inner, ok := class.InnerClasses.Get(names[2]); ok {
					return inner, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("%s does not found", strings.Join(names, "."))
}

func (r *TypeRefResolver) ConvertType(n *ast.TypeRef) (*ast.ClassType, error) {
	// convert list from array
	for n.Dimmension > 0 {
		name := n.Name
		params := n.Parameters
		n.Name = []string{"List"}
		n.Parameters = []ast.Node{
			&ast.TypeRef{
				Name:       name,
				Parameters: params,
			},
		}
		n.Dimmension--
	}

	t, err := r.ResolveType(n.Name)
	if err != nil {
		return nil, err
	}
	if t.IsGeneric() {
		types := make([]*ast.ClassType, len(n.Parameters))
		var err error
		for i, p := range n.Parameters {
			types[i], err = r.ConvertType(p.(*ast.TypeRef))
			if err != nil {
				return nil, err
			}
		}
		return &ast.ClassType{
			Name:            t.Name,
			Constructors:    t.Constructors,
			InstanceMethods: t.InstanceMethods,
			StaticMethods:   t.StaticMethods,
			Extra: map[string]interface{}{
				"generics": types,
			},
		}, nil
	}
	return t, nil
}
