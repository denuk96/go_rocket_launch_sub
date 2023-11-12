package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"go_rocket_launch_sub/internal/pkg/model"
	"go_rocket_launch_sub/internal/pkg/repository"
	"net/mail"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const tokenTTL = 96 * time.Hour

type AuthService struct {
	repo repository.Authorisation
}

func NewAuthService(repo repository.Authorisation) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp(user model.User) (string, error) {
	if err := validateMail(user.Email, s.repo); err != nil {
		return "", err
	}

	user.PasswordHash = hashPassword(user.PasswordHash)
	userId, err := s.repo.CreateUser(user)

	if err != nil {
		return "", err
	}

	userUUID, err := uuid.Parse(userId)
	if err != nil {
		return "", err
	}

	return generateToken(userUUID)
}

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId uuid.UUID `json:"user_id"`
}

func (s *AuthService) SignIn(email, password string) (string, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil || user.PasswordHash != hashPassword(password) {
		return "", errors.New("email or password invalid")
	}

	return generateToken(user.Id)
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	signingKey := os.Getenv("SECRET_KEY")

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId.String(), nil
}

func generateToken(userId uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		userId,
	})
	signingKey := os.Getenv("SECRET_KEY")

	return token.SignedString([]byte(signingKey))
}

func hashPassword(password string) string {
	salt := os.Getenv("PASSWORD_SALT")
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func validateMail(email string, authRepo repository.Authorisation) error {
	_, err := mail.ParseAddress(email)
	user, _ := authRepo.FindUserByEmail(email)

	if user != (model.User{}) {
		return errors.New("email already exists")
	}

	return err
}
