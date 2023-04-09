package repository

import "gorm.io/gorm"

type Authorisation interface {
}

type Subscription interface {
}

type Repository struct {
	Authorisation
	Subscription
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{}
}
