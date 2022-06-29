package searches

import "github.com/steve-care-software/tokendb/domain/queries/searches"

// Application represents a search application
type Application interface {
	Compile(script string) (searches.Search, error)
	Execute(qery searches.Search) ([]byte, error)
}
