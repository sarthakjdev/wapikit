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

var CampaignTag = newCampaignTagTable("public", "CampaignTag", "")

type campaignTagTable struct {
	postgres.Table

	// Columns
	CreatedAt  postgres.ColumnTimestampz
	UpdatedAt  postgres.ColumnTimestampz
	CampaignId postgres.ColumnString
	TagId      postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type CampaignTagTable struct {
	campaignTagTable

	EXCLUDED campaignTagTable
}

// AS creates new CampaignTagTable with assigned alias
func (a CampaignTagTable) AS(alias string) *CampaignTagTable {
	return newCampaignTagTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CampaignTagTable with assigned schema name
func (a CampaignTagTable) FromSchema(schemaName string) *CampaignTagTable {
	return newCampaignTagTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CampaignTagTable with assigned table prefix
func (a CampaignTagTable) WithPrefix(prefix string) *CampaignTagTable {
	return newCampaignTagTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CampaignTagTable with assigned table suffix
func (a CampaignTagTable) WithSuffix(suffix string) *CampaignTagTable {
	return newCampaignTagTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCampaignTagTable(schemaName, tableName, alias string) *CampaignTagTable {
	return &CampaignTagTable{
		campaignTagTable: newCampaignTagTableImpl(schemaName, tableName, alias),
		EXCLUDED:         newCampaignTagTableImpl("", "excluded", ""),
	}
}

func newCampaignTagTableImpl(schemaName, tableName, alias string) campaignTagTable {
	var (
		CreatedAtColumn  = postgres.TimestampzColumn("CreatedAt")
		UpdatedAtColumn  = postgres.TimestampzColumn("UpdatedAt")
		CampaignIdColumn = postgres.StringColumn("CampaignId")
		TagIdColumn      = postgres.StringColumn("TagId")
		allColumns       = postgres.ColumnList{CreatedAtColumn, UpdatedAtColumn, CampaignIdColumn, TagIdColumn}
		mutableColumns   = postgres.ColumnList{CreatedAtColumn, UpdatedAtColumn}
	)

	return campaignTagTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		CreatedAt:  CreatedAtColumn,
		UpdatedAt:  UpdatedAtColumn,
		CampaignId: CampaignIdColumn,
		TagId:      TagIdColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
