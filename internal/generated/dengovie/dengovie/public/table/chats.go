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

var Chats = newChatsTable("public", "chats", "")

type chatsTable struct {
	postgres.Table

	//Columns
	ChatID      postgres.ColumnInteger
	TgChatID    postgres.ColumnString
	Description postgres.ColumnString
	CreatedAt   postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type ChatsTable struct {
	chatsTable

	EXCLUDED chatsTable
}

// AS creates new ChatsTable with assigned alias
func (a ChatsTable) AS(alias string) *ChatsTable {
	return newChatsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ChatsTable with assigned schema name
func (a ChatsTable) FromSchema(schemaName string) *ChatsTable {
	return newChatsTable(schemaName, a.TableName(), a.Alias())
}

func newChatsTable(schemaName, tableName, alias string) *ChatsTable {
	return &ChatsTable{
		chatsTable: newChatsTableImpl(schemaName, tableName, alias),
		EXCLUDED:   newChatsTableImpl("", "excluded", ""),
	}
}

func newChatsTableImpl(schemaName, tableName, alias string) chatsTable {
	var (
		ChatIDColumn      = postgres.IntegerColumn("chat_id")
		TgChatIDColumn    = postgres.StringColumn("tg_chat_id")
		DescriptionColumn = postgres.StringColumn("description")
		CreatedAtColumn   = postgres.TimestampzColumn("created_at")
		allColumns        = postgres.ColumnList{ChatIDColumn, TgChatIDColumn, DescriptionColumn, CreatedAtColumn}
		mutableColumns    = postgres.ColumnList{TgChatIDColumn, DescriptionColumn, CreatedAtColumn}
	)

	return chatsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ChatID:      ChatIDColumn,
		TgChatID:    TgChatIDColumn,
		Description: DescriptionColumn,
		CreatedAt:   CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
