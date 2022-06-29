package commons

type comparison struct {
	comparee    Comparee
	expectation Expectation
	isNot       bool
}

func createComparison(
	comparee Comparee,
	expectation Expectation,
	isNot bool,
) Comparison {
	out := comparison{
		comparee:    comparee,
		expectation: expectation,
		isNot:       isNot,
	}

	return &out
}

// Comparee returns the comparee
func (obj *comparison) Comparee() Comparee {
	return obj.comparee
}

// Expectation returns the expectation
func (obj *comparison) Expectation() Expectation {
	return obj.expectation
}

// IsNot returns true if not, false otherwise
func (obj *comparison) IsNot() bool {
	return obj.isNot
}
