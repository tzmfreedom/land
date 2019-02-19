package builtin

import (
	"database/sql"

	"fmt"
	"strings"

	"math/rand"
	"time"

	"github.com/k0kubun/pp"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tzmfreedom/land/ast"
)

type databaseDriver struct {
	db *sql.DB
}

var DatabaseDriver = NewDatabaseDriver()

func NewDatabaseDriver() *databaseDriver {
	// TODO: implment not sqlite3
	db, _ := sql.Open("sqlite3", "./database.sqlite3")
	return &databaseDriver{db}
}

func (d *databaseDriver) Query(n *ast.Soql, interpreter ast.Visitor) []*ast.Object {
	builder := SqlBuilder{interpreter: interpreter}
	sql, selectFields, relations := builder.Build(n)
	// pp.Println(sql)

	rows, err := d.db.Query(sql)
	if err != nil {
		panic(err)
	}

	classType, _ := PrimitiveClassMap().Get(n.FromObject)
	records := []*ast.Object{}
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
		record := ast.CreateObject(classType)
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
				relationField = ast.CreateObject(relationType)
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

func (d *databaseDriver) Begin() {
	d.db.Exec("BEGIN;")
}

func (d *databaseDriver) Rollback() {
	d.db.Exec("ROLLBACK;")
}

func (d *databaseDriver) Execute(dmlType string, sObjectType string, records []*ast.Object, upsertKey string) *ast.Object {
	saveResults := make([]*ast.Object, len(records))
	for i, record := range records {
		var query string

		switch dmlType {
		case "insert":
			fields := []string{}
			values := []string{}
			rand.Seed(time.Now().UnixNano())
			record.InstanceFields.Set("Id", NewString(string(rand.Int())))
			for name, field := range record.InstanceFields.All() {
				// TODO: convert type
				if field == Null {
					continue
				}
				fields = append(fields, name)
				values = append(values, fmt.Sprintf("'%s'", field.StringValue()))
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
		result, err := d.db.Exec(query)
		if err != nil {
			panic(err)
		}
		obj := ast.CreateObject(saveResultType)
		obj.Extra["isSuccess"] = NewBoolean(err == nil)
		if err == nil {
			id, err := result.LastInsertId()
			if err != nil {
				panic(err)
			}
			obj.Extra["id"] = NewString(fmt.Sprintf("%d", id))
		} else {
			obj.Extra["errors"] = err.Error()
		}
		saveResults[i] = obj
	}
	listObject := ast.CreateObject(ListType)
	listObject.Extra["records"] = saveResults
	return listObject
}

func (d *databaseDriver) ExecuteRaw(query string, args ...interface{}) error {
	_, err := d.db.Exec(query, args...)
	return err
}

var dbTypeMapper = map[string]string{
	"string":                     "TEXT",
	"picklist":                   "TEXT",
	"multipicklist":              "TEXT",
	"combobox":                   "TEXT",
	"reference":                  "TEXT",
	"boolean":                    "INT",
	"currency":                   "REAL",
	"textarea":                   "TEXT",
	"int":                        "INT",
	"double":                     "REAL",
	"percent":                    "REAL",
	"id":                         "TEXT",
	"date":                       "TEXT",
	"datetime":                   "TEXT",
	"time":                       "TEXT",
	"url":                        "TEXT",
	"email":                      "TEXT",
	"encryptedstring":            "TEXT",
	"datacategorygroupreference": "TEXT",
	"location":                   "TEXT",
	"address":                    "TEXT",
	"anyType":                    "TEXT",
	"complexvalue":               "TEXT",
	"phone":                      "TEXT",
}

func CreateDatabase(src string) error {
	loader := NewMetaFileLoader(src)
	sobjects, err := loader.Load()
	if err != nil {
		return err
	}
	for name, sobject := range sobjects {
		fields := make([]string, len(sobject.Fields))
		for i, field := range sobject.Fields {
			if field.Name == "id" {
				fields[i] = "id VARCHAR NOT NULL PRIMARY KEY"
			} else {
				if _, ok := dbTypeMapper[field.Type]; !ok {
					return fmt.Errorf("undefined type mapper %s", field.Type)
				}
				fields[i] = fmt.Sprintf("`%s` %s", field.Name, dbTypeMapper[field.Type])
			}
		}
		query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (%s);", name, strings.Join(fields, ", "))
		err := DatabaseDriver.ExecuteRaw(query)
		if err != nil {
			return err
		}
	}
	return nil
}

func Seed(username, password, endpoint, src string) error {
	loader := NewMetaFileLoader(src)
	sobjects, err := loader.Load()
	if err != nil {
		return err
	}
	client := NewSoapClient(username, password, endpoint)
	for name, sobject := range sobjects {
		fields := make([]string, len(sobject.Fields))
		for i, field := range sobject.Fields {
			fields[i] = fmt.Sprintf("%s", field.Name)
		}
		soql := fmt.Sprintf("SELECT %s FROM %s", strings.Join(fields, ","), name)
		r, err := client.Query(soql)
		if err != nil {
			return err
		}
		for _, record := range r.Records {
			insertFields := make([]string, len(record.Fields)+1)
			insertValues := make([]interface{}, len(record.Fields)+1)
			placeholders := make([]string, len(record.Fields)+1)
			insertFields[0] = "`id`"
			insertValues[0] = "'" + record.Id + "'"
			placeholders[0] = "?"
			i := 1
			for key, insertField := range record.Fields {
				insertFields[i] = "`" + key + "`"
				insertValues[i] = "'" + insertField.(string) + "'"
				placeholders[i] = "?"
				i++
			}
			query := fmt.Sprintf(
				"INSERT INTO `%s`(%s) VALUES (%s);",
				name,
				strings.Join(insertFields, ", "),
				strings.Join(placeholders, ", "),
			)
			err := DatabaseDriver.ExecuteRaw(query, insertValues...)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
