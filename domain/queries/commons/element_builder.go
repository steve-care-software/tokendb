package commons

import "errors"

type elementBuilder struct {
	expectation Expectation
	comparison  Comparison
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		expectation: nil,
		comparison:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithExpectation adds an expectation to the builder
func (app *elementBuilder) WithExpectation(expectation Expectation) ElementBuilder {
	app.expectation = expectation
	return app
}

// WithComparison adds a comparison to the builder
func (app *elementBuilder) WithComparison(comparison Comparison) ElementBuilder {
	app.comparison = comparison
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.expectation != nil {
		return createElementWithExpectation(app.expectation), nil
	}

	if app.comparison != nil {
		return createElementWithComparison(app.comparison), nil
	}

	return nil, errors.New("the Element is invalid")
}
