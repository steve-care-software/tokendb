package transactions

type transactions struct {
	list []Transaction
}

func createTransactions(
	list []Transaction,
) Transactions {
	out := transactions{
		list: list,
	}

	return &out
}

// List returns the transactions list
func (obj *transactions) List() []Transaction {
	return obj.list
}
