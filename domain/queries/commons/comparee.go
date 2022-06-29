package commons

type comparee struct {
	element Element
	isAnd   bool
}

func createComparee(
	element Element,
	isAnd bool,
) Comparee {
	out := comparee{
		element: element,
		isAnd:   isAnd,
	}

	return &out
}

// Element returns the element
func (obj *comparee) Element() Element {
	return obj.element
}

// IsAnd returns true if and, false is or
func (obj *comparee) IsAnd() bool {
	return obj.isAnd
}
