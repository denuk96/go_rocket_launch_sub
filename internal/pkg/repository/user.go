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

func (psql *UserPsql) CreateUser(user model.User) (string, error) {
	result := psql.db.Create(&user)

	return user.Id.String(), result.Error
}

func (psql *UserPsql) FindUserById(id uuid.UUID) (model.User, error) {
	var user model.User
	result := psql.db.Take(&user, id)

	return user, result.Error
}

func (psql *UserPsql) FindUserByEmail(email string) (model.User, error) {
	var user model.User
	result := psql.db.First(&user, "email = ?", email)

	return user, result.Error
}
