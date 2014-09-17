package model

import (
	"fmt"
	"time"
)

type Chirp struct {
	Id          int       `json:"id"`
	Userid      int       `json:"userid"`
	Type        string    `json:"type"`
	Message     string    `json:"message"`
	Create_date time.Time `json:"create_date"`
	Active      bool      `json:"active"`
}

func (c *Chirp) StmtCountActive(state bool) string {
	if state == false {
		return fmt.Sprintf("SELECT COUNT(*) FROM chirps WHERE active = FALSE")
	}
	return fmt.Sprintf("SELECT COUNT(*) FROM chirps WHERE active = TRUE")
}

func (c *Chirp) StmtGetById() string {
	return fmt.Sprintf("SELECT * FROM chirps WHERE id = %v", c.Id)
}

func (c *Chirp) StmtInsert() string {
	cols := "userid, type, message"
	return fmt.Sprintf("INSERT INTO chirps (%s) VALUES (%v, '%s', '%s') RETURNING id",
		cols, c.Userid, c.Type, c.Message)
}

func (c *Chirp) StmtSelect(clause string) string {
	return "SELECT * FROM chirps " + clause
}

func (c *Chirp) StmtSelectByUserId() string {
	return fmt.Sprintf("SELECT * FROM chirps WHERE userid = %v AND active = TRUE", c.Userid)
}

func (c *Chirp) StmtSelectActive() string {
	return "SELECT * FROM chirps WHERE active = TRUE"
}

func (c *Chirp) StmtSetActive(state bool) string {
	if state == false {
		return fmt.Sprintf("UPDATE chirps SET active = FALSE WHERE id = %v", c.Id)
	}
	return fmt.Sprintf("UPDATE chirps SET active = TRUE WHERE id = %v", c.Id)
}

func (c *Chirp) StmtUpdate() string {
	cols := "userid, type, message, create_date"
	return fmt.Sprintf("UPDATE chirps SET (%s) = (%v, '%s', '%s', '%s') WHERE id = %v RETURNING id",
		cols, c.Userid, c.Type, c.Message, c.Create_date.Format(time.RFC3339), c.Id)
}
