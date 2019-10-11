package mysql

import "github.com/xxxgo/jet/internal/jet"

// Statement is common interface for all statements(SELECT, INSERT, UPDATE, DELETE, LOCK)
type Statement = jet.Statement

// Projection is interface for all projection types. Types that can be part of, for instance SELECT clause.
type Projection = jet.Projection
