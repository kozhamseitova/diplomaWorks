package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/kozhamseitova/diplomaWorks/pkg/model"
	"github.com/kozhamseitova/diplomaWorks/pkg/repository"
)

const salt = "cbkjhfdcbdslkjdcbjhb"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
