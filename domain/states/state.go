package states

import (
	"time"

	"github.com/steve-care-software/tokendb/domain/pointers"
)

type state struct {
	hash      []byte
	height    uint
	pointer   pointers.Pointer
	createdOn time.Time
	previous  []byte
}

func createState(
	hash []byte,
	height uint,
	pointer pointers.Pointer,
	createdOn time.Time,
) State {
	return createStateInternally(hash, height, pointer, createdOn, nil)
}

func createStateWithPrevious(
	hash []byte,
	height uint,
	pointer pointers.Pointer,
	createdOn time.Time,
	previous []byte,
) State {
	return createStateInternally(hash, height, pointer, createdOn, previous)
}

func createStateInternally(
	hash []byte,
	height uint,
	pointer pointers.Pointer,
	createdOn time.Time,
	previous []byte,
) State {
	out := state{
		hash:      hash,
		height:    height,
		pointer:   pointer,
		createdOn: createdOn,
		previous:  previous,
	}

	return &out
}

// Hash returns the hash
func (obj *state) Hash() []byte {
	return obj.hash
}

// Height returns the height
func (obj *state) Height() uint {
	return obj.height
}

// Pointer returns the pointer
func (obj *state) Pointer() pointers.Pointer {
	return obj.pointer
}

// CreatedOn returns the creation time
func (obj *state) CreatedOn() time.Time {
	return obj.createdOn
}

// HasPrevious returns true if there is previous, false otherwise
func (obj *state) HasPrevious() bool {
	return obj.previous != nil
}

// Previous returns the previous, if any
func (obj *state) Previous() []byte {
	return obj.previous
}
