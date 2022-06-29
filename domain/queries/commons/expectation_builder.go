package commons

import (
	"errors"

	"github.com/steve-care-software/selector/domain/selectors"
)

type expectationBuilder struct {
	selector      selectors.Selector
	expectedValue []byte
	isNot         bool
}

func createExpectationBuilder() ExpectationBuilder {
	out := expectationBuilder{
		selector:      nil,
		expectedValue: nil,
		isNot:         false,
	}

	return &out
}

// Create initializes the builder
func (app *expectationBuilder) Create() ExpectationBuilder {
	return createExpectationBuilder()
}

// WithSelector adds a selector to the builder
func (app *expectationBuilder) WithSelector(selector selectors.Selector) ExpectationBuilder {
	app.selector = selector
	return app
}

// WithExpectedValue adds an expectedValue to the builder
func (app *expectationBuilder) WithExpectedValue(expectedValue []byte) ExpectationBuilder {
	app.expectedValue = expectedValue
	return app
}

// IsNot flags the bulder as not
func (app *expectationBuilder) IsNot() ExpectationBuilder {
	app.isNot = true
	return app
}

// Now builds a new Expectation instance
func (app *expectationBuilder) Now() (Expectation, error) {
	if app.selector == nil {
		return nil, errors.New("the selector is mandatory in order to build an Expectation instance")
	}

	if app.expectedValue != nil && len(app.expectedValue) <= 0 {
		app.expectedValue = nil
	}

	if app.expectedValue == nil {
		return nil, errors.New("the expectedValue is mandatory in order to build an Expectation instance")
	}

	return createExpectation(app.selector, app.expectedValue, app.isNot), nil
}
