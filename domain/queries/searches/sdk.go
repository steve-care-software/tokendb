package searches

import (
	"github.com/steve-care-software/selector/domain/selectors"
	"github.com/steve-care-software/tokendb/domain/queries/commons"
	"github.com/steve-care-software/validator/domain/grammars"
)

// NewBuilder initializes the builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a search builder
type Builder interface {
	Create() Builder
	WithSelector(selector selectors.Selector) Builder
	WithMatcher(matcher grammars.Grammar) Builder
	WithRange(rnge commons.Range) Builder
	Now() (Search, error)
}

// Search represents a search query
type Search interface {
	Selector() selectors.Selector
	Matcher() grammars.Grammar
	Range() commons.Range
}
