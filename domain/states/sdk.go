package states

import "time"

// Builder represents a state builder
type Builder interface {
	Create() Builder
	WithHeight(height uint) Builder
	WithPrevious(previous []byte) Builder
	WithPointers(pointers Pointers) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (State, error)
}

// State represents a state
type State interface {
	Height() uint
	Previous() []byte
	Pointers() Pointers
	CreatedOn() time.Time
}

// PointersBuilder represents a pointers builder
type PointersBuilder interface {
	Create() PointersBuilder
	WithIndex(index uint) PointersBuilder
	WithNext(next Pointers) PointersBuilder
	Now() (Pointers, error)
}

// Pointers represents data pointers
type Pointers interface {
	Index() uint
	HasNext() bool
	Next() Pointers
}
