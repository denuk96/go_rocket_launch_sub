package repository

// import "gorm.io/gorm"

type Authorisation interface {
}

type Subscription interface {
}

type Repository struct {
	Authorisation
	Subscription
}

func NewRepository() *Repository {
	return &Repository{}
}
