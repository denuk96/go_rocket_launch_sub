package service

import (
	"crypto/sha1"
	"fmt"
	"go_rocket_launch_sub/internal/pkg/model"
	"go_rocket_launch_sub/internal/pkg/repository"
	"os"

	"github.com/google/uuid"
)

type AuthService struct {
	repo repository.Authorisation
}

func NewAuthService(repo repository.Authorisation) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (uuid.UUID, error) {
	user.PasswordHash = hashPassword(user.PasswordHash)
	return s.repo.CreateUser(user)
}

func hashPassword(password string) string {
	salt := os.Getenv("PASSWORD_SALT")
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
