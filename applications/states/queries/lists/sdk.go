package lists

import "github.com/steve-care-software/tokendb/domain/queries/lists"

// Application represents a list application
type Application interface {
	Compile(script string) (lists.List, error)
	Execute(qery lists.List) ([]byte, uint, error)
}
