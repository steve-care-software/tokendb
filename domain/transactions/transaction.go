package transactions

import "github.com/steve-care-software/tokendb/domain/states"

type transaction struct {
	hash  []byte
	state states.State
	data  []byte
}

func createTransaction(
	hash []byte,
	state states.State,
	data []byte,
) Transaction {
	out := transaction{
		hash:  hash,
		state: state,
		data:  data,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() []byte {
	return obj.hash
}

// State returns the state
func (obj *transaction) State() states.State {
	return obj.state
}

// Data returns the data
func (obj *transaction) Data() []byte {
	return obj.data
}
