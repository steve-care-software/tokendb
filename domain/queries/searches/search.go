package searches

import (
	"github.com/steve-care-software/tokendb/domain/queries/commons"
	"github.com/steve-care-software/selector/domain/selectors"
	"github.com/steve-care-software/validator/domain/grammars"
)

type search struct {
	selector selectors.Selector
	matcher  grammars.Grammar
	rnge     commons.Range
}

func createSearch(
	selector selectors.Selector,
	matcher grammars.Grammar,
	rnge commons.Range,
) Search {
	out := search{
		selector: selector,
		matcher:  matcher,
		rnge:     rnge,
	}

	return &out
}

// Selector returns the selector
func (obj *search) Selector() selectors.Selector {
	return obj.selector
}

// Matcher returns the matcher
func (obj *search) Matcher() grammars.Grammar {
	return obj.matcher
}

// Range returns the range
func (obj *search) Range() commons.Range {
	return obj.rnge
}
