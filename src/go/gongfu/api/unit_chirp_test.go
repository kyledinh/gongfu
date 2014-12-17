package api

import (
	"go/gongfu/cfg"
	"go/gongfu/model"
	"go/gongfu/util"
	"log"
	"testing"
)

func Test_Chirp(t *testing.T) {

	cfg.Setup("../../../../cfg.json")
	util.SetupDB()

	var a model.Chirp
	var b model.Chirp
	var c model.Chirp
	var token = util.GenHash(6)
	//var id int = 110
	a.Userid = 1
	a.Type = token
	a.Message = token + "-message"

	log.Println(a)
	// Test StmtInsert
	row := util.DB.QueryRow(a.StmtInsert())
	err := row.Scan(&b.Id)
	log.Println(b)
	if err != nil {
		t.Error("FAILED QueryRow : ", b)
	}

	// Test StmtGetById
	row = util.DB.QueryRow(b.StmtGetById())
	err = row.Scan(&b.Id, &b.Userid, &b.Type, &b.Message, &b.Create_date, &b.Active)
	if err != nil {
		t.Error("FAILED Prepare Statement : ", err)
	}

	if a.Userid != b.Userid {
		t.Error("FAILED not same User : ", a.Userid)
	}
	if a.Type != b.Type {
		t.Error("FAILED not same Type : ", a.Type)
	}
	if a.Message != b.Message {
		t.Error("FAILED not same Message : ", a.Message)
	}

	// Test StmtUpdate
	b.Message = "updated-msg"
	b.Type = "updated"
	log.Println("b is now updated : ", b)

	_, err = util.DB.Exec(b.StmtUpdate())
	log.Println("Update query : ", b.StmtUpdate())
	log.Println("This, the error : ", err)

	c.Id = b.Id
	log.Println("Using this id: ", c.Id, b.Id)

	err = util.DB.QueryRow(c.StmtGetById()).Scan(&c.Id, &c.Userid, &c.Type, &c.Message, &c.Create_date, &c.Active)
	if err != nil {
		t.Error("FAILED scan row : ", err)
	}

	log.Println("C now looks like : ", c)

	if c.Type != "updated" {
		t.Error("FAILED update Name : ", c.Type)
	}
	if c.Message != "updated-msg" {
		t.Error("FAILED update Message : ", c.Message)
	}

	// Test StmtSetActive
	_, err = util.DB.Exec(c.StmtSetActive(false))
	if err != nil {
		t.Error("FAILED StmtSetActive to false : ", err)
	}

	err = util.DB.QueryRow(c.StmtGetById()).Scan(&c.Id, &c.Userid, &c.Type, &c.Message, &c.Create_date, &c.Active)
	if err != nil {
		t.Error("FAILED scan row : ", err)
	}

	if c.Active != false {
		t.Error("FAILED updated SetActive(false) : ", c.Active)
	}

	// Test StmtSetActive
	_, err = util.DB.Exec(c.StmtSetActive(true))
	if err != nil {
		t.Error("FAILED StmtSetActive to false : ", err)
	}

	// Test StmtSelect
	rows, err := util.DB.Query(b.StmtSelect(""))
	if err != nil {
		t.Error("FAILED StmtSelect : ", err)
	}

	for rows.Next() {
		rows.Scan(&b.Id, &b.Userid, &b.Type, &b.Message, &b.Create_date, &b.Active)
		log.Println("This is a unit row : ", b)
	}

	// Test StmtSelectActive
	rows, err = util.DB.Query(a.StmtSelectActive())
	if err != nil {
		t.Error("FAILED StmtSelectActive : ", err)
	}
	for rows.Next() {
		rows.Scan(&b.Id, &b.Userid, &b.Type, &b.Message, &b.Create_date, &b.Active)
		log.Println("This is an active row : ", b)
	}

	// Test StmtSelectActive
	rows, err = util.DB.Query(a.StmtSelectByUserId())
	if err != nil {
		t.Error("FAILED StmtSelectByUserId : ", err)
	}
	for rows.Next() {
		rows.Scan(&b.Id, &b.Userid, &b.Type, &b.Message, &b.Create_date, &b.Active)
		log.Println("This is has for unit row : ", b)
	}

	// Test StmtCountActive
	var cnt int
	err = util.DB.QueryRow(c.StmtCountActive(true)).Scan(&cnt)
	if err != nil {
		t.Error("FAILED StmtCountActive", err)
	}
	if cnt < 1 {
		t.Error("FAILED found no active Units : ", cnt)
	}

}
