package dictionaries

import "github.com/steve-care-software/tokendb/domain/dictionaries"

// NewApplication creates a new application instance
func NewApplication(
	repository dictionaries.Repository,
) Application {
	return createApplication(repository)
}

// Application represents the dictionary application
type Application interface {
	Retrieve() (dictionaries.Dictionary, error)
}
