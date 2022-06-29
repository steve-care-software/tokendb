package commons

import "github.com/steve-care-software/selector/domain/selectors"

type expectation struct {
	selector      selectors.Selector
	expectedValue []byte
	isNot         bool
}

func createExpectation(
	selector selectors.Selector,
	expectedValue []byte,
	isNot bool,
) Expectation {
	out := expectation{
		selector:      selector,
		expectedValue: expectedValue,
		isNot:         isNot,
	}

	return &out
}

// Selector returns the selector
func (obj *expectation) Selector() selectors.Selector {
	return obj.selector
}

// ExpectedValue returns the expectedValue
func (obj *expectation) ExpectedValue() []byte {
	return obj.expectedValue
}

// IsNot returns true if not, false otherwise
func (obj *expectation) IsNot() bool {
	return obj.isNot
}
