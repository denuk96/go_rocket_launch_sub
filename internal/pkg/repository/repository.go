package repository

import (
	"github.com/google/uuid"
	"go_rocket_launch_sub/internal/pkg/model"
	"gorm.io/gorm"
)

type Authorisation interface {
	CreateUser(user model.User) (uuid.UUID, error)
}

type Subscription interface {
}

type Repository struct {
	Authorisation
	Subscription
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorisation: NewUserPsql(db),
	}
}
