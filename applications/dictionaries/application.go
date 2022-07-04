package dictionaries

import "github.com/steve-care-software/tokendb/domain/dictionaries"

type application struct {
	repository dictionaries.Repository
}

func createApplication(
	repository dictionaries.Repository,
) Application {
	out := application{
		repository: repository,
	}

	return &out
}

// Retrieve retrieves the dictionary
func (app *application) Retrieve() (dictionaries.Dictionary, error) {
	return app.repository.Retrieve()
}
