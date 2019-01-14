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

func (d *databaseDriver) Query(n *ast.Soql, interpreter ast.Visitor) []*Object {
	builder := SqlBuilder{interpreter: interpreter}
	sql, selectFields, relations := builder.Build(n)
	// pp.Println(sql)

	rows, err := d.db.Query(sql)
	if err != nil {
		panic(err)
	}

	classType, _ := PrimitiveClassMap().Get(n.FromObject)
	records := []*Object{}
	for rows.Next() {
		dispatches := make([]interface{}, len(selectFields))
		for i, _ := range selectFields {
			var temp string
			dispatches[i] = &temp
		}
		err := rows.Scan(dispatches...)
		if err != nil {
			panic(err)
		}
		record := CreateObject(classType)
		for i, field := range selectFields {
			tmpTable := field[0]
			fieldName := field[1]
			value := NewString(*dispatches[i].(*string))

			if tmpTable == "t0" {
				record.InstanceFields.Set(fieldName, value)
				continue
			}
			relationInfo := relations[tmpTable]
			relationField, ok := record.InstanceFields.Get(relationInfo.RelationshipName)
			if ok {
				relationField.InstanceFields.Set(fieldName, value)
			} else {
				relationType, _ := PrimitiveClassMap().Get(relationInfo.ReferenceTo)
				relationField = CreateObject(relationType)
				relationField.InstanceFields.Set(fieldName, value)
				record.InstanceFields.Set(relationInfo.RelationshipName, relationField)
			}
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

func (d *databaseDriver) Execute(dmlType string, sObjectType string, records []*Object, upsertKey string) {
	for _, record := range records {
		var query string

		switch dmlType {
		case "insert":
			fields := []string{}
			values := []string{}
			for name, field := range record.InstanceFields.All() {
				// TODO: convert type
				if field == Null {
					continue
				}
				fields = append(fields, name)
				values = append(values, field.StringValue())
			}
			query = fmt.Sprintf(
				"INSERT INTO %s(%s) VALUES (%s)",
				sObjectType,
				strings.Join(fields, ", "),
				strings.Join(values, ", "),
			)
		case "update":
			updateFields := []string{}
			for name, field := range record.InstanceFields.All() {
				// TODO: convert type
				if field == Null {
					continue
				}
				updateFields = append(updateFields, fmt.Sprintf("%s = '%s'", name, field.StringValue()))
			}
			id, ok := record.InstanceFields.Get("Id")
			if !ok {
				panic("id does not exist")
			}
			query = fmt.Sprintf(
				"UPDATE %s SET %s WHERE id = '%s'",
				sObjectType,
				strings.Join(updateFields, ", "),
				id.StringValue(),
			)
		case "upsert":
			// TODO: implement
		case "delete":
			id, ok := record.InstanceFields.Get("Id")
			if !ok {
				panic("id does not exist")
			}
			query = fmt.Sprintf(
				"DELETE FROM %s WHERE id = '%s'",
				sObjectType,
				id.StringValue(),
			)
		}
		d.db.Exec(query)
	}
}

func (d *databaseDriver) ExecuteRaw(query string) {
	d.db.Exec(query)
}

func Seed() {
	DatabaseDriver.ExecuteRaw(`
INSERT INTO Account(id, name) VALUES ('12345', 'hoge');
INSERT INTO Account(id, name) VALUES ('abcde', 'fuga');
INSERT INTO Contact(id, lastname, firstname, accountid) VALUES ('a', 'l1', 'r1', '12345');
INSERT INTO Contact(id, lastname, firstname, accountid) VALUES ('b', 'l2', 'r2', 'abcde');
`)
}

func Setup() {
	DatabaseDriver.ExecuteRaw(`
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
}
