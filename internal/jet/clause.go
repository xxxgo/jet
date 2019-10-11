package jet

import (
	"github.com/xxxgo/jet/internal/utils"
)

// Clause interface
type Clause interface {
	Serialize(statementType StatementType, out *SQLBuilder)
}

// ClauseWithProjections interface
type ClauseWithProjections interface {
	Clause

	projections() ProjectionList
}

// ClauseSelect struct
type ClauseSelect struct {
	Distinct    bool
	Projections []Projection
}

func (s *ClauseSelect) projections() ProjectionList {
	return s.Projections
}

// Serialize serializes clause into SQLBuilder
func (s *ClauseSelect) Serialize(statementType StatementType, out *SQLBuilder) {
	out.NewLine()
	out.WriteString("SELECT")

	if s.Distinct {
		out.WriteString("DISTINCT")
	}

	if len(s.Projections) == 0 {
		panic("jet: SELECT clause has to have at least one projection")
	}

	out.WriteProjections(statementType, s.Projections)
}

// ClauseFrom struct
type ClauseFrom struct {
	Table Serializer
}

// Serialize serializes clause into SQLBuilder
func (f *ClauseFrom) Serialize(statementType StatementType, out *SQLBuilder) {
	if f.Table == nil {
		return
	}
	out.NewLine()
	out.WriteString("FROM")

	out.IncreaseIdent()
	f.Table.serialize(statementType, out)
	out.DecreaseIdent()
}

// ClauseWhere struct
type ClauseWhere struct {
	Condition BoolExpression
	Mandatory bool
}

// Serialize serializes clause into SQLBuilder
func (c *ClauseWhere) Serialize(statementType StatementType, out *SQLBuilder) {
	if c.Condition == nil {
		if c.Mandatory {
			panic("jet: WHERE clause not set")
		}
		return
	}
	out.NewLine()
	out.WriteString("WHERE")

	out.IncreaseIdent()
	c.Condition.serialize(statementType, out, noWrap)
	out.DecreaseIdent()
}

// ClauseGroupBy struct
type ClauseGroupBy struct {
	List []GroupByClause
}

// Serialize serializes clause into SQLBuilder
func (c *ClauseGroupBy) Serialize(statementType StatementType, out *SQLBuilder) {
	if len(c.List) == 0 {
		return
	}

	out.NewLine()
	out.WriteString("GROUP BY")

	out.IncreaseIdent()

	for i, c := range c.List {
		if i > 0 {
			out.WriteString(", ")
		}

		if c == nil {
			panic("jet: nil clause in GROUP BY list")
		}

		c.serializeForGroupBy(statementType, out)
	}

	out.DecreaseIdent()
}

// ClauseHaving struct
type ClauseHaving struct {
	Condition BoolExpression
}

// Serialize serializes clause into SQLBuilder
func (c *ClauseHaving) Serialize(statementType StatementType, out *SQLBuilder) {
	if c.Condition == nil {
		return
	}

	out.NewLine()
	out.WriteString("HAVING")

	out.IncreaseIdent()
	c.Condition.serialize(statementType, out, noWrap)
	out.DecreaseIdent()
}

// ClauseOrderBy struct
type ClauseOrderBy struct {
	List        []OrderByClause
	SkipNewLine bool
}

// Serialize serializes clause into SQLBuilder
func (o *ClauseOrderBy) Serialize(statementType StatementType, out *SQLBuilder) {
	if o.List == nil {
		return
	}

	if !o.SkipNewLine {
		out.NewLine()
	}
	out.WriteString("ORDER BY")

	out.IncreaseIdent()

	for i, value := range o.List {
		if i > 0 {
			out.WriteString(", ")
		}

		value.serializeForOrderBy(statementType, out)
	}

	out.DecreaseIdent()
}

// ClauseLimit struct
type ClauseLimit struct {
	Count int64
}

// Serialize serializes clause into SQLBuilder
func (l *ClauseLimit) Serialize(statementType StatementType, out *SQLBuilder) {
	if l.Count >= 0 {
		out.NewLine()
		out.WriteString("LIMIT")
		out.insertParametrizedArgument(l.Count)
	}
}

// ClauseOffset struct
type ClauseOffset struct {
	Count int64
}

// Serialize serializes clause into SQLBuilder
func (o *ClauseOffset) Serialize(statementType StatementType, out *SQLBuilder) {
	if o.Count >= 0 {
		out.NewLine()
		out.WriteString("OFFSET")
		out.insertParametrizedArgument(o.Count)
	}
}

// ClauseFor struct
type ClauseFor struct {
	Lock RowLock
}

// Serialize serializes clause into SQLBuilder
func (f *ClauseFor) Serialize(statementType StatementType, out *SQLBuilder) {
	if f.Lock == nil {
		return
	}

	out.NewLine()
	out.WriteString("FOR")
	f.Lock.serialize(statementType, out)
}

// ClauseSetStmtOperator struct
type ClauseSetStmtOperator struct {
	Operator string
	All      bool
	Selects  []StatementWithProjections
	OrderBy  ClauseOrderBy
	Limit    ClauseLimit
	Offset   ClauseOffset
}

func (s *ClauseSetStmtOperator) projections() ProjectionList {
	if len(s.Selects) > 0 {
		return s.Selects[0].projections()
	}
	return nil
}

// Serialize serializes clause into SQLBuilder
func (s *ClauseSetStmtOperator) Serialize(statementType StatementType, out *SQLBuilder) {
	if len(s.Selects) < 2 {
		panic("jet: UNION Statement must contain at least two SELECT statements")
	}

	for i, selectStmt := range s.Selects {
		out.NewLine()
		if i > 0 {
			out.WriteString(s.Operator)

			if s.All {
				out.WriteString("ALL")
			}
			out.NewLine()
		}

		if selectStmt == nil {
			panic("jet: select statement of '" + s.Operator + "' is nil")
		}

		selectStmt.serialize(statementType, out)
	}

	s.OrderBy.Serialize(statementType, out)
	s.Limit.Serialize(statementType, out)
	s.Offset.Serialize(statementType, out)
}

// ClauseUpdate struct
type ClauseUpdate struct {
	Table SerializerTable
}

// Serialize serializes clause into SQLBuilder
func (u *ClauseUpdate) Serialize(statementType StatementType, out *SQLBuilder) {
	out.NewLine()
	out.WriteString("UPDATE")

	if utils.IsNil(u.Table) {
		panic("jet: table to update is nil")
	}

	u.Table.serialize(statementType, out)
}

// ClauseSet struct
type ClauseSet struct {
	Columns []Column
	Values  []Serializer
}

// Serialize serializes clause into SQLBuilder
func (s *ClauseSet) Serialize(statementType StatementType, out *SQLBuilder) {
	out.NewLine()
	out.WriteString("SET")

	if len(s.Columns) != len(s.Values) {
		panic("jet: mismatch in numbers of columns and values for SET clause")
	}

	out.IncreaseIdent(4)
	for i, column := range s.Columns {
		if i > 0 {
			out.WriteString(", ")
			out.NewLine()
		}

		if column == nil {
			panic("jet: nil column in columns list for SET clause")
		}

		out.WriteString(column.Name())

		out.WriteString(" = ")

		s.Values[i].serialize(UpdateStatementType, out)
	}
	out.DecreaseIdent(4)
}

// ClauseInsert struct
type ClauseInsert struct {
	Table   SerializerTable
	Columns []Column
}

// GetColumns gets list of columns for insert
func (i *ClauseInsert) GetColumns() []Column {
	if len(i.Columns) > 0 {
		return i.Columns
	}

	return i.Table.columns()
}

// Serialize serializes clause into SQLBuilder
func (i *ClauseInsert) Serialize(statementType StatementType, out *SQLBuilder) {
	out.NewLine()
	out.WriteString("INSERT INTO")

	if utils.IsNil(i.Table) {
		panic("jet: table is nil for INSERT clause")
	}

	i.Table.serialize(statementType, out)

	if len(i.Columns) > 0 {
		out.WriteString("(")

		SerializeColumnNames(i.Columns, out)

		out.WriteString(")")
	}
}

// ClauseValuesQuery struct
type ClauseValuesQuery struct {
	ClauseValues
	ClauseQuery
}

// Serialize serializes clause into SQLBuilder
func (v *ClauseValuesQuery) Serialize(statementType StatementType, out *SQLBuilder) {
	if len(v.Rows) == 0 && v.Query == nil {
		panic("jet: VALUES or QUERY has to be specified for INSERT statement")
	}

	if len(v.Rows) > 0 && v.Query != nil {
		panic("jet: VALUES or QUERY has to be specified for INSERT statement")
	}

	v.ClauseValues.Serialize(statementType, out)
	v.ClauseQuery.Serialize(statementType, out)
}

// ClauseValues struct
type ClauseValues struct {
	Rows [][]Serializer
}

// Serialize serializes clause into SQLBuilder
func (v *ClauseValues) Serialize(statementType StatementType, out *SQLBuilder) {
	if len(v.Rows) == 0 {
		return
	}

	out.WriteString("VALUES")

	for rowIndex, row := range v.Rows {
		if rowIndex > 0 {
			out.WriteString(",")
		}

		out.IncreaseIdent()
		out.NewLine()
		out.WriteString("(")

		SerializeClauseList(statementType, row, out)

		out.WriteByte(')')
		out.DecreaseIdent()
	}
}

// ClauseQuery struct
type ClauseQuery struct {
	Query SerializerStatement
}

// Serialize serializes clause into SQLBuilder
func (v *ClauseQuery) Serialize(statementType StatementType, out *SQLBuilder) {
	if v.Query == nil {
		return
	}

	v.Query.serialize(statementType, out)
}

// ClauseDelete struct
type ClauseDelete struct {
	Table SerializerTable
}

// Serialize serializes clause into SQLBuilder
func (d *ClauseDelete) Serialize(statementType StatementType, out *SQLBuilder) {
	out.NewLine()
	out.WriteString("DELETE FROM")

	if d.Table == nil {
		panic("jet: nil table in DELETE clause")
	}

	d.Table.serialize(statementType, out)
}

// ClauseStatementBegin struct
type ClauseStatementBegin struct {
	Name   string
	Tables []SerializerTable
}

// Serialize serializes clause into SQLBuilder
func (d *ClauseStatementBegin) Serialize(statementType StatementType, out *SQLBuilder) {
	out.NewLine()
	out.WriteString(d.Name)

	for i, table := range d.Tables {
		if i > 0 {
			out.WriteString(", ")
		}

		table.serialize(statementType, out)
	}
}

// ClauseOptional struct
type ClauseOptional struct {
	Name      string
	Show      bool
	InNewLine bool
}

// Serialize serializes clause into SQLBuilder
func (d *ClauseOptional) Serialize(statementType StatementType, out *SQLBuilder) {
	if !d.Show {
		return
	}
	if d.InNewLine {
		out.NewLine()
	}
	out.WriteString(d.Name)
}

// ClauseIn struct
type ClauseIn struct {
	LockMode string
}

// Serialize serializes clause into SQLBuilder
func (i *ClauseIn) Serialize(statementType StatementType, out *SQLBuilder) {
	if i.LockMode == "" {
		return
	}

	out.WriteString("IN")
	out.WriteString(string(i.LockMode))
	out.WriteString("MODE")
}

// WindowDefinition struct
type WindowDefinition struct {
	Name   string
	Window Window
}

// ClauseWindow struct
type ClauseWindow struct {
	Definitions []WindowDefinition
}

// Serialize serializes clause into SQLBuilder
func (i *ClauseWindow) Serialize(statementType StatementType, out *SQLBuilder) {
	if len(i.Definitions) == 0 {
		return
	}

	out.NewLine()
	out.WriteString("WINDOW")

	for i, def := range i.Definitions {
		if i > 0 {
			out.WriteString(", ")
		}
		out.WriteString(def.Name)
		out.WriteString("AS")
		if def.Window == nil {
			out.WriteString("()")
			continue
		}
		def.Window.serialize(statementType, out)
	}
}
