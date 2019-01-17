package compiler

import (
	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/builtin"
)

type TypeRefResolver struct {
	resolver *builtin.TypeRefResolver
}

func NewTypeRefResolver(classTypes *ast.ClassMap, nameSpaceStore *builtin.NameSpaceStore) *TypeRefResolver {
	return &TypeRefResolver{
		resolver: &builtin.TypeRefResolver{
			ClassTypes: classTypes,
			NameSpaces: nameSpaceStore,
		},
	}
}

func (r *TypeRefResolver) Resolve(n *ast.ClassType) (*ast.ClassType, error) {
	var err error
	if n.SuperClassRef != nil {
		n.SuperClass, err = r.resolver.ResolveType(n.SuperClassRef.Name)
		if err != nil {
			return nil, err
		}
	}

	if n.ImplementClassRefs != nil {
		n.ImplementClasses = make([]*ast.ClassType, len(n.ImplementClassRefs))
		for i, impl := range n.ImplementClassRefs {
			n.ImplementClasses[i], err = r.resolver.ResolveType(impl.Name)
			if err != nil {
				return nil, err
			}
		}
	}

	for _, c := range n.InnerClasses.Data {
		_, err = r.Resolve(c)
		if err != nil {
			return nil, err
		}
	}
	return n, nil
}
