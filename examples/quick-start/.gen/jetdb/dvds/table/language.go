//
// Code generated by go-jet DO NOT EDIT.
// Generated at Thursday, 08-Aug-19 16:59:58 CEST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/xxxgo/jet/postgres"
)

var Language = newLanguageTable()

type LanguageTable struct {
	postgres.Table

	//Columns
	LanguageID postgres.ColumnInteger
	Name       postgres.ColumnString
	LastUpdate postgres.ColumnTimestamp

	AllColumns     postgres.IColumnList
	MutableColumns postgres.IColumnList
}

// creates new LanguageTable with assigned alias
func (a *LanguageTable) AS(alias string) *LanguageTable {
	aliasTable := newLanguageTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newLanguageTable() *LanguageTable {
	var (
		LanguageIDColumn = postgres.IntegerColumn("language_id")
		NameColumn       = postgres.StringColumn("name")
		LastUpdateColumn = postgres.TimestampColumn("last_update")
	)

	return &LanguageTable{
		Table: postgres.NewTable("dvds", "language", LanguageIDColumn, NameColumn, LastUpdateColumn),

		//Columns
		LanguageID: LanguageIDColumn,
		Name:       NameColumn,
		LastUpdate: LastUpdateColumn,

		AllColumns:     postgres.ColumnList(LanguageIDColumn, NameColumn, LastUpdateColumn),
		MutableColumns: postgres.ColumnList(NameColumn, LastUpdateColumn),
	}
}
