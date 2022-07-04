package states

import (
	"crypto/sha512"
	"errors"
	"strconv"
	"time"

	"github.com/steve-care-software/tokendb/domain/pointers"
)

type builder struct {
	hash       []byte
	pHeight    *uint
	pointer    pointers.Pointer
	pCreatedOn *time.Time
	previous   []byte
}

func createBuilder() Builder {
	out := builder{
		hash:       nil,
		pHeight:    nil,
		pointer:    nil,
		pCreatedOn: nil,
		previous:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithHeight adds an height to the builder
func (app *builder) WithHeight(height uint) Builder {
	app.pHeight = &height
	return app
}

// WithPrevious adds previous to the builder
func (app *builder) WithPrevious(previous []byte) Builder {
	app.previous = previous
	return app
}

// WithPointer adds a pointer to the builder
func (app *builder) WithPointer(pointer pointers.Pointer) Builder {
	app.pointer = pointer
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.pCreatedOn = &createdOn
	return app
}

// Now buildsa new State instance
func (app *builder) Now() (State, error) {
	if app.previous == nil {
		value := uint(0)
		app.pHeight = &value
	}

	if app.pHeight == nil {
		return nil, errors.New("the height is mandatory in order to build a State instance")
	}

	if app.pointer == nil {
		return nil, errors.New("the pointer is mandatory in order to build a State instance")
	}

	if app.pCreatedOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a State instance")
	}

	data := []byte{}
	data = append(data, []byte(strconv.Itoa(int(*app.pHeight)))...)
	data = append(data, app.pointer.Hash()...)

	createdOn := *app.pCreatedOn
	data = append(data, []byte(strconv.Itoa(int(createdOn.UnixNano())))...)
	if app.previous != nil {
		data = append(data, app.previous...)
	}

	hash := sha512.New()
	_, err := hash.Write([]byte(data))
	if err != nil {
		return nil, err
	}

	if app.previous != nil {
		return createStateWithPrevious(hash.Sum(nil), *app.pHeight, app.pointer, *app.pCreatedOn, app.previous), nil
	}

	return createState(hash.Sum(nil), *app.pHeight, app.pointer, *app.pCreatedOn), nil
}
