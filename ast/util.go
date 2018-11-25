package ast

func IsDecendants(n Node, typeName string) bool {
	parent := n.GetParent()
	if parent == nil {
		return false
	}
	if parent.GetType() == typeName {
		return true
	}
	return IsDecendants(parent, typeName)
}

func IsParent(n Node, typeName string) bool {
	parent := n.GetParent()
	if parent != nil && parent.GetType() == typeName {
		return true
	}
	return false
}
