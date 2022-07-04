package dictionaries

import (
	"errors"
	"time"

	"github.com/steve-care-software/selector/domain/selectors"
)

type builder struct {
	tokens     Tokens
	selectors  []selectors.Selector
	pCreatedOn *time.Time
}

func createBuilder() Builder {
	out := builder{
		tokens:     nil,
		selectors:  nil,
		pCreatedOn: nil,
	}

	return &out
}

// Create intiializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithTokens add tokens to the builder
func (app *builder) WithTokens(tokens Tokens) Builder {
	app.tokens = tokens
	return app
}

// WithSelectors add selectors to the builder
func (app *builder) WithSelectors(selectors []selectors.Selector) Builder {
	app.selectors = selectors
	return app
}

// CreatedOn adds creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new Dictionary instance
func (app *builder) Now() (Dictionary, error) {
	if app.tokens == nil {
		return nil, errors.New("the tokens is mandatory in order to build a Dictionary instance")
	}

	if app.selectors != nil && len(app.selectors) <= 0 {
		app.selectors = nil
	}

	if app.selectors == nil {
		return nil, errors.New("there must be at least 1 Selector in order to build a Dictionary instance")
	}

	if app.pCreatedOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a Dictionary instance")
	}

	return createDictionary(app.tokens, app.selectors, *app.pCreatedOn), nil
}
