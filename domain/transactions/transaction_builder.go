package transactions

import (
	"crypto/sha512"
	"errors"

	"github.com/steve-care-software/tokendb/domain/states"
)

type transactionBuilder struct {
	state states.State
	data  []byte
}

func createTransactionBuilder() TransactionBuilder {
	out := transactionBuilder{
		state: nil,
		data:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *transactionBuilder) Create() TransactionBuilder {
	return createTransactionBuilder()
}

// WithState adds a state to the builder
func (app *transactionBuilder) WithState(state states.State) TransactionBuilder {
	app.state = state
	return app
}

// WithData adds data to the builder
func (app *transactionBuilder) WithData(data []byte) TransactionBuilder {
	app.data = data
	return app
}

// Now builds a new Transaction instance
func (app *transactionBuilder) Now() (Transaction, error) {
	if app.state == nil {
		return nil, errors.New("the state is mandatory in order to build a Transaction instance")
	}

	if app.data == nil {
		return nil, errors.New("the data is mandatory in order to build a Transaction instance")
	}

	data := []byte{}
	data = append(data, app.state.Hash()...)
	data = append(data, app.data...)

	hash := sha512.New()
	_, err := hash.Write([]byte(data))
	if err != nil {
		return nil, err
	}

	return createTransaction(hash.Sum(nil), app.state, app.data), nil
}
