package queries

import (
	"github.com/steve-care-software/tokendb/applications/states/queries/lists"
	"github.com/steve-care-software/tokendb/applications/states/queries/retrieves"
	"github.com/steve-care-software/tokendb/applications/states/queries/searches"
)

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithHeight(height uint) Builder
	Now() (Application, error)
}

// Application represents the queries application
type Application interface {
	Retrieve() retrieves.Application
	Search() searches.Application
	List() lists.Application
}
