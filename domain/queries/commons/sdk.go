package commons

import (
	"github.com/steve-care-software/selector/domain/selectors"
)

// NewRangeBuilder creates a new range builder
func NewRangeBuilder() RangeBuilder {
	return createRangeBuilder()
}

// NewCursorBuilder creates a new cursor builder
func NewCursorBuilder() CursorBuilder {
	return createCursorBuilder()
}

// NewComparisonBuilder creates a new comparison builder
func NewComparisonBuilder() ComparisonBuilder {
	return createComparisonBuilder()
}

// NewCompareeBuilder creates a new comparee builder
func NewCompareeBuilder() CompareeBuilder {
	return createCompareeBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// NewExpectationBuilder creates a new expectation builder
func NewExpectationBuilder() ExpectationBuilder {
	return createExpectationBuilder()
}

// RangeBuilder represents a range builder
type RangeBuilder interface {
	Create() RangeBuilder
	WithCursor(cursor Cursor) RangeBuilder
	WithAmount(amount uint) RangeBuilder
	Now() (Range, error)
}

// Range represents the range
type Range interface {
	Cursor() Cursor
	Amount() uint
}

// CursorBuilder represents a cursor builder
type CursorBuilder interface {
	Create() CursorBuilder
	WithIndex(index uint) CursorBuilder
	WithComparison(comparison Comparison) CursorBuilder
	Now() (Cursor, error)
}

// Cursor represents the cursor
type Cursor interface {
	Index() uint
	HasComparison() bool
	Comparison() Comparison
}

// ComparisonBuilder represents a comparison builder
type ComparisonBuilder interface {
	Create() ComparisonBuilder
	WithComparee(comparee Comparee) ComparisonBuilder
	WithExpectation(expectation Expectation) ComparisonBuilder
	IsNot() ComparisonBuilder
	Now() (Comparison, error)
}

// Comparison represents a comparison
type Comparison interface {
	Comparee() Comparee
	Expectation() Expectation
	IsNot() bool
}

// CompareeBuilder represents a comparee builder
type CompareeBuilder interface {
	Create() CompareeBuilder
	WithElement(element Element) CompareeBuilder
	IsAnd() CompareeBuilder
	Now() (Comparee, error)
}

// Comparee represents a comparee
type Comparee interface {
	Element() Element
	IsAnd() bool
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithExpectation(expectation Expectation) ElementBuilder
	WithComparison(comparison Comparison) ElementBuilder
	Now() (Element, error)
}

// Element represents a comparison element
type Element interface {
	IsExpectation() bool
	Expectation() Expectation
	IsComparison() bool
	Comparison() Comparison
}

// ExpectationBuilder represents the expectation builder
type ExpectationBuilder interface {
	Create() ExpectationBuilder
	WithSelector(selector selectors.Selector) ExpectationBuilder
	WithExpectedValue(expectedValue []byte) ExpectationBuilder
	IsNot() ExpectationBuilder
	Now() (Expectation, error)
}

// Expectation represents an expectation
type Expectation interface {
	Selector() selectors.Selector
	ExpectedValue() []byte
	IsNot() bool
}
