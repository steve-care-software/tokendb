package dictionaries

import (
	"time"

	"github.com/steve-care-software/selector/domain/selectors"
)

type dictionary struct {
	tokens    Tokens
	selectors []selectors.Selector
	createdOn time.Time
}

func createDictionary(
	tokens Tokens,
	selectors []selectors.Selector,
	createdOn time.Time,
) Dictionary {
	out := dictionary{
		tokens:    tokens,
		selectors: selectors,
		createdOn: createdOn,
	}

	return &out
}

// Tokens returns the tokens
func (obj *dictionary) Tokens() Tokens {
	return obj.tokens
}

// Selectors returns the selectors
func (obj *dictionary) Selectors() []selectors.Selector {
	return obj.selectors
}

// CreatedOn returns the creation time
func (obj *dictionary) CreatedOn() time.Time {
	return obj.createdOn
}
