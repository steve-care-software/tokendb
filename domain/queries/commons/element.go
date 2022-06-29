package commons

type element struct {
	expectation Expectation
	comparison  Comparison
}

func createElementWithExpectation(
	expectation Expectation,
) Element {
	return createElementInternally(expectation, nil)
}

func createElementWithComparison(
	comparison Comparison,
) Element {
	return createElementInternally(nil, comparison)
}

func createElementInternally(
	expectation Expectation,
	comparison Comparison,
) Element {
	out := element{
		expectation: expectation,
		comparison:  comparison,
	}

	return &out
}

// IsExpectation returns true if there is an expectation, false otherwise
func (obj *element) IsExpectation() bool {
	return obj.expectation != nil
}

// Expectation returns the expectation, if any
func (obj *element) Expectation() Expectation {
	return obj.expectation
}

// IsComparison returns true if there is a comparison, false otherwise
func (obj *element) IsComparison() bool {
	return obj.comparison != nil
}

// Comparison returns the comparison, if any
func (obj *element) Comparison() Comparison {
	return obj.comparison
}
