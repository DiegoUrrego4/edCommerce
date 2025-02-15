package user

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/DiegoUrrego4/edCommerce/model"
)

type User struct {
	storage Storage
}

func New(s Storage) *User {
	return &User{s}
}

func (u *User) Create(newUser *model.User) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}

	newUser.ID = id
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%s %w", "bcrypt.GenerateFromPassword()", err)
	}

	newUser.Password = string(encryptedPassword)
	if newUser.Details == nil {
		newUser.Details = []byte("{}")
	}

	newUser.CreatedAt = time.Now().Unix()

	err = u.storage.Create(newUser)
	if err != nil {
		return fmt.Errorf("%s %w", "storage.Create()", err)
	}

	newUser.Password = ""

	return nil
}

func (u *User) GetByEmail(email string) (model.User, error) {
	user, err := u.storage.GetByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("%s %w", "storage.GetByEmail()", err)
	}

	return user, nil
}

func (u *User) GetAll() (model.Users, error) {
	users, err := u.storage.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%s %w", "storage.GetAll()", err)
	}

	return users, nil
}
