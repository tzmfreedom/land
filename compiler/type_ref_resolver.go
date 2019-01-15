package compiler

import (
	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/builtin"
)

type TypeRefResolver struct {
	resolver *builtin.TypeRefResolver
}

func NewTypeRefResolver(classTypes *builtin.ClassMap, nameSpaceStore *builtin.NameSpaceStore) *TypeRefResolver {
	return &TypeRefResolver{
		resolver: &builtin.TypeRefResolver{
			ClassTypes: classTypes,
			NameSpaces: nameSpaceStore,
		},
	}
}

func (r *TypeRefResolver) Resolve(n *builtin.ClassType) (*builtin.ClassType, error) {
	var err error
	if n.SuperClassRef != nil {
		n.SuperClass, err = r.resolver.ResolveType(n.SuperClassRef.(*ast.TypeRef).Name)
		if err != nil {
			return nil, err
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
