package dictionaries

import (
	"time"

	"github.com/steve-care-software/selector/domain/selectors"
)

// Builder represents a dictionary builder
type Builder interface {
	Create() Builder
	WithTokens(tokens Tokens) Builder
	WithSelectors(selectors Selectors) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Dictionary, error)
}

// Dictionary represents the state dictionary
type Dictionary interface {
	Tokens() Tokens
	Selectors() Selectors
	CreatedOn() time.Time
}

// SelectorsBuilder represents a selectors builder
type SelectorsBuilder interface {
	Create() SelectorsBuilder
	WithList(list []selectors.Selector) SelectorsBuilder
	Now() (Selectors, error)
}

// Selectors represents selectors
type Selectors interface {
	List() []selectors.Selector
	Matches(input selectors.Selector) ([]selectors.Selector, error)
}

// TokensBuilder represents a tokens builder
type TokensBuilder interface {
	Create() TokensBuilder
	WithList(list []Token) TokensBuilder
	Now() (Tokens, error)
}

// Tokens represents tokens
type Tokens interface {
	List() []Token
	Find(name string) (Token, error)
}

// TokenBuilder represents a token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithIndex(index uint) TokenBuilder
	WithName(name string) TokenBuilder
	Now() (Token, error)
}

// Token represents the token
type Token interface {
	Index() uint
	Name() string
}

// Repository represents the dictionary repository
type Repository interface {
	Retrieve() (Dictionary, error)
}

// Service represents the dictionary service
type Service interface {
	Insert(dictionary Dictionary) error
}
