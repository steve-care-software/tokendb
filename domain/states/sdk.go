package states

import (
	"time"

	"github.com/steve-care-software/tokendb/domain/pointers"
)

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
	FindByHeight(height uint) (State, error)
	FindByHash(hash []byte) (State, error)
	Previous() []byte
	Pointer() pointers.Pointer
	CreatedOn() time.Time
}

// Repository represents the state repository
type Repository interface {
	Retrieve() (State, error)
}

// Service represents the state service
type Service interface {
	Insert(state State) error
}
