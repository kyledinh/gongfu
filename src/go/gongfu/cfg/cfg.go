package cfg

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Unique identifier for cookie in this app
var COOKIE_NAME string = "GONGFU-COOKIE"
var HASHKEY string = "gongfu-hash-key"
var BLOCKKEY string = "gongfu-block-key"

// defaults
var APPNAME string = "gongfu"
var DB_SERVER string = "localhost"
var DB_NAME string = "gongfu"
var DB_USER string = "gongfu"
var DB_PASSWORD string = "gongfu"
var HTTP_PORT string = ""
var SERVER string = "gongfu"

var conf JsonCfg

type JsonCfg struct {
	Appname     string
	Http_port   string
	Server      string
	Db_server   string
	Db_name     string
	Db_user     string
	Db_password string
}

// ConfigFrom is a method to read json properties
func (cfg *JsonCfg) ConfigFrom(path string) (err error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = json.Unmarshal(content, &cfg)
	if err != nil {
		log.Printf("bad json", err)
	}

	return
}

// ConfigSetup is the init function called from app.go, to read the cfg.json file
// If error reading the cfg.json; defaults are used
func Setup(path string) {
	err := conf.ConfigFrom(path)
	if err != nil {
		log.Printf("######### CRITICAL: Failed to load cfg.json %v", err)
		log.Printf("######### USING DEFAULT CFG SETTINGS!!!")
	} else {
		APPNAME = conf.Appname
		DB_SERVER = conf.Db_server
		DB_NAME = conf.Db_name
		DB_USER = conf.Db_user
		DB_PASSWORD = conf.Db_password
		HTTP_PORT = conf.Http_port
		SERVER = conf.Server
		log.Printf("... loaded cfg.json ...")
		log.Printf("DB_NAME: %v", DB_NAME)
		log.Printf("DB_USER: %v", DB_USER)
		log.Printf("DB_PASSWORD: %v", DB_PASSWORD)
		log.Printf("COOKIE_NAME: %v", COOKIE_NAME)
	}
}
