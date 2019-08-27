package postgres

import "github.com/xxxgo/jet/v2/internal/jet"

// RowLock is interface for SELECT statement row lock types
type RowLock = jet.RowLock

// Row lock types
var (
	UPDATE        = jet.NewRowLock("UPDATE")
	NO_KEY_UPDATE = jet.NewRowLock("NO KEY UPDATE")
	SHARE         = jet.NewRowLock("SHARE")
	KEY_SHARE     = jet.NewRowLock("KEY SHARE")
)

// SelectStatement is interface for PostgreSQL SELECT statement
type SelectStatement interface {
	Statement
	jet.HasProjections
	Expression

	DISTINCT() SelectStatement
	FROM(table ReadableTable) SelectStatement
	WHERE(expression BoolExpression) SelectStatement
	GROUP_BY(groupByClauses ...jet.GroupByClause) SelectStatement
	HAVING(boolExpression BoolExpression) SelectStatement
	ORDER_BY(orderByClauses ...jet.OrderByClause) SelectStatement
	LIMIT(limit int64) SelectStatement
	OFFSET(offset int64) SelectStatement
	FOR(lock RowLock) SelectStatement

	UNION(rhs SelectStatement) setStatement
	UNION_ALL(rhs SelectStatement) setStatement
	INTERSECT(rhs SelectStatement) setStatement
	INTERSECT_ALL(rhs SelectStatement) setStatement
	EXCEPT(rhs SelectStatement) setStatement
	EXCEPT_ALL(rhs SelectStatement) setStatement

	AsTable(alias string) SelectTable
}

//SELECT creates new SelectStatement with list of projections
func SELECT(projection Projection, projections ...Projection) SelectStatement {
	return newSelectStatement(nil, append([]Projection{projection}, projections...))
}

func newSelectStatement(table ReadableTable, projections []Projection) SelectStatement {
	newSelect := &selectStatementImpl{}
	newSelect.ExpressionStatement = jet.NewExpressionStatementImpl(Dialect, jet.SelectStatementType, newSelect, &newSelect.Select,
		&newSelect.From, &newSelect.Where, &newSelect.GroupBy, &newSelect.Having, &newSelect.OrderBy,
		&newSelect.Limit, &newSelect.Offset, &newSelect.For)

	//	statementImpl = jet.NewStatementImpl(Dialect, jet.SelectStatementType, newSelect, &newSelect.Select,
	//	&newSelect.From, &newSelect.Where, &newSelect.GroupBy, &newSelect.Having, &newSelect.OrderBy,
	//	&newSelect.Limit, &newSelect.Offset, &newSelect.For)
	//
	//newSelect.expressionStatementImpl.expressionInterfaceImpl.Parent = newSelect

	newSelect.Select.Projections = toJetProjectionList(projections)
	newSelect.From.Table = table
	newSelect.Limit.Count = -1
	newSelect.Offset.Count = -1

	newSelect.setOperatorsImpl.parent = newSelect

	return newSelect
}

type selectStatementImpl struct {
	jet.ExpressionStatement
	setOperatorsImpl

	Select  jet.ClauseSelect
	From    jet.ClauseFrom
	Where   jet.ClauseWhere
	GroupBy jet.ClauseGroupBy
	Having  jet.ClauseHaving
	OrderBy jet.ClauseOrderBy
	Limit   jet.ClauseLimit
	Offset  jet.ClauseOffset
	For     jet.ClauseFor
}

func (s *selectStatementImpl) DISTINCT() SelectStatement {
	s.Select.Distinct = true
	return s
}

func (s *selectStatementImpl) FROM(table ReadableTable) SelectStatement {
	s.From.Table = table
	return s
}

func (s *selectStatementImpl) WHERE(condition BoolExpression) SelectStatement {
	s.Where.Condition = condition
	return s
}

func (s *selectStatementImpl) GROUP_BY(groupByClauses ...jet.GroupByClause) SelectStatement {
	s.GroupBy.List = groupByClauses
	return s
}

func (s *selectStatementImpl) HAVING(boolExpression BoolExpression) SelectStatement {
	s.Having.Condition = boolExpression
	return s
}

func (s *selectStatementImpl) ORDER_BY(orderByClauses ...jet.OrderByClause) SelectStatement {
	s.OrderBy.List = orderByClauses
	return s
}

func (s *selectStatementImpl) LIMIT(limit int64) SelectStatement {
	s.Limit.Count = limit
	return s
}

func (s *selectStatementImpl) OFFSET(offset int64) SelectStatement {
	s.Offset.Count = offset
	return s
}

func (s *selectStatementImpl) FOR(lock RowLock) SelectStatement {
	s.For.Lock = lock
	return s
}

func (s *selectStatementImpl) AsTable(alias string) SelectTable {
	return newSelectTable(s, alias)
}
