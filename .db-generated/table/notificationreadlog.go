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

var NotificationReadLog = newNotificationReadLogTable("public", "NotificationReadLog", "")

type notificationReadLogTable struct {
	postgres.Table

	// Columns
	UniqueId       postgres.ColumnString
	CreatedAt      postgres.ColumnTimestamp
	UpdatedAt      postgres.ColumnTimestamp
	ReadByUserId   postgres.ColumnString
	NotificationId postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type NotificationReadLogTable struct {
	notificationReadLogTable

	EXCLUDED notificationReadLogTable
}

// AS creates new NotificationReadLogTable with assigned alias
func (a NotificationReadLogTable) AS(alias string) *NotificationReadLogTable {
	return newNotificationReadLogTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new NotificationReadLogTable with assigned schema name
func (a NotificationReadLogTable) FromSchema(schemaName string) *NotificationReadLogTable {
	return newNotificationReadLogTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new NotificationReadLogTable with assigned table prefix
func (a NotificationReadLogTable) WithPrefix(prefix string) *NotificationReadLogTable {
	return newNotificationReadLogTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new NotificationReadLogTable with assigned table suffix
func (a NotificationReadLogTable) WithSuffix(suffix string) *NotificationReadLogTable {
	return newNotificationReadLogTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newNotificationReadLogTable(schemaName, tableName, alias string) *NotificationReadLogTable {
	return &NotificationReadLogTable{
		notificationReadLogTable: newNotificationReadLogTableImpl(schemaName, tableName, alias),
		EXCLUDED:                 newNotificationReadLogTableImpl("", "excluded", ""),
	}
}

func newNotificationReadLogTableImpl(schemaName, tableName, alias string) notificationReadLogTable {
	var (
		UniqueIdColumn       = postgres.StringColumn("UniqueId")
		CreatedAtColumn      = postgres.TimestampColumn("CreatedAt")
		UpdatedAtColumn      = postgres.TimestampColumn("UpdatedAt")
		ReadByUserIdColumn   = postgres.StringColumn("ReadByUserId")
		NotificationIdColumn = postgres.StringColumn("NotificationId")
		allColumns           = postgres.ColumnList{UniqueIdColumn, CreatedAtColumn, UpdatedAtColumn, ReadByUserIdColumn, NotificationIdColumn}
		mutableColumns       = postgres.ColumnList{CreatedAtColumn, UpdatedAtColumn, ReadByUserIdColumn, NotificationIdColumn}
	)

	return notificationReadLogTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UniqueId:       UniqueIdColumn,
		CreatedAt:      CreatedAtColumn,
		UpdatedAt:      UpdatedAtColumn,
		ReadByUserId:   ReadByUserIdColumn,
		NotificationId: NotificationIdColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
