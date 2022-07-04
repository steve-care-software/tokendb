package states

import (
	"time"

	"github.com/steve-care-software/tokendb/domain/pointers"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a state builder
type Builder interface {
	Create() Builder
	WithHeight(height uint) Builder
	WithPrevious(previous []byte) Builder
	WithPointer(pointer pointers.Pointer) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (State, error)
}

// State represents a state
type State interface {
	Hash() []byte
	Height() uint
	Pointer() pointers.Pointer
	CreatedOn() time.Time
	HasPrevious() bool
	Previous() []byte
}

// Repository represents the state repository
type Repository interface {
	RetrieveHead() (State, error)
	RetrieveByHeight(height uint) (State, error)
	RetrieveByHash(hash []byte) (State, error)
}

// Service represents the state service
type Service interface {
	Insert(state State) error
}
