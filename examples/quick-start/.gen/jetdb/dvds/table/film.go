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

var Film = newFilmTable()

type FilmTable struct {
	postgres.Table

	//Columns
	FilmID          postgres.ColumnInteger
	Title           postgres.ColumnString
	Description     postgres.ColumnString
	ReleaseYear     postgres.ColumnInteger
	LanguageID      postgres.ColumnInteger
	RentalDuration  postgres.ColumnInteger
	RentalRate      postgres.ColumnFloat
	Length          postgres.ColumnInteger
	ReplacementCost postgres.ColumnFloat
	Rating          postgres.ColumnString
	LastUpdate      postgres.ColumnTimestamp
	SpecialFeatures postgres.ColumnString
	Fulltext        postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new FilmTable with assigned alias
func (a *FilmTable) AS(alias string) *FilmTable {
	aliasTable := newFilmTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newFilmTable() *FilmTable {
	var (
		FilmIDColumn          = postgres.IntegerColumn("film_id")
		TitleColumn           = postgres.StringColumn("title")
		DescriptionColumn     = postgres.StringColumn("description")
		ReleaseYearColumn     = postgres.IntegerColumn("release_year")
		LanguageIDColumn      = postgres.IntegerColumn("language_id")
		RentalDurationColumn  = postgres.IntegerColumn("rental_duration")
		RentalRateColumn      = postgres.FloatColumn("rental_rate")
		LengthColumn          = postgres.IntegerColumn("length")
		ReplacementCostColumn = postgres.FloatColumn("replacement_cost")
		RatingColumn          = postgres.StringColumn("rating")
		LastUpdateColumn      = postgres.TimestampColumn("last_update")
		SpecialFeaturesColumn = postgres.StringColumn("special_features")
		FulltextColumn        = postgres.StringColumn("fulltext")
	)

	return &FilmTable{
		Table: postgres.NewTable("dvds", "film", FilmIDColumn, TitleColumn, DescriptionColumn, ReleaseYearColumn, LanguageIDColumn, RentalDurationColumn, RentalRateColumn, LengthColumn, ReplacementCostColumn, RatingColumn, LastUpdateColumn, SpecialFeaturesColumn, FulltextColumn),

		//Columns
		FilmID:          FilmIDColumn,
		Title:           TitleColumn,
		Description:     DescriptionColumn,
		ReleaseYear:     ReleaseYearColumn,
		LanguageID:      LanguageIDColumn,
		RentalDuration:  RentalDurationColumn,
		RentalRate:      RentalRateColumn,
		Length:          LengthColumn,
		ReplacementCost: ReplacementCostColumn,
		Rating:          RatingColumn,
		LastUpdate:      LastUpdateColumn,
		SpecialFeatures: SpecialFeaturesColumn,
		Fulltext:        FulltextColumn,

		AllColumns:     postgres.ColumnList{FilmIDColumn, TitleColumn, DescriptionColumn, ReleaseYearColumn, LanguageIDColumn, RentalDurationColumn, RentalRateColumn, LengthColumn, ReplacementCostColumn, RatingColumn, LastUpdateColumn, SpecialFeaturesColumn, FulltextColumn},
		MutableColumns: postgres.ColumnList{TitleColumn, DescriptionColumn, ReleaseYearColumn, LanguageIDColumn, RentalDurationColumn, RentalRateColumn, LengthColumn, ReplacementCostColumn, RatingColumn, LastUpdateColumn, SpecialFeaturesColumn, FulltextColumn},
	}
}
