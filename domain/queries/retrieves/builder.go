package retrieves

import (
	"errors"

	"github.com/steve-care-software/tokendb/domain/queries/commons"
	"github.com/steve-care-software/selector/domain/selectors"
)

type builder struct {
	selector selectors.Selector
	cursor   commons.Cursor
}

func createBuilder() Builder {
	out := builder{
		selector: nil,
		cursor:   nil,
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

// WithCursor adds a cursor to the builder
func (app *builder) WithCursor(cursor commons.Cursor) Builder {
	app.cursor = cursor
	return app
}

// Now builds a new Retrieve instance
func (app *builder) Now() (Retrieve, error) {
	if app.selector == nil {
		return nil, errors.New("the selector is mandatory in order to build a Retrieve instance")
	}

	if app.cursor == nil {
		return nil, errors.New("the cursor is mandatory in order to build a Retrieve instance")
	}

	return createRetrieve(app.selector, app.cursor), nil
}
