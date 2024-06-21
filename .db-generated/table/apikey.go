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

var ApiKey = newApiKeyTable("public", "ApiKey", "")

type apiKeyTable struct {
	postgres.Table

	// Columns
	UniqueId       postgres.ColumnString
	CreatedAt      postgres.ColumnTimestamp
	UpdatedAt      postgres.ColumnTimestamp
	MemberId       postgres.ColumnString
	Key            postgres.ColumnString
	OrganizationId postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type ApiKeyTable struct {
	apiKeyTable

	EXCLUDED apiKeyTable
}

// AS creates new ApiKeyTable with assigned alias
func (a ApiKeyTable) AS(alias string) *ApiKeyTable {
	return newApiKeyTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ApiKeyTable with assigned schema name
func (a ApiKeyTable) FromSchema(schemaName string) *ApiKeyTable {
	return newApiKeyTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ApiKeyTable with assigned table prefix
func (a ApiKeyTable) WithPrefix(prefix string) *ApiKeyTable {
	return newApiKeyTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ApiKeyTable with assigned table suffix
func (a ApiKeyTable) WithSuffix(suffix string) *ApiKeyTable {
	return newApiKeyTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newApiKeyTable(schemaName, tableName, alias string) *ApiKeyTable {
	return &ApiKeyTable{
		apiKeyTable: newApiKeyTableImpl(schemaName, tableName, alias),
		EXCLUDED:    newApiKeyTableImpl("", "excluded", ""),
	}
}

func newApiKeyTableImpl(schemaName, tableName, alias string) apiKeyTable {
	var (
		UniqueIdColumn       = postgres.StringColumn("UniqueId")
		CreatedAtColumn      = postgres.TimestampColumn("CreatedAt")
		UpdatedAtColumn      = postgres.TimestampColumn("UpdatedAt")
		MemberIdColumn       = postgres.StringColumn("MemberId")
		KeyColumn            = postgres.StringColumn("Key")
		OrganizationIdColumn = postgres.StringColumn("OrganizationId")
		allColumns           = postgres.ColumnList{UniqueIdColumn, CreatedAtColumn, UpdatedAtColumn, MemberIdColumn, KeyColumn, OrganizationIdColumn}
		mutableColumns       = postgres.ColumnList{CreatedAtColumn, UpdatedAtColumn, MemberIdColumn, KeyColumn, OrganizationIdColumn}
	)

	return apiKeyTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UniqueId:       UniqueIdColumn,
		CreatedAt:      CreatedAtColumn,
		UpdatedAt:      UpdatedAtColumn,
		MemberId:       MemberIdColumn,
		Key:            KeyColumn,
		OrganizationId: OrganizationIdColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
