package postgres

import "github.com/xxxgo/jet/internal/jet"

// UNION effectively appends the result of sub-queries(select statements) into single query.
// It eliminates duplicate rows from its result.
func UNION(lhs, rhs jet.StatementWithProjections, selects ...jet.StatementWithProjections) setStatement {
	return newSetStatementImpl(union, false, toSelectList(lhs, rhs, selects...))
}

// UNION_ALL effectively appends the result of sub-queries(select statements) into single query.
// It does not eliminates duplicate rows from its result.
func UNION_ALL(lhs, rhs jet.StatementWithProjections, selects ...jet.StatementWithProjections) setStatement {
	return newSetStatementImpl(union, true, toSelectList(lhs, rhs, selects...))
}

// INTERSECT returns all rows that are in query results.
// It eliminates duplicate rows from its result.
func INTERSECT(lhs, rhs jet.StatementWithProjections, selects ...jet.StatementWithProjections) setStatement {
	return newSetStatementImpl(intersect, false, toSelectList(lhs, rhs, selects...))
}

// INTERSECT_ALL returns all rows that are in query results.
// It does not eliminates duplicate rows from its result.
func INTERSECT_ALL(lhs, rhs jet.StatementWithProjections, selects ...jet.StatementWithProjections) setStatement {
	return newSetStatementImpl(intersect, true, toSelectList(lhs, rhs, selects...))
}

// EXCEPT returns all rows that are in the result of query lhs but not in the result of query rhs.
// It eliminates duplicate rows from its result.
func EXCEPT(lhs, rhs jet.StatementWithProjections) setStatement {
	return newSetStatementImpl(except, false, toSelectList(lhs, rhs))
}

// EXCEPT_ALL returns all rows that are in the result of query lhs but not in the result of query rhs.
// It does not eliminates duplicate rows from its result.
func EXCEPT_ALL(lhs, rhs jet.StatementWithProjections) setStatement {
	return newSetStatementImpl(except, true, toSelectList(lhs, rhs))
}

type setStatement interface {
	setOperators

	ORDER_BY(orderByClauses ...jet.OrderByClause) setStatement

	LIMIT(limit int64) setStatement
	OFFSET(offset int64) setStatement

	AsTable(alias string) SelectTable
}

type setOperators interface {
	Statement
	jet.HasProjections
	Expression

	UNION(rhs SelectStatement) setStatement
	UNION_ALL(rhs SelectStatement) setStatement
	INTERSECT(rhs SelectStatement) setStatement
	INTERSECT_ALL(rhs SelectStatement) setStatement
	EXCEPT(rhs SelectStatement) setStatement
	EXCEPT_ALL(rhs SelectStatement) setStatement
}

type setOperatorsImpl struct {
	parent setOperators
}

func (s *setOperatorsImpl) UNION(rhs SelectStatement) setStatement {
	return UNION(s.parent, rhs)
}

func (s *setOperatorsImpl) UNION_ALL(rhs SelectStatement) setStatement {
	return UNION_ALL(s.parent, rhs)
}

func (s *setOperatorsImpl) INTERSECT(rhs SelectStatement) setStatement {
	return INTERSECT(s.parent, rhs)
}

func (s *setOperatorsImpl) INTERSECT_ALL(rhs SelectStatement) setStatement {
	return INTERSECT_ALL(s.parent, rhs)
}

func (s *setOperatorsImpl) EXCEPT(rhs SelectStatement) setStatement {
	return EXCEPT(s.parent, rhs)
}

func (s *setOperatorsImpl) EXCEPT_ALL(rhs SelectStatement) setStatement {
	return EXCEPT_ALL(s.parent, rhs)
}

type setStatementImpl struct {
	jet.ExpressionStatement

	setOperatorsImpl

	setOperator jet.ClauseSetStmtOperator
}

func newSetStatementImpl(operator string, all bool, selects []jet.StatementWithProjections) setStatement {
	newSetStatement := &setStatementImpl{}
	newSetStatement.ExpressionStatement = jet.NewExpressionStatementImpl(Dialect, jet.SetStatementType, newSetStatement,
		&newSetStatement.setOperator)

	newSetStatement.setOperator.Operator = operator
	newSetStatement.setOperator.All = all
	newSetStatement.setOperator.Selects = selects
	newSetStatement.setOperator.Limit.Count = -1
	newSetStatement.setOperator.Offset.Count = -1

	newSetStatement.setOperatorsImpl.parent = newSetStatement

	return newSetStatement
}

func (s *setStatementImpl) ORDER_BY(orderByClauses ...jet.OrderByClause) setStatement {
	s.setOperator.OrderBy.List = orderByClauses
	return s
}

func (s *setStatementImpl) LIMIT(limit int64) setStatement {
	s.setOperator.Limit.Count = limit
	return s
}

func (s *setStatementImpl) OFFSET(offset int64) setStatement {
	s.setOperator.Offset.Count = offset
	return s
}

func (s *setStatementImpl) AsTable(alias string) SelectTable {
	return newSelectTable(s, alias)
}

const (
	union     = "UNION"
	intersect = "INTERSECT"
	except    = "EXCEPT"
)

func toSelectList(lhs, rhs jet.StatementWithProjections, selects ...jet.StatementWithProjections) []jet.StatementWithProjections {
	return append([]jet.StatementWithProjections{lhs, rhs}, selects...)
}
