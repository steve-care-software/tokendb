package dictionaries

import (
	"time"

	"github.com/steve-care-software/selector/domain/selectors"
	"github.com/steve-care-software/validator/domain/grammars"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewTokensBuilder creates a new tokens builder
func NewTokensBuilder() TokensBuilder {
	return createTokensBuilder()
}

// NewTokenBuilder creates a new token builder
func NewTokenBuilder() TokenBuilder {
	return createTokenBuilder()
}

// Adapter represents the dictionary adapter
type Adapter interface {
	ToDictionary(grammar grammars.Grammar) (Dictionary, error)
}

// Builder represents a dictionary builder
type Builder interface {
	Create() Builder
	WithTokens(tokens Tokens) Builder
	WithSelectors(selectors []selectors.Selector) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Dictionary, error)
}

// Dictionary represents the state dictionary
type Dictionary interface {
	Tokens() Tokens
	Selectors() []selectors.Selector
	CreatedOn() time.Time
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
