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

var OrganizationMember = newOrganizationMemberTable("public", "OrganizationMember", "")

type organizationMemberTable struct {
	postgres.Table

	// Columns
	UniqueId       postgres.ColumnString
	CreatedAt      postgres.ColumnTimestamp
	UpdatedAt      postgres.ColumnTimestamp
	Role           postgres.ColumnString
	OrganizationId postgres.ColumnString
	UserId         postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type OrganizationMemberTable struct {
	organizationMemberTable

	EXCLUDED organizationMemberTable
}

// AS creates new OrganizationMemberTable with assigned alias
func (a OrganizationMemberTable) AS(alias string) *OrganizationMemberTable {
	return newOrganizationMemberTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new OrganizationMemberTable with assigned schema name
func (a OrganizationMemberTable) FromSchema(schemaName string) *OrganizationMemberTable {
	return newOrganizationMemberTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new OrganizationMemberTable with assigned table prefix
func (a OrganizationMemberTable) WithPrefix(prefix string) *OrganizationMemberTable {
	return newOrganizationMemberTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new OrganizationMemberTable with assigned table suffix
func (a OrganizationMemberTable) WithSuffix(suffix string) *OrganizationMemberTable {
	return newOrganizationMemberTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newOrganizationMemberTable(schemaName, tableName, alias string) *OrganizationMemberTable {
	return &OrganizationMemberTable{
		organizationMemberTable: newOrganizationMemberTableImpl(schemaName, tableName, alias),
		EXCLUDED:                newOrganizationMemberTableImpl("", "excluded", ""),
	}
}

func newOrganizationMemberTableImpl(schemaName, tableName, alias string) organizationMemberTable {
	var (
		UniqueIdColumn       = postgres.StringColumn("UniqueId")
		CreatedAtColumn      = postgres.TimestampColumn("CreatedAt")
		UpdatedAtColumn      = postgres.TimestampColumn("UpdatedAt")
		RoleColumn           = postgres.StringColumn("Role")
		OrganizationIdColumn = postgres.StringColumn("OrganizationId")
		UserIdColumn         = postgres.StringColumn("UserId")
		allColumns           = postgres.ColumnList{UniqueIdColumn, CreatedAtColumn, UpdatedAtColumn, RoleColumn, OrganizationIdColumn, UserIdColumn}
		mutableColumns       = postgres.ColumnList{CreatedAtColumn, UpdatedAtColumn, RoleColumn, OrganizationIdColumn, UserIdColumn}
	)

	return organizationMemberTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UniqueId:       UniqueIdColumn,
		CreatedAt:      CreatedAtColumn,
		UpdatedAt:      UpdatedAtColumn,
		Role:           RoleColumn,
		OrganizationId: OrganizationIdColumn,
		UserId:         UserIdColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
