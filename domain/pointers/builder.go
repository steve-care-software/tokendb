package pointers

import (
	"crypto/sha512"
	"errors"
	"strconv"
)

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

	data := []byte{}
	data = append(data, []byte(strconv.Itoa(int(*app.pIndex)))...)
	if app.next != nil {
		nextHash := app.next.Hash()
		data = append(data, nextHash...)
	}

	hash := sha512.New()
	_, err := hash.Write([]byte(data))
	if err != nil {
		return nil, err
	}

	if app.next != nil {
		return createPointerWithNext(hash.Sum(nil), *app.pIndex, app.next), nil
	}

	return createPointer(hash.Sum(nil), *app.pIndex), nil
}
