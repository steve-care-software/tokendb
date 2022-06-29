package lists

import (
	"github.com/steve-care-software/tokendb/domain/queries/commons"
	"github.com/steve-care-software/selector/domain/selectors"
)

type list struct {
	selector selectors.Selector
	rnge     commons.Range
}

func createList(
	selector selectors.Selector,
	rnge commons.Range,
) List {
	out := list{
		selector: selector,
		rnge:     rnge,
	}

	return &out
}

// Selector returns the selectors
func (obj *list) Selector() selectors.Selector {
	return obj.selector
}

// Range returns the range
func (obj *list) Range() commons.Range {
	return obj.rnge
}
