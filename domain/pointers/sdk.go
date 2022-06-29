package pointers

// Builder represents a pointers builder
type Builder interface {
	Create() Builder
	WithIndex(index uint) Builder
	WithNext(next Pointer) Builder
	Now() (Pointer, error)
}

// Pointer represents data pointer
type Pointer interface {
	Index() uint
	HasNext() bool
	Next() Pointer
}
