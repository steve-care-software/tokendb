package commons

type cursor struct {
	index      uint
	comparison Comparison
}

func createCursor(
	index uint,
) Cursor {
	return createCursorInternally(index, nil)
}

func createCursorWithComparison(
	index uint,
	comparison Comparison,
) Cursor {
	return createCursorInternally(index, comparison)
}

func createCursorInternally(
	index uint,
	comparison Comparison,
) Cursor {
	out := cursor{
		index:      index,
		comparison: comparison,
	}

	return &out
}

// Index returns the index
func (obj *cursor) Index() uint {
	return obj.index
}

// HasComparison returns true if there is a comparison, false otherwise
func (obj *cursor) HasComparison() bool {
	return obj.comparison != nil
}

// Comparison returns the comparison, if any
func (obj *cursor) Comparison() Comparison {
	return obj.comparison
}
