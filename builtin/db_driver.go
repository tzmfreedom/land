package builtin

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseDriver struct {
	db *sql.DB
}

var databaseDriver = NewDatabaseDriver()

func NewDatabaseDriver() *DatabaseDriver {
	db, _ := sql.Open("sqlite3", "./database.sqlite3")
	return &DatabaseDriver{db}
}

type Scanner struct {
	ClassType *ClassType
	Fields *FieldMap
}

func NewScanner(t *ClassType) *Scanner {
	return &Scanner{
		t,
		NewFieldMap(),
	}
}

func (s *Scanner) Scan(src interface{}) error {
	return nil
}

func (d *DatabaseDriver) Query(soql string) []*Object {
	fields := make([]string, 3)
	rows, _ := d.db.Query(soql)
	records := []*Object{}
	for rows.Next() {
		dispatch := make([]interface{}, len(fields))
		err := rows.Scan(dispatch...)
		if err != nil {
			panic(err)
		}
		record := CreateObject(StringType)
		for i, field := range fields {
			record.InstanceFields.Set(field, NewString(dispatch[i].(string)))
		}
		records = append(records, record)
	}
	return records
}
