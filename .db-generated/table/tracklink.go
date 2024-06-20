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

var TrackLink = newTrackLinkTable("public", "TrackLink", "")

type trackLinkTable struct {
	postgres.Table

	// Columns
	UniqueId   postgres.ColumnString
	CreatedAt  postgres.ColumnTimestamp
	UpdatedAt  postgres.ColumnTimestamp
	CampaignId postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TrackLinkTable struct {
	trackLinkTable

	EXCLUDED trackLinkTable
}

// AS creates new TrackLinkTable with assigned alias
func (a TrackLinkTable) AS(alias string) *TrackLinkTable {
	return newTrackLinkTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TrackLinkTable with assigned schema name
func (a TrackLinkTable) FromSchema(schemaName string) *TrackLinkTable {
	return newTrackLinkTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TrackLinkTable with assigned table prefix
func (a TrackLinkTable) WithPrefix(prefix string) *TrackLinkTable {
	return newTrackLinkTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TrackLinkTable with assigned table suffix
func (a TrackLinkTable) WithSuffix(suffix string) *TrackLinkTable {
	return newTrackLinkTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTrackLinkTable(schemaName, tableName, alias string) *TrackLinkTable {
	return &TrackLinkTable{
		trackLinkTable: newTrackLinkTableImpl(schemaName, tableName, alias),
		EXCLUDED:       newTrackLinkTableImpl("", "excluded", ""),
	}
}

func newTrackLinkTableImpl(schemaName, tableName, alias string) trackLinkTable {
	var (
		UniqueIdColumn   = postgres.StringColumn("UniqueId")
		CreatedAtColumn  = postgres.TimestampColumn("CreatedAt")
		UpdatedAtColumn  = postgres.TimestampColumn("UpdatedAt")
		CampaignIdColumn = postgres.StringColumn("CampaignId")
		allColumns       = postgres.ColumnList{UniqueIdColumn, CreatedAtColumn, UpdatedAtColumn, CampaignIdColumn}
		mutableColumns   = postgres.ColumnList{CreatedAtColumn, UpdatedAtColumn, CampaignIdColumn}
	)

	return trackLinkTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UniqueId:   UniqueIdColumn,
		CreatedAt:  CreatedAtColumn,
		UpdatedAt:  UpdatedAtColumn,
		CampaignId: CampaignIdColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}