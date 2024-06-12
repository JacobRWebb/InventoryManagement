package dbstore

import (
	"errors"
	"fmt"
	"time"

	"github.com/JacobRWebb/InventoryManagement/pkg/models"
	"github.com/JacobRWebb/InventoryManagement/pkg/utils/hash"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserStore struct {
	db           *gorm.DB
	passwordhash hash.PasswordHash
}

type NewUserStoreParams struct {
	DB           *gorm.DB
	PasswordHash hash.PasswordHash
}

var (
	ErrEmailTaken = errors.New("email is taken")
)

func NewUserStore(params NewUserStoreParams) *UserStore {
	return &UserStore{
		db:           params.DB,
		passwordhash: params.PasswordHash,
	}
}

func (s *UserStore) CreateUser(email string, password string) error {
	encodedPassword, err := s.passwordhash.GenerateFromPassword(password)

	if err != nil {
		return err
	}

	existingUser, err := s.GetUserByEmail(email)

	if err != nil {
		return err
	}

	fmt.Printf("User: %v", existingUser)

	if existingUser != nil {
		return ErrEmailTaken
	}

	return s.db.Create(&models.User{
		Id:           uuid.New(),
		Email:        email,
		PasswordHash: encodedPassword,
		CreatedAt:    time.Now(),
		ModifiedAt:   time.Now(),
	}).Error
}

func (s *UserStore) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := s.db.Where(&models.User{Email: email}).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
