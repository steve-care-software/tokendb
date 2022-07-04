package dictionaries

type token struct {
	index uint
	name  string
}

func createToken(
	index uint,
	name string,
) Token {
	out := token{
		index: index,
		name:  name,
	}

	return &out
}

// Index returns the index
func (obj *token) Index() uint {
	return obj.index
}

// Name returns the name
func (obj *token) Name() string {
	return obj.name
}
