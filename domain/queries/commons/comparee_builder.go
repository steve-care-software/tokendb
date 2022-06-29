package commons

import "errors"

type compareeBuilder struct {
	element Element
	isAnd   bool
}

func createCompareeBuilder() CompareeBuilder {
	out := compareeBuilder{
		element: nil,
		isAnd:   false,
	}

	return &out
}

// Create initializes the builder
func (app *compareeBuilder) Create() CompareeBuilder {
	return createCompareeBuilder()
}

// WithElement adds an element to the builder
func (app *compareeBuilder) WithElement(element Element) CompareeBuilder {
	app.element = element
	return app
}

// IsAnd flags the builder as and
func (app *compareeBuilder) IsAnd() CompareeBuilder {
	app.isAnd = true
	return app
}

// Now builds a new Comparee instance
func (app *compareeBuilder) Now() (Comparee, error) {
	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build a Comparee instance")
	}

	return createComparee(app.element, app.isAnd), nil
}
