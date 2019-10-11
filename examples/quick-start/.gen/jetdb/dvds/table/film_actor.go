//
// Code generated by go-jet DO NOT EDIT.
// Generated at Thursday, 26-Sep-19 12:02:13 CEST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/xxxgo/jet/postgres"
)

var FilmActor = newFilmActorTable()

type FilmActorTable struct {
	postgres.Table

	//Columns
	ActorID    postgres.ColumnInteger
	FilmID     postgres.ColumnInteger
	LastUpdate postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new FilmActorTable with assigned alias
func (a *FilmActorTable) AS(alias string) *FilmActorTable {
	aliasTable := newFilmActorTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newFilmActorTable() *FilmActorTable {
	var (
		ActorIDColumn    = postgres.IntegerColumn("actor_id")
		FilmIDColumn     = postgres.IntegerColumn("film_id")
		LastUpdateColumn = postgres.TimestampColumn("last_update")
	)

	return &FilmActorTable{
		Table: postgres.NewTable("dvds", "film_actor", ActorIDColumn, FilmIDColumn, LastUpdateColumn),

		//Columns
		ActorID:    ActorIDColumn,
		FilmID:     FilmIDColumn,
		LastUpdate: LastUpdateColumn,

		AllColumns:     postgres.ColumnList{ActorIDColumn, FilmIDColumn, LastUpdateColumn},
		MutableColumns: postgres.ColumnList{LastUpdateColumn},
	}
}
