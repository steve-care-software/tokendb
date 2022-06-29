package transactions

import (
	"github.com/steve-care-software/tokendb/domain/pointers"
	"github.com/steve-care-software/tokendb/domain/states"
)

// Builder represents the transaction builder
type Builder interface {
	Create() Builder
	WithState(state states.State) Builder
	WithData(data []byte) Builder
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
}

// Service represents a transaction service
type Service interface {
	Insert(trx Transaction) error
}
