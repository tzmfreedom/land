package builtin

import (
	"database/sql"

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
	sql, fields := builder.Build(n)

	//pp.Println(sql)
	rows, err := d.db.Query(sql)
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
			if len(field) == 1 {
				fieldName := field[0]
				record.InstanceFields.Set(fieldName, NewString(*dispatches[i].(*string)))
			} else {
				fieldName := field[0]
				relation, ok := record.InstanceFields.Get(fieldName)
				value := NewString(*dispatches[i].(*string))
				if ok {
					relation.InstanceFields.Set(field[1], value)
				} else {
					// TODO: duplicate code
					relationType, _ := PrimitiveClassMap().Get(fieldName)
					relation = CreateObject(relationType)
					relation.InstanceFields.Set(field[1], value)
					record.InstanceFields.Set(fieldName, relation)
				}
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
