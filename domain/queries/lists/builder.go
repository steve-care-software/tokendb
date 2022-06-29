package lists

import (
	"errors"

	"github.com/steve-care-software/tokendb/domain/queries/commons"
	"github.com/steve-care-software/selector/domain/selectors"
)

type builder struct {
	selector selectors.Selector
	rnge     commons.Range
}

func createBuilder() Builder {
	out := builder{
		selector: nil,
		rnge:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithSelector adds a selector to the builder
func (app *builder) WithSelector(selector selectors.Selector) Builder {
	app.selector = selector
	return app
}

// WithRange adds a range to the builder
func (app *builder) WithRange(rnge commons.Range) Builder {
	app.rnge = rnge
	return app
}

// Now builds a new List instance
func (app *builder) Now() (List, error) {
	if app.selector == nil {
		return nil, errors.New("the selector is mandatory in order to build a List instance")
	}

	if app.rnge == nil {
		return nil, errors.New("the range is mandatory in order to build a List instance")
	}

	return createList(app.selector, app.rnge), nil
}
