package commons

type rnge struct {
	cursor Cursor
	amount uint
}

func createRange(
	cursor Cursor,
	amount uint,
) Range {
	out := rnge{
		cursor: cursor,
		amount: amount,
	}

	return &out
}

// Cursor returns the cursor
func (obj *rnge) Cursor() Cursor {
	return obj.cursor
}

// Amount returns the amount
func (obj *rnge) Amount() uint {
	return obj.amount
}
