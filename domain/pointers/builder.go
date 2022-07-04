package pointers

import "errors"

type builder struct {
	pIndex *uint
	next   Pointer
}

func createBuilder() Builder {
	out := builder{
		pIndex: nil,
		next:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index uint) Builder {
	app.pIndex = &index
	return app
}

// WithNext adds a next pointer to the builder
func (app *builder) WithNext(next Pointer) Builder {
	app.next = next
	return app
}

// Now builds a new Pointer instance
func (app *builder) Now() (Pointer, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Pointer instance")
	}

	if app.next != nil {
		return createPointerWithNext(*app.pIndex, app.next), nil
	}

	return createPointer(*app.pIndex), nil
}
