package dictionaries

import "github.com/steve-care-software/tokendb/domain/dictionaries"

// Application represents the dictionary application
type Application interface {
	Retrieve() (dictionaries.Dictionary, error)
}
