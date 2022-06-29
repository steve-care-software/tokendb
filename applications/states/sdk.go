package states

import (
	"github.com/steve-care-software/tokendb/applications/states/queries"
	"github.com/steve-care-software/tokendb/domain/states"
)

// Application represents the state application
type Application interface {
	Query(height uint) queries.Application
	RetrieveByHeight(height uint) (states.State, error)
	RetrieveByHash(hash []byte) (states.State, error)
	Height() uint
}
