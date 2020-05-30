package user

import (
	"github.com/rithikjain/MongoNotes/pkg"
	"github.com/rithikjain/MongoNotes/pkg/entities"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type Service interface {
	Register(user *entities.User) (*entities.User, error)

	Login(email, password string) (*entities.User, error)

	GetUserByID(id string) (*entities.User, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func Validate(user *entities.User) (bool, error) {
	if !strings.Contains(user.Email, "@") {
		return false, pkg.ErrEmail
	}

	if len(user.Password) < 6 || len(user.Password) > 60 {
		return false, pkg.ErrPassword
	}
	return true, nil
}

func (s *service) Register(user *entities.User) (*entities.User, error) {
	// Validation
	validate, err := Validate(user)
	if !validate {
		return nil, err
	}
	exists, err := s.repo.DoesEmailExist(user.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, pkg.ErrExists
	}
	pass, err := HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = pass
	return s.repo.Register(user)
}

func (s *service) Login(email, password string) (*entities.User, error) {
	user := &entities.User{}
	us, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if CheckPasswordHash(password, user.Password) {
		return us, nil
	}
	return nil, pkg.ErrNotFound
}

func (s *service) GetUserByID(id string) (*entities.User, error) {
	return s.repo.FindByID(id)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
