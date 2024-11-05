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

var Debts = newDebtsTable("public", "debts", "")

type debtsTable struct {
	postgres.Table

	//Columns
	CollectorID postgres.ColumnInteger
	DebtorID    postgres.ColumnInteger
	ChatID      postgres.ColumnInteger
	Amount      postgres.ColumnInteger
	UpdatedAt   postgres.ColumnTimestampz
	CreatedAt   postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type DebtsTable struct {
	debtsTable

	EXCLUDED debtsTable
}

// AS creates new DebtsTable with assigned alias
func (a DebtsTable) AS(alias string) *DebtsTable {
	return newDebtsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new DebtsTable with assigned schema name
func (a DebtsTable) FromSchema(schemaName string) *DebtsTable {
	return newDebtsTable(schemaName, a.TableName(), a.Alias())
}

func newDebtsTable(schemaName, tableName, alias string) *DebtsTable {
	return &DebtsTable{
		debtsTable: newDebtsTableImpl(schemaName, tableName, alias),
		EXCLUDED:   newDebtsTableImpl("", "excluded", ""),
	}
}

func newDebtsTableImpl(schemaName, tableName, alias string) debtsTable {
	var (
		CollectorIDColumn = postgres.IntegerColumn("collector_id")
		DebtorIDColumn    = postgres.IntegerColumn("debtor_id")
		ChatIDColumn      = postgres.IntegerColumn("chat_id")
		AmountColumn      = postgres.IntegerColumn("amount")
		UpdatedAtColumn   = postgres.TimestampzColumn("updated_at")
		CreatedAtColumn   = postgres.TimestampzColumn("created_at")
		allColumns        = postgres.ColumnList{CollectorIDColumn, DebtorIDColumn, ChatIDColumn, AmountColumn, UpdatedAtColumn, CreatedAtColumn}
		mutableColumns    = postgres.ColumnList{AmountColumn, UpdatedAtColumn, CreatedAtColumn}
	)

	return debtsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		CollectorID: CollectorIDColumn,
		DebtorID:    DebtorIDColumn,
		ChatID:      ChatIDColumn,
		Amount:      AmountColumn,
		UpdatedAt:   UpdatedAtColumn,
		CreatedAt:   CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
