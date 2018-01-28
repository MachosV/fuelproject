package models

import "fmt"

type User struct {
	email    string
	password string
	uuids    [][]byte
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) AppendUuid(uuid []byte) {
	u.uuids = append(u.uuids, uuid)
}

func (u User) Debug() (string, string, [][]byte) {
	fmt.Printf("%s %s %s\n", u.email, u.password, u.uuids)
	return u.email, u.password, u.uuids
}
