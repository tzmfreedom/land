package builtin

import (
	"fmt"
	"strings"

	"github.com/tzmfreedom/goland/ast"
)

type Relation struct {
	Name        string
	ReferenceTo string
}

type SqlBuilder struct {
	interpreter ast.Visitor
}

func (b *SqlBuilder) Build(n *ast.Soql) (string, [][]string) {
	selectClause, selectFields := createSelectClause(n)

	relations := createRelations(n.FromObject, selectFields)

	leftJoinClause := createLeftJoins(n.FromObject, relations)

	whereClause := b.createWhere(n.Where)
	if whereClause != "" {
		whereClause = " WHERE " + whereClause
	}

	sql := fmt.Sprintf(
		"SELECT %s FROM %s%s%s",
		selectClause,
		n.FromObject,
		leftJoinClause,
		whereClause,
	)
	return sql, selectFields
}

func (b *SqlBuilder) createWhere(n ast.Node) string {
	switch val := n.(type) {
	case *ast.WhereCondition:
		var field string
		switch f := val.Field.(type) {
		case *ast.SelectField:
			field = strings.Join(f.Value, ".")
		case *ast.SoqlFunction:
			field = f.Name + "()"
		}
		value, _ := val.Expression.Accept(b.interpreter)
		return fmt.Sprintf("%s %s '%s'", field, val.Op, String(value.(*Object)))
	case *ast.WhereBinaryOperator:
		where := ""
		if val.Left != nil {
			where += b.createWhere(val.Left)
		}
		if val.Right != nil {
			where += fmt.Sprintf("%s %s", val.Op, b.createWhere(val.Right))
		}
		return where
	}
	return ""
}

func createSelectClause(n *ast.Soql) (string, [][]string) {
	selectFields := make([][]string, len(n.SelectFields))
	for i, selectField := range n.SelectFields {
		v := selectField.(*ast.SelectField).Value
		if len(v) == 1 {
			selectFields[i] = []string{n.FromObject, v[0]}
		} else {
			selectFields[i] = v
		}
	}

	tempFields := make([]string, len(n.SelectFields))
	for i, selectField := range selectFields {
		tempFields[i] = strings.Join(selectField, ".")
	}
	return strings.Join(tempFields, ", "), selectFields
}

func createRelations(from string, fields [][]string) map[string]Relation {
	relations := map[string]Relation{}
	for _, field := range fields {
		if from == field[0] {
			continue
		}

		sObject := sObjects[from]
		relationshipName := field[0]
		var targetField SobjectField
		for _, sObjectField := range sObject.Fields {
			if sObjectField.RelationshipName == relationshipName {
				targetField = sObjectField
				break
			}
		}
		relations[targetField.RelationshipName] = Relation{
			Name:        targetField.Name,
			ReferenceTo: targetField.ReferenceTo[0],
		}
	}
	return relations
}

func createLeftJoins(from string, relations map[string]Relation) string {
	leftJoins := []string{}
	for _, relation := range relations {
		leftJoins = append(
			leftJoins,
			fmt.Sprintf(
				"LEFT JOIN %s ON %s.%s = %s.id",
				relation.ReferenceTo,
				from,
				relation.Name,
				relation.ReferenceTo,
			),
		)
	}

	if len(leftJoins) == 0 {
		return ""
	}
	return " " + strings.Join(leftJoins, " AND ")
}
