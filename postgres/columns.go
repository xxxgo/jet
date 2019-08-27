package postgres

import "github.com/xxxgo/jet/v2/internal/jet"

// Column is common column interface for all types of columns.
type Column jet.ColumnExpression

// IColumnList is used to store list of columns for later reuse as single projection or
// column list for UPDATE and INSERT statement.
type IColumnList jet.IColumnList

// ColumnList function returns list of columns that be used as projection or column list for UPDATE and INSERT statement.
var ColumnList = jet.ColumnList

// ColumnBool is interface for SQL boolean columns.
type ColumnBool jet.ColumnBool

// BoolColumn creates named bool column.
var BoolColumn = jet.BoolColumn

// ColumnString is interface for SQL text, character, character varying
// bytea, uuid columns and enums types.
type ColumnString jet.ColumnString

// StringColumn creates named string column.
var StringColumn = jet.StringColumn

// ColumnInteger is interface for SQL smallint, integer, bigint columns.
type ColumnInteger jet.ColumnInteger

// IntegerColumn creates named integer column.
var IntegerColumn = jet.IntegerColumn

// ColumnFloat is interface for SQL real, numeric, decimal or double precision column.
type ColumnFloat jet.ColumnFloat

// FloatColumn creates named float column.
var FloatColumn = jet.FloatColumn

// ColumnDate is interface of SQL date columns.
type ColumnDate jet.ColumnDate

// DateColumn creates named date column.
var DateColumn = jet.DateColumn

// ColumnTime is interface for SQL time column.
type ColumnTime jet.ColumnTime

// TimeColumn creates named time column
var TimeColumn = jet.TimeColumn

// ColumnTimez is interface of SQL time with time zone columns.
type ColumnTimez jet.ColumnTimez

// TimezColumn creates named time with time zone column.
var TimezColumn = jet.TimezColumn

// ColumnTimestamp is interface of SQL timestamp columns.
type ColumnTimestamp jet.ColumnTimestamp

// TimestampColumn creates named timestamp column
var TimestampColumn = jet.TimestampColumn

// ColumnTimestampz is interface of SQL timestamp with timezone columns.
type ColumnTimestampz jet.ColumnTimestampz

// TimestampzColumn creates named timestamp with time zone column.
var TimestampzColumn = jet.TimestampzColumn
