package pointers

type pointer struct {
	hash  []byte
	index uint
	next  Pointer
}

func createPointer(
	hash []byte,
	index uint,
) Pointer {
	return createPointerInternally(hash, index, nil)
}

func createPointerWithNext(
	hash []byte,
	index uint,
	next Pointer,
) Pointer {
	return createPointerInternally(hash, index, next)
}

func createPointerInternally(
	hash []byte,
	index uint,
	next Pointer,
) Pointer {
	out := pointer{
		hash:  hash,
		index: index,
		next:  next,
	}

	return &out
}

// Hash returns the hash
func (obj *pointer) Hash() []byte {
	return obj.hash
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
