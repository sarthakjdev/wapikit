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

var OrganisationRole = newOrganisationRoleTable("public", "OrganisationRole", "")

type organisationRoleTable struct {
	postgres.Table

	// Columns
	UniqueId       postgres.ColumnString
	CreatedAt      postgres.ColumnTimestamp
	UpdatedAt      postgres.ColumnTimestamp
	Name           postgres.ColumnString
	OrganisationId postgres.ColumnString
	Permissions    postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type OrganisationRoleTable struct {
	organisationRoleTable

	EXCLUDED organisationRoleTable
}

// AS creates new OrganisationRoleTable with assigned alias
func (a OrganisationRoleTable) AS(alias string) *OrganisationRoleTable {
	return newOrganisationRoleTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new OrganisationRoleTable with assigned schema name
func (a OrganisationRoleTable) FromSchema(schemaName string) *OrganisationRoleTable {
	return newOrganisationRoleTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new OrganisationRoleTable with assigned table prefix
func (a OrganisationRoleTable) WithPrefix(prefix string) *OrganisationRoleTable {
	return newOrganisationRoleTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new OrganisationRoleTable with assigned table suffix
func (a OrganisationRoleTable) WithSuffix(suffix string) *OrganisationRoleTable {
	return newOrganisationRoleTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newOrganisationRoleTable(schemaName, tableName, alias string) *OrganisationRoleTable {
	return &OrganisationRoleTable{
		organisationRoleTable: newOrganisationRoleTableImpl(schemaName, tableName, alias),
		EXCLUDED:              newOrganisationRoleTableImpl("", "excluded", ""),
	}
}

func newOrganisationRoleTableImpl(schemaName, tableName, alias string) organisationRoleTable {
	var (
		UniqueIdColumn       = postgres.StringColumn("UniqueId")
		CreatedAtColumn      = postgres.TimestampColumn("CreatedAt")
		UpdatedAtColumn      = postgres.TimestampColumn("UpdatedAt")
		NameColumn           = postgres.StringColumn("Name")
		OrganisationIdColumn = postgres.StringColumn("OrganisationId")
		PermissionsColumn    = postgres.StringColumn("Permissions")
		allColumns           = postgres.ColumnList{UniqueIdColumn, CreatedAtColumn, UpdatedAtColumn, NameColumn, OrganisationIdColumn, PermissionsColumn}
		mutableColumns       = postgres.ColumnList{CreatedAtColumn, UpdatedAtColumn, NameColumn, OrganisationIdColumn, PermissionsColumn}
	)

	return organisationRoleTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UniqueId:       UniqueIdColumn,
		CreatedAt:      CreatedAtColumn,
		UpdatedAt:      UpdatedAtColumn,
		Name:           NameColumn,
		OrganisationId: OrganisationIdColumn,
		Permissions:    PermissionsColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}