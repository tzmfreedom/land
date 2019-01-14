package builtin

import (
	"fmt"
	"strings"

	"github.com/tzmfreedom/goland/ast"
)

type Relation struct {
	RelationshipName string
	FieldName        string
	ReferenceTo      string
}

type SqlBuilder struct {
	interpreter ast.Visitor
}

func (b *SqlBuilder) Build(n *ast.Soql) (string, [][]string, map[string]Relation) {
	tmpTableMap := map[string]string{}
	selectClause, selectFields := createSelectClause(n, tmpTableMap)
	whereClause := b.createWhere(n.Where, tmpTableMap)
	if whereClause != "" {
		whereClause = " WHERE " + whereClause
	}

	relations := createRelations(n.FromObject, tmpTableMap)

	leftJoinClause := createLeftJoins(relations)

	sql := fmt.Sprintf(
		"SELECT %s FROM %s t0%s%s",
		selectClause,
		n.FromObject,
		leftJoinClause,
		whereClause,
	)
	return sql, selectFields, relations
}

func (b *SqlBuilder) createWhere(n ast.Node, tmpTableMap map[string]string) string {
	switch val := n.(type) {
	case *ast.WhereCondition:
		var field string
		switch f := val.Field.(type) {
		case *ast.SelectField:
			if len(f.Value) == 1 {
				field = fmt.Sprintf("t0.%s", f.Value[0])
			} else {
				tmpTable, ok := tmpTableMap[f.Value[0]]
				if !ok {
					// TODO: key from string to integer?
					tmpTableIndex := len(tmpTable) + 1
					tmpTable = fmt.Sprintf("t%d", tmpTableIndex)
					tmpTableMap[f.Value[0]] = tmpTable
				}
				field = fmt.Sprintf("%s.%s", tmpTable, strings.Join(f.Value[1:], "."))
			}
		case *ast.SoqlFunction:
			field = f.Name + "()"
		}
		value, err := val.Expression.Accept(b.interpreter)
		if err != nil {
			panic(err)
		}
		return fmt.Sprintf("%s %s '%s'", field, val.Op, String(value.(*Object)))
	case *ast.WhereBinaryOperator:
		where := ""
		if val.Left != nil {
			where += b.createWhere(val.Left, tmpTableMap)
		}
		if val.Right != nil {
			where += fmt.Sprintf("%s %s", val.Op, b.createWhere(val.Right, tmpTableMap))
		}
		return where
	}
	return ""
}

func createSelectClause(n *ast.Soql, tmpTableMap map[string]string) (string, [][]string) {
	selectFields := make([][]string, len(n.SelectFields))
	// TODO: case insensitive
	tmpTableIndex := 1
	for i, selectField := range n.SelectFields {
		v := selectField.(*ast.SelectField).Value
		if len(v) == 1 {
			selectFields[i] = []string{"t0", v[0]}
		} else {
			relationshipName := v[0]
			if _, ok := tmpTableMap[relationshipName]; !ok {
				tmpTableMap[relationshipName] = fmt.Sprintf("t%d", tmpTableIndex)
				tmpTableIndex++
			}
			tmpTable := tmpTableMap[relationshipName]
			// TODO: recursive relation
			selectFields[i] = []string{tmpTable, v[1]}
		}
	}

	tempFields := make([]string, len(selectFields))
	for i, selectField := range selectFields {
		tempFields[i] = strings.Join(selectField, ".")
	}
	return strings.Join(tempFields, ", "), selectFields
}

func createRelations(from string, tmpTableMap map[string]string) map[string]Relation {
	relations := map[string]Relation{}
	sObject := sObjects[from]
	for relationshipName, tmpTableName := range tmpTableMap {
		var targetField SobjectField
		for _, sObjectField := range sObject.Fields {
			if sObjectField.RelationshipName == relationshipName {
				targetField = sObjectField
				break
			}
		}
		relations[tmpTableName] = Relation{
			RelationshipName: relationshipName,
			FieldName:        targetField.Name,
			ReferenceTo:      targetField.ReferenceTo[0],
		}
	}
	return relations
}

func createLeftJoins(relations map[string]Relation) string {
	leftJoins := []string{}
	for tmpTable, relation := range relations {
		leftJoins = append(
			leftJoins,
			fmt.Sprintf(
				"LEFT JOIN %s %s ON %s.%s = %s.id",
				relation.ReferenceTo,
				tmpTable,
				"t0", // TODO: recursive relation
				relation.FieldName,
				tmpTable,
			),
		)
	}

	if len(leftJoins) == 0 {
		return ""
	}
	return " " + strings.Join(leftJoins, " ")
}
