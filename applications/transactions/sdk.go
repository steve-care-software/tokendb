package transactions

// Application represents the transaction application
type Application interface {
	Begin(height uint) ([]byte, error)
	Transact(context []byte, data []byte) error
	Commit(context []byte, message []byte) error
	RollBack(context []byte) error
	Push() error
}
