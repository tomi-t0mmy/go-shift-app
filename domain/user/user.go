package user

import "errors"

type User struct {
	email string
	Name  string
}

func NewUser(email, name string) (*User, error) {
	if email == "" {
		return nil, errors.New("メールアドレスを入力してください")
	}
	if name == "" {
		return nil, errors.New("名前を入力してください")
	}

	user := User{
		email: email,
		Name:  name,
	}

	return &user, nil
}
