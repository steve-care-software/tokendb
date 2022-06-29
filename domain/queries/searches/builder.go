package searches

import (
	"errors"

	"github.com/steve-care-software/tokendb/domain/queries/commons"
	"github.com/steve-care-software/selector/domain/selectors"
	"github.com/steve-care-software/validator/domain/grammars"
)

type builder struct {
	selector selectors.Selector
	matcher  grammars.Grammar
	rnge     commons.Range
}

func createBuilder() Builder {
	out := builder{
		selector: nil,
		matcher:  nil,
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

// WithMatcher adds a matcher to the builder
func (app *builder) WithMatcher(matcher grammars.Grammar) Builder {
	app.matcher = matcher
	return app
}

// WithRange adds a range to the builder
func (app *builder) WithRange(rnge commons.Range) Builder {
	app.rnge = rnge
	return app
}

// Now builds a new Search instance
func (app *builder) Now() (Search, error) {
	if app.selector == nil {
		return nil, errors.New("the selector is mandatory in order to build a Search instance")
	}

	if app.matcher == nil {
		return nil, errors.New("the matcher is mandatory in order to build a Search instance")
	}

	if app.rnge == nil {
		return nil, errors.New("the range is mandatory in order to build a Search instance")
	}

	return createSearch(app.selector, app.matcher, app.rnge), nil
}
