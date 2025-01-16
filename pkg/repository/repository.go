package repository

// GenericRepository mendefinisikan operasi CRUD umum.
type GenericRepository[T any] interface {
	Create(entity T) error
	FindByID(id uint) (*T, error)
	Update(entity T) error
	Delete(id uint) error
}
