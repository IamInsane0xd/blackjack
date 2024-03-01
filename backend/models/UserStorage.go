package models

import "errors"

var (
	UserAlreadyExistsErr = errors.New("user already exists")
	UserNotFoundErr      = errors.New("user not found")
)

type UserStorage struct {
	Users map[string]string "json:\"users\""
}

func NewUserStorage() *UserStorage {
	return &UserStorage{make(map[string]string)}
}

func (u *UserStorage) Add(uName string, pass string) error {
	if _, ok := u.Users[uName]; ok {
		return UserAlreadyExistsErr
	}

	u.Users[uName] = pass
	return nil
}

func (u *UserStorage) Find(uName string) (string, error) {
	val, ok := u.Users[uName]

	if !ok {
		return "", UserNotFoundErr
	}

	return val, nil
}
