package model

import (
	"go/gongfu/util"
	"fmt"
)

type User struct {
	Id        int    `json:"id"`
	Accountid int    `json:"accountid"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Active    bool   `json:"active"`
}

func (u *User) HashPassword(str string) string {
	hash, _ := util.HashPassword(str)
	u.Password = string(hash)
	return u.Password
}

func (u *User) StmtCountActive(state bool) string {
	if state == false {
		return fmt.Sprintf("SELECT COUNT(*) FROM users WHERE active = FALSE")
	}
	return fmt.Sprintf("SELECT COUNT(*) FROM users WHERE active = TRUE")
}

func (u *User) StmtGetById() string {
	return fmt.Sprintf("SELECT * FROM users WHERE id = %v", u.Id)
}

func (u *User) StmtGetByEmail() string {
	return fmt.Sprintf("SELECT * FROM users WHERE email = '%s'", u.Email)
}

func (u *User) StmtInsert() string {
	return fmt.Sprintf("insert into users (accountid, name, email, password) values (%v, '%s', '%s', '%s') RETURNING id",
		u.Accountid, u.Name, u.Email, u.Password)
}

func (u *User) StmtSelect(clause string) string {
	return "SELECT * FROM users " + clause
}

func (u *User) StmtSelectByAccountId() string {
	return fmt.Sprintf("SELECT * FROM users WHERE accountid = %v AND active = TRUE", u.Accountid)
}

func (u *User) StmtSelectActive() string {
	return "SELECT * FROM users WHERE active = TRUE"
}

func (u *User) StmtSetActive(state bool) string {
	if state == false {
		return fmt.Sprintf("UPDATE users SET active = FALSE WHERE id = %v", u.Id)
	}
	return fmt.Sprintf("UPDATE users SET active = TRUE WHERE id = %v", u.Id)
}

func (u *User) StmtUpdate() string {
	return fmt.Sprintf("UPDATE users SET (accountid, name, email, password) = (%v, '%s', '%s', '%s') WHERE id = %v RETURNING id",
		u.Accountid, u.Name, u.Email, u.Password, u.Id)
}
