package postgres

import "github.com/xxxgo/jet/internal/jet"

// Expression is common interface for all expressions.
// Can be Bool, Int, Float, String, Date, Time, Timez, Timestamp or Timestampz expressions.
type Expression = jet.Expression

// BoolExpression interface
type BoolExpression = jet.BoolExpression

// StringExpression interface
type StringExpression = jet.StringExpression

// IntegerExpression interface
type IntegerExpression = jet.IntegerExpression

//FloatExpression is interface
type FloatExpression = jet.FloatExpression

// TimeExpression interface
type TimeExpression = jet.TimeExpression

// TimezExpression interface for 'time with time zone' types
type TimezExpression = jet.TimezExpression

// DateExpression is interface for date types
type DateExpression = jet.DateExpression

// TimestampExpression interface
type TimestampExpression = jet.TimestampExpression

// TimestampzExpression interface
type TimestampzExpression = jet.TimestampzExpression

// BoolExp is bool expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as bool expression.
// Does not add sql cast to generated sql builder output.
var BoolExp = jet.BoolExp

// IntExp is int expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as int expression.
// Does not add sql cast to generated sql builder output.
var IntExp = jet.IntExp

// FloatExp is date expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as float expression.
// Does not add sql cast to generated sql builder output.
var FloatExp = jet.FloatExp

// TimeExp is time expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as time expression.
// Does not add sql cast to generated sql builder output.
var TimeExp = jet.TimeExp

// StringExp is string expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as string expression.
// Does not add sql cast to generated sql builder output.
var StringExp = jet.StringExp

// TimezExp is time with time zone expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as time with time zone expression.
// Does not add sql cast to generated sql builder output.
var TimezExp = jet.TimezExp

// DateExp is date expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as date expression.
// Does not add sql cast to generated sql builder output.
var DateExp = jet.DateExp

// TimestampExp is timestamp expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as timestamp expression.
// Does not add sql cast to generated sql builder output.
var TimestampExp = jet.TimestampExp

// TimestampzExp is timestamp with time zone expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as timestamp with time zone expression.
// Does not add sql cast to generated sql builder output.
var TimestampzExp = jet.TimestampzExp

// Raw can be used for any unsupported functions, operators or expressions.
// For example: Raw("current_database()")
var Raw = jet.Raw

// NewEnumValue creates new named enum value
var NewEnumValue = jet.NewEnumValue
