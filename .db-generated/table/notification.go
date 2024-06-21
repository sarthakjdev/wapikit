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

var Notification = newNotificationTable("public", "Notification", "")

type notificationTable struct {
	postgres.Table

	// Columns
	UniqueId    postgres.ColumnString
	CreatedAt   postgres.ColumnTimestamp
	UpdatedAt   postgres.ColumnTimestamp
	CtaUrl      postgres.ColumnString
	Title       postgres.ColumnString
	Description postgres.ColumnString
	Type        postgres.ColumnString
	IsBroadcast postgres.ColumnBool
	UserId      postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type NotificationTable struct {
	notificationTable

	EXCLUDED notificationTable
}

// AS creates new NotificationTable with assigned alias
func (a NotificationTable) AS(alias string) *NotificationTable {
	return newNotificationTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new NotificationTable with assigned schema name
func (a NotificationTable) FromSchema(schemaName string) *NotificationTable {
	return newNotificationTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new NotificationTable with assigned table prefix
func (a NotificationTable) WithPrefix(prefix string) *NotificationTable {
	return newNotificationTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new NotificationTable with assigned table suffix
func (a NotificationTable) WithSuffix(suffix string) *NotificationTable {
	return newNotificationTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newNotificationTable(schemaName, tableName, alias string) *NotificationTable {
	return &NotificationTable{
		notificationTable: newNotificationTableImpl(schemaName, tableName, alias),
		EXCLUDED:          newNotificationTableImpl("", "excluded", ""),
	}
}

func newNotificationTableImpl(schemaName, tableName, alias string) notificationTable {
	var (
		UniqueIdColumn    = postgres.StringColumn("UniqueId")
		CreatedAtColumn   = postgres.TimestampColumn("CreatedAt")
		UpdatedAtColumn   = postgres.TimestampColumn("UpdatedAt")
		CtaUrlColumn      = postgres.StringColumn("ctaUrl")
		TitleColumn       = postgres.StringColumn("title")
		DescriptionColumn = postgres.StringColumn("description")
		TypeColumn        = postgres.StringColumn("type")
		IsBroadcastColumn = postgres.BoolColumn("isBroadcast")
		UserIdColumn      = postgres.StringColumn("UserId")
		allColumns        = postgres.ColumnList{UniqueIdColumn, CreatedAtColumn, UpdatedAtColumn, CtaUrlColumn, TitleColumn, DescriptionColumn, TypeColumn, IsBroadcastColumn, UserIdColumn}
		mutableColumns    = postgres.ColumnList{CreatedAtColumn, UpdatedAtColumn, CtaUrlColumn, TitleColumn, DescriptionColumn, TypeColumn, IsBroadcastColumn, UserIdColumn}
	)

	return notificationTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UniqueId:    UniqueIdColumn,
		CreatedAt:   CreatedAtColumn,
		UpdatedAt:   UpdatedAtColumn,
		CtaUrl:      CtaUrlColumn,
		Title:       TitleColumn,
		Description: DescriptionColumn,
		Type:        TypeColumn,
		IsBroadcast: IsBroadcastColumn,
		UserId:      UserIdColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
