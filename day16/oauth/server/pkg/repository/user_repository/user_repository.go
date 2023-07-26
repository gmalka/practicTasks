package userrepository

import "errors"

type User struct {
	Username string
	Password string
	Secret string
	Clients map[string]struct{}
}

type UserRepotisory struct {
	users map[string]User
}

func NewUserRepository() UserRepotisory {
	return UserRepotisory{users: make(map[string]User)}
}

func (ur UserRepotisory) AddUser(username string, user User) {
	ur.users[username] = user
}

func (ur UserRepotisory) GetByName(username string) (User, error) {
	if v, ok := ur.users[username]; !ok {
		return User{}, errors.New("No such user exists")
	} else {
		return v, nil
	}
}

func (ur UserRepotisory) AddClientToUser(username, client string) error {
	if v, ok := ur.users[username]; !ok {
		return errors.New("No such user exists")
	} else {
		v.Clients[client] = struct{}{}
		ur.users[username] = v
		return nil
	}
}

func (ur UserRepotisory) GetUser(username, password string) (User, error) {
	if v, ok := ur.users[username]; !ok {
		return User{}, errors.New("No such user exists")
	} else {
		if v.Password != password {
			return User{}, errors.New("No such user exists")
		}
		return v, nil
	}
}

func (ur UserRepotisory) CreateUser(username, password, secret string) error {
	if _, ok := ur.users[username]; ok {
		return errors.New("User already exists")
	}

	ur.users[username] = User{Username: username, Password: password, Secret: secret, Clients: make(map[string]struct{})}
	return nil
}