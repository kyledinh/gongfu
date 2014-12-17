package api

import (
	"go/gongfu/cfg"
	"go/gongfu/model"
	"go/gongfu/util"
	"log"
	"testing"
)

func Test_User(t *testing.T) {

	cfg.Setup("../../../../cfg.json")
	util.SetupDB()

	var a model.User
	var b model.User
	var c model.User
	var token = util.GenHash(6)

	a.Accountid = 1
	a.Name = token
	a.Email = token + "@user.com"
	a.HashPassword("123")

	log.Println("a is : ", a)

	// Test StmtInsert
	row := util.DB.QueryRow(a.StmtInsert())
	err := row.Scan(&b.Id)
	log.Println("b is : ", b)
	if err != nil {
		t.Error("FAILED QueryRow : ", b)
	}

	// Test StmtGetById
	row = util.DB.QueryRow(b.StmtGetById())
	err = row.Scan(&b.Id, &b.Accountid, &b.Name, &b.Email, &b.Password, &b.Active)
	if err != nil {
		t.Error("FAILED Prepare Statement : ", err)
	}

	if a.Name != b.Name {
		t.Error("FAILED not same Name : ", a.Name)
	}
	if a.Password != b.Password {
		t.Error("FAILED not same Password : ", a.Password)
	}
	if a.Email != b.Email {
		t.Error("FAILED not same Email : ", a.Email)
	}

	// Test StmtUpdate
	b.Email = "updated" + token + "@email.com"
	b.Name = "updated"
	log.Println("b is now to be : ", b)

	_, err = util.DB.Exec(b.StmtUpdate())
	log.Println("Update query : ", b.StmtUpdate())
	log.Println("This, the error : ", err)

	c.Id = b.Id
	log.Println("Using this id: ", c.Id, b.Id)

	err = util.DB.QueryRow(c.StmtGetById()).Scan(&c.Id, &c.Accountid, &c.Name, &c.Email, &c.Password, &c.Active)
	if err != nil {
		t.Error("FAILED scan row : ", err)
	}

	log.Println("C now looks like : ", c)

	if c.Name != "updated" {
		t.Error("FAILED update Name : ", c.Name)
	}
	if c.Email != "updated"+token+"@email.com" {
		t.Error("FAILED update Email : ", c.Email)
	}

	// Test StmtSetActive
	_, err = util.DB.Exec(b.StmtSetActive(false))
	if err != nil {
		t.Error("FAILED StmtSetActive to false : ", err)
	}

	err = util.DB.QueryRow(c.StmtGetById()).Scan(&c.Id, &c.Accountid, &c.Name, &c.Email, &c.Password, &c.Active)
	if err != nil {
		t.Error("FAILED scan row : ", err)
	}

	if c.Active != false {
		t.Error("FAILED updated SetActive(false) : ", c.Active)
	}

	// Test StmtSelect
	rows, err := util.DB.Query(a.StmtSelect(""))
	if err != nil {
		t.Error("FAILED StmtSelect : ", err)
	}

	for rows.Next() {
		rows.Scan(&b.Id, &b.Accountid, &b.Name, &b.Email, &b.Password, &b.Active)
		log.Println("This is a row : ", b)
	}

	// Test StmtSelectActive
	rows, err = util.DB.Query(a.StmtSelectActive())
	if err != nil {
		t.Error("FAILED StmtSelectActive : ", err)
	}
	for rows.Next() {
		rows.Scan(&b.Id, &b.Accountid, &b.Name, &b.Email, &b.Password, &b.Active)
		log.Println("This is an active row : ", b)
	}

	// Test StmtSelectActive
	rows, err = util.DB.Query(a.StmtSelectByAccountId())
	if err != nil {
		t.Error("FAILED StmtSelectByAccountId : ", err)
	}
	for rows.Next() {
		rows.Scan(&b.Id, &b.Accountid, &b.Name, &b.Email, &b.Password, &b.Active)
		log.Println("This is has for this userid row : ", b)
	}

	// Test StmtCountActive
	var cnt int
	err = util.DB.QueryRow(b.StmtCountActive(true)).Scan(&cnt)
	if err != nil {
		t.Error("FAILED StmtCountActive", err)
	}
	if cnt < 1 {
		t.Error("FAILED found no active Users : ", cnt)
	}
}
