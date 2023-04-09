package repository

import (
	"go_rocket_launch_sub/internal/pkg/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	usersTable = "users"
)

type UserPsql struct {
	db *gorm.DB
}

func NewUserPsql(db *gorm.DB) *UserPsql {
	return &UserPsql{db: db}
}

func (psql *UserPsql) CreateUser(user model.User) (uuid.UUID, error) {
	result := psql.db.Create(&user)

	return user.Id, result.Error
}
