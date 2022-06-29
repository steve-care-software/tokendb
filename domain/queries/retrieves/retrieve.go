package retrieves

import (
	"github.com/steve-care-software/tokendb/domain/queries/commons"
	"github.com/steve-care-software/selector/domain/selectors"
)

type retrieve struct {
	selector selectors.Selector
	cursor   commons.Cursor
}

func createRetrieve(
	selector selectors.Selector,
	cursor commons.Cursor,
) Retrieve {
	out := retrieve{
		selector: selector,
		cursor:   cursor,
	}

	return &out
}

// Selector returns the selector
func (obj *retrieve) Selector() selectors.Selector {
	return obj.selector
}

// Cursor returns the cursor
func (obj *retrieve) Cursor() commons.Cursor {
	return obj.cursor
}
