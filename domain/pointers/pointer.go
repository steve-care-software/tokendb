package pointers

type pointer struct {
	index uint
	next  Pointer
}

func createPointer(
	index uint,
) Pointer {
	return createPointerInternally(index, nil)
}

func createPointerWithNext(
	index uint,
	next Pointer,
) Pointer {
	return createPointerInternally(index, next)
}

func createPointerInternally(
	index uint,
	next Pointer,
) Pointer {
	out := pointer{
		index: index,
		next:  next,
	}

	return &out
}

// Index returns the index
func (obj *pointer) Index() uint {
	return obj.index
}

// HasNext returns true if there is a next pointer, false otherwise
func (obj *pointer) HasNext() bool {
	return obj.next != nil
}

// Next returns the next pointer, if any
func (obj *pointer) Next() Pointer {
	return obj.next
}
