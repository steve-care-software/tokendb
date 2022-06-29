package commons

import "errors"

type cursorBuilder struct {
	pIndex     *uint
	comparison Comparison
}

func createCursorBuilder() CursorBuilder {
	out := cursorBuilder{
		pIndex:     nil,
		comparison: nil,
	}

	return &out
}

// Create initializes the builder
func (app *cursorBuilder) Create() CursorBuilder {
	return createCursorBuilder()
}

// WithIndex adds an index to the builder
func (app *cursorBuilder) WithIndex(index uint) CursorBuilder {
	app.pIndex = &index
	return app
}

// WithComparison adds a comparison to the builder
func (app *cursorBuilder) WithComparison(comparison Comparison) CursorBuilder {
	app.comparison = comparison
	return app
}

// Now builds a new Cursor index
func (app *cursorBuilder) Now() (Cursor, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Comparison instance")
	}

	if app.comparison != nil {
		return createCursorWithComparison(*app.pIndex, app.comparison), nil
	}

	return createCursor(*app.pIndex), nil
}
