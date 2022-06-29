package application

import (
	"github.com/steve-care-software/tokendb/applications/states"
	"github.com/steve-care-software/tokendb/applications/transactions"
	"github.com/steve-care-software/tokendb/domain/dictionaries"
)

// Application represents the application
type Application interface {
	State() states.Application
	Dictionary() dictionaries.Dictionary
	Transactions() transactions.Application
}
