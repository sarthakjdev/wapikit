//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var ContactListTag = newContactListTagTable("public", "ContactListTag", "")

type contactListTagTable struct {
	postgres.Table

	// Columns
	CreatedAt     postgres.ColumnTimestampz
	UpdatedAt     postgres.ColumnTimestampz
	ContactListId postgres.ColumnString
	TagId         postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type ContactListTagTable struct {
	contactListTagTable

	EXCLUDED contactListTagTable
}

// AS creates new ContactListTagTable with assigned alias
func (a ContactListTagTable) AS(alias string) *ContactListTagTable {
	return newContactListTagTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ContactListTagTable with assigned schema name
func (a ContactListTagTable) FromSchema(schemaName string) *ContactListTagTable {
	return newContactListTagTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ContactListTagTable with assigned table prefix
func (a ContactListTagTable) WithPrefix(prefix string) *ContactListTagTable {
	return newContactListTagTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ContactListTagTable with assigned table suffix
func (a ContactListTagTable) WithSuffix(suffix string) *ContactListTagTable {
	return newContactListTagTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newContactListTagTable(schemaName, tableName, alias string) *ContactListTagTable {
	return &ContactListTagTable{
		contactListTagTable: newContactListTagTableImpl(schemaName, tableName, alias),
		EXCLUDED:            newContactListTagTableImpl("", "excluded", ""),
	}
}

func newContactListTagTableImpl(schemaName, tableName, alias string) contactListTagTable {
	var (
		CreatedAtColumn     = postgres.TimestampzColumn("CreatedAt")
		UpdatedAtColumn     = postgres.TimestampzColumn("UpdatedAt")
		ContactListIdColumn = postgres.StringColumn("ContactListId")
		TagIdColumn         = postgres.StringColumn("TagId")
		allColumns          = postgres.ColumnList{CreatedAtColumn, UpdatedAtColumn, ContactListIdColumn, TagIdColumn}
		mutableColumns      = postgres.ColumnList{CreatedAtColumn, UpdatedAtColumn}
	)

	return contactListTagTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		CreatedAt:     CreatedAtColumn,
		UpdatedAt:     UpdatedAtColumn,
		ContactListId: ContactListIdColumn,
		TagId:         TagIdColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
