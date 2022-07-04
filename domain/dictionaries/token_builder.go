package dictionaries

import "errors"

type tokenBuilder struct {
	pIndex *uint
	name   string
}

func createTokenBuilder() TokenBuilder {
	out := tokenBuilder{
		pIndex: nil,
		name:   "",
	}

	return &out
}

// Create initializes the builder
func (app *tokenBuilder) Create() TokenBuilder {
	return createTokenBuilder()
}

// WithIndex adds an index to the builder
func (app *tokenBuilder) WithIndex(index uint) TokenBuilder {
	app.pIndex = &index
	return app
}

// WithName adds a name to the builder
func (app *tokenBuilder) WithName(name string) TokenBuilder {
	app.name = name
	return app
}

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Token instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Token instance")
	}

	return createToken(*app.pIndex, app.name), nil
}
