package models

import (
	"github.com/pkg/errors"
	"strconv"
)

type (
	User struct {
		ID string `json:"id"`
		DisplayName string `json:"displayName"`
		Email string `json:"email"`
		Password string `json:"password"`
	}

	Users []User
)

func (c *Users) Get () error {
	*c = make(Users, 0)

	for _, r := range userDatabase {
		*c = append(*c, r)
	}

	return nil
}

func (c *Users) GetByID (ids []string) (error) {
	*c = make(Users, 0)

	for _, id := range ids {
		r, ok := userDatabase[id]
		if !ok {
			return errors.New(`Record not found.`)
		}
		*c = append(*c, r)
	}

	return nil
}

func (m *User) Get (id string) error {
	var r, ok = userDatabase[id]
	if !ok {
		return errors.New(`Record not found.`)
	}

	*m = r
	return nil
}

func (m *User) Save () error {
	if m.ID == `` {
		m.ID = strconv.Itoa(userIndex)
		userIndex++
	}
	userDatabase[m.ID] = *m
	return nil
}

func (m *User) Delete () error {
	_, ok := userDatabase[m.ID]
	if !ok {
		return errors.New(`Record not found.`)
	}

	delete(userDatabase, m.ID)
	return nil
}