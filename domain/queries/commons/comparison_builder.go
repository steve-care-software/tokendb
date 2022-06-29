package commons

import "errors"

type comparisonBuilder struct {
	comparee    Comparee
	expectation Expectation
	isNot       bool
}

func createComparisonBuilder() ComparisonBuilder {
	out := comparisonBuilder{
		comparee:    nil,
		expectation: nil,
		isNot:       false,
	}

	return &out
}

// Create initializes the builder
func (app *comparisonBuilder) Create() ComparisonBuilder {
	return createComparisonBuilder()
}

// WithComparee adds a comparee the builder
func (app *comparisonBuilder) WithComparee(comparee Comparee) ComparisonBuilder {
	app.comparee = comparee
	return app
}

// WithExpectation expectation to the builder
func (app *comparisonBuilder) WithExpectation(expectation Expectation) ComparisonBuilder {
	app.expectation = expectation
	return app
}

// IsNot flags the builder as not
func (app *comparisonBuilder) IsNot() ComparisonBuilder {
	app.isNot = true
	return app
}

// Now builds a new Comparison instance
func (app *comparisonBuilder) Now() (Comparison, error) {
	if app.comparee == nil {
		return nil, errors.New("the comparee is mandatory in order to build a Comparison instance")
	}

	if app.expectation == nil {
		return nil, errors.New("the expectation is mandatory in order to build a Comparison instance")
	}

	return createComparison(app.comparee, app.expectation, app.isNot), nil
}
