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

var Events = newEventsTable("public", "events", "")

type eventsTable struct {
	postgres.Table

	//Columns
	EventID     postgres.ColumnInteger
	CollectorID postgres.ColumnInteger
	DebtorID    postgres.ColumnInteger
	ChatID      postgres.ColumnInteger
	Amount      postgres.ColumnInteger
	Description postgres.ColumnString
	CreatedAt   postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type EventsTable struct {
	eventsTable

	EXCLUDED eventsTable
}

// AS creates new EventsTable with assigned alias
func (a EventsTable) AS(alias string) *EventsTable {
	return newEventsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new EventsTable with assigned schema name
func (a EventsTable) FromSchema(schemaName string) *EventsTable {
	return newEventsTable(schemaName, a.TableName(), a.Alias())
}

func newEventsTable(schemaName, tableName, alias string) *EventsTable {
	return &EventsTable{
		eventsTable: newEventsTableImpl(schemaName, tableName, alias),
		EXCLUDED:    newEventsTableImpl("", "excluded", ""),
	}
}

func newEventsTableImpl(schemaName, tableName, alias string) eventsTable {
	var (
		EventIDColumn     = postgres.IntegerColumn("event_id")
		CollectorIDColumn = postgres.IntegerColumn("collector_id")
		DebtorIDColumn    = postgres.IntegerColumn("debtor_id")
		ChatIDColumn      = postgres.IntegerColumn("chat_id")
		AmountColumn      = postgres.IntegerColumn("amount")
		DescriptionColumn = postgres.StringColumn("description")
		CreatedAtColumn   = postgres.TimestampzColumn("created_at")
		allColumns        = postgres.ColumnList{EventIDColumn, CollectorIDColumn, DebtorIDColumn, ChatIDColumn, AmountColumn, DescriptionColumn, CreatedAtColumn}
		mutableColumns    = postgres.ColumnList{CollectorIDColumn, DebtorIDColumn, ChatIDColumn, AmountColumn, DescriptionColumn, CreatedAtColumn}
	)

	return eventsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		EventID:     EventIDColumn,
		CollectorID: CollectorIDColumn,
		DebtorID:    DebtorIDColumn,
		ChatID:      ChatIDColumn,
		Amount:      AmountColumn,
		Description: DescriptionColumn,
		CreatedAt:   CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
