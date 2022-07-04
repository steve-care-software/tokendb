package transactions

import (
	"github.com/steve-care-software/tokendb/domain/pointers"
	"github.com/steve-care-software/tokendb/domain/states"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewTransactionBuilder creates a new transaction builder
func NewTransactionBuilder() TransactionBuilder {
	return createTransactionBuilder()
}

// Builder represents the transactions builder
type Builder interface {
	Create() Builder
	WithList(list []Transaction) Builder
	Now() (Transactions, error)
}

// Transactions represents transactions
type Transactions interface {
	List() []Transaction
}

// TransactionBuilder represents the transaction builder
type TransactionBuilder interface {
	Create() TransactionBuilder
	WithState(state states.State) TransactionBuilder
	WithData(data []byte) TransactionBuilder
	Now() (Transaction, error)
}

// Transaction represents a transaction
type Transaction interface {
	Hash() []byte
	State() states.State
	Data() []byte
}

// RepositoryBuilder represents a repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithState(state states.State) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents the transaction repository
type Repository interface {
	Retrieve(pointer pointers.Pointer) (Transaction, error)
	UnCommitted() (Transactions, error)
}

// Service represents a transaction service
type Service interface {
	Insert(trx Transaction) error
}
