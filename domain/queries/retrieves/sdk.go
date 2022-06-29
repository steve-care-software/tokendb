package retrieves

import (
	"github.com/steve-care-software/selector/domain/selectors"
	"github.com/steve-care-software/tokendb/domain/queries/commons"
)

// NewBuilder initializes the builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a retrieve builder
type Builder interface {
	Create() Builder
	WithSelector(selector selectors.Selector) Builder
	WithCursor(cursor commons.Cursor) Builder
	Now() (Retrieve, error)
}

// Retrieve represents a retrieval query
type Retrieve interface {
	Selector() selectors.Selector
	Cursor() commons.Cursor
}
