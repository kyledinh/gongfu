package cfg

import (
	"testing"
)

func Test_LoadingConfig(t *testing.T) {
	var jc JsonCfg
	var err error
	var test_cfg = "test_cfg.json"

	if err = jc.ConfigFrom(test_cfg); err != nil {
		t.Error("FAILED to Test_LoadingConfig : %s", err)
	}
	if jc.Db_name != "test-name" {
		t.Error("FAILED DB_NAME : %s", DB_NAME)
	}

	Setup(test_cfg)

	if APPNAME != "test-gongfu" {
		t.Error("FAILED APPNAME : %s", APPNAME)
	}
	if DB_NAME != "test-name" {
		t.Error("FAILED DB_NAME : %s", DB_NAME)
	}
	if DB_USER != "test-user" {
		t.Error("FAILED DB_USER : %s", DB_USER)
	}
	if SERVER != "test-gongfu" {
		t.Error("FAILED Setup to configure root var SERVER : %s", SERVER)
	}
}
