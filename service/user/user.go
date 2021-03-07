package user

import "github.com/Sanketkhote/microService/service/user/model"

type User interface {
	SaveUser(data model.UserModel) (bool, error)
}

type user struct {
}

var users = make(map[string]model.UserModel)

func NewUser() User {

	return &user{}
}

func (u *user) SaveUser(data model.UserModel) (bool, error) {
	user, ok := users[data.Name]
	if ok {
		return true, nil
	}
	users[data.Name] = user
	return false, nil
}
