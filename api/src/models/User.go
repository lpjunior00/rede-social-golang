package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID           uint64    `json:"id,omitempty"` //omit empty serve pra tirar do json se nao tiver valor
	Name         string    `json:"name,omitempty"`
	Nickname     string    `json:"nickname,omitempty"`
	Email        string    `json:"email,omitempty"`
	Password     string    `json:"password,omitempty"`
	CreationDate time.Time `json:"creationDate,omitempty"`
}

func (user *User) Prepare(etapa string) error {
	if erro := user.validate(etapa); erro != nil {
		return erro
	}

	if erro := user.removeSpaces(etapa); erro != nil {
		return erro
	}

	return nil
}

func (user *User) validate(etapa string) error {
	if user.Name == "" {
		return errors.New("Name is a required field and cannot be empty")
	}

	if user.Nickname == "" {
		return errors.New("Nickname is a required field and cannot be empty")
	}

	if user.Email == "" {
		return errors.New("Email is a required field and cannot be empty")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("Email is invalid!")
	}

	if etapa == "create" && user.Password == "" {
		return errors.New("Password is a required field and cannot be empty")
	}

	return nil
}

func (user *User) removeSpaces(etapa string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nickname = strings.TrimSpace(user.Nickname)
	user.Email = strings.TrimSpace(user.Email)

	if etapa == "create" {
		passwordHash, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}

		user.Password = string(passwordHash)
	}

	return nil
}
