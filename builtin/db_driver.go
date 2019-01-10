package builtin

import (
	"database/sql"

	"fmt"
	"strings"

	"github.com/k0kubun/pp"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tzmfreedom/goland/ast"
)

type databaseDriver struct {
	db *sql.DB
}

var DatabaseDriver = NewDatabaseDriver()

func NewDatabaseDriver() *databaseDriver {
	db, _ := sql.Open("sqlite3", "./database.sqlite3")
	return &databaseDriver{db}
}

type Relation struct {
	Name        string
	ReferenceTo string
}

func (d *databaseDriver) Query(n *ast.Soql) []*Object {
	visitor := &ast.TosVisitor{}
	r, err := n.Accept(visitor)
	if err != nil {
		panic(err)
	}
	relations := map[string]Relation{}
	fields := make([][]string, len(n.SelectFields))
	for i, field := range n.SelectFields {
		switch f := field.(type) {
		case *ast.SelectField:
			fields[i] = f.Value
			// relation
			if len(f.Value) == 2 {
				sObject := sObjects[n.FromObject]
				relationshipName := f.Value[0]
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
		}
	}
	leftJoins := []string{}
	for _, relation := range relations {
		leftJoins = append(
			leftJoins,
			fmt.Sprintf(
				"LEFT JOIN %s ON %s.%s = %s.id",
				relation.ReferenceTo,
				n.FromObject,
				relation.Name,
				relation.ReferenceTo,
			),
		)
	}

	soql := r.(string)
	soql = soql[1:len(soql)-1] + strings.Join(leftJoins, " ")
	soql = strings.Replace(soql, " id", " "+n.FromObject+".id", -1)
	soql = strings.Replace(soql, " name", " "+n.FromObject+".name", -1)
	pp.Println(soql)
	rows, err := d.db.Query(soql)
	if err != nil {
		panic(err)
	}

	classType, _ := PrimitiveClassMap().Get(n.FromObject)
	records := []*Object{}
	for rows.Next() {
		dispatches := make([]interface{}, len(fields))
		for i, _ := range fields {
			var temp string
			dispatches[i] = &temp
		}
		err := rows.Scan(dispatches...)
		if err != nil {
			panic(err)
		}
		record := CreateObject(classType)
		for i, field := range fields {
			fieldName := field[0]
			record.InstanceFields.Set(fieldName, NewString(*dispatches[i].(*string)))
		}
		records = append(records, record)
	}
	return records
}

func (d *databaseDriver) QueryRaw(query string) {
	rows, err := d.db.Query(query)
	if err != nil {
		panic(err)
	}
	pp.Println(rows)
}

func (d *databaseDriver) Execute(query string) {
	d.db.Exec(query)
}

func seed() {
	DatabaseDriver.Execute(`
INSERT INTO Account(id, name) VALUES ('12345', 'hoge');
INSERT INTO Account(id, name) VALUES ('abcde', 'fuga');
INSERT INTO Contact(id, lastname, firstname, accountid) VALUES ('a', 'l1', 'r1', '12345');
INSERT INTO Contact(id, lastname, firstname, accountid) VALUES ('b', 'l2', 'r2', 'abcde');
`)
}

func init() {
	DatabaseDriver.Execute(`
CREATE TABLE IF NOT EXISTS Account (
	id VARCHAR NOT NULL PRIMARY KEY,
	name TEXT
);

CREATE TABLE IF NOT EXISTS Contact (
	id VARCHAR NOT NULL PRIMARY KEY,
	lastname TEXT,
	firstname TEXT,
	accountid TEXT	
);
`)
	seed()
}
