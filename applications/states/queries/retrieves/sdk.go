package retrieves

import "github.com/steve-care-software/tokendb/domain/queries/retrieves"

// Application represents a retrieve application
type Application interface {
	Compile(script string) (retrieves.Retrieve, error)
	Execute(qery retrieves.Retrieve) ([]byte, error)
}
