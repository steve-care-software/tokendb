package commons

import "errors"

type rangeBuilder struct {
	cursor Cursor
	amount uint
}

func createRangeBuilder() RangeBuilder {
	out := rangeBuilder{
		cursor: nil,
		amount: 0,
	}

	return &out
}

// Create initializes the builder
func (app *rangeBuilder) Create() RangeBuilder {
	return createRangeBuilder()
}

// WithCursor adds a cursor to the builder
func (app *rangeBuilder) WithCursor(cursor Cursor) RangeBuilder {
	app.cursor = cursor
	return app
}

// WithAmount adds an amount to the builder
func (app *rangeBuilder) WithAmount(amount uint) RangeBuilder {
	app.amount = amount
	return app
}

// Now builds a new Range instance
func (app *rangeBuilder) Now() (Range, error) {
	if app.cursor == nil {
		return nil, errors.New("the cursor is mandatory in order to build a Range instance")
	}

	if app.amount <= 0 {
		return nil, errors.New("the amount must be greater than zero (0) in order to build a Range instance")
	}

	return createRange(app.cursor, app.amount), nil
}
