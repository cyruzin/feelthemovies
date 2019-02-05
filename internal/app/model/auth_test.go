package model

import (
	"testing"

	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
)

func TestCheckApiTokenModel(t *testing.T) {
	check, err := db.CheckAPIToken("bb9b6ed1-8688-44f4-9f35-f75e62ef83f1")
	if err != nil {
		t.Errorf("CheckApiToken error: %s", err)
	}
	if !check {
		t.Errorf("CheckApiToken error: %s", err)
	}
	data, err := helper.ToJSON(check)
	if err != nil {
		t.Errorf("CheckApiToken - ToJSON - error: %s", err)
	}
	t.Log(data)
}

func TestAuthenticateModel(t *testing.T) {
	dbPass, err := db.Authenticate("xorycx@gmail.com")
	if err != nil {
		t.Errorf("Authenticate error: %s", err)
	}
	checkPass := helper.CheckPasswordHash("qw12erty", dbPass)
	if !checkPass {
		t.Errorf("Authenticate error: %s", err)
	}
}

func TestGetAuthInfoModel(t *testing.T) {
	_, err := db.GetAuthInfo("xorycx@gmail.com")
	if err != nil {
		t.Errorf("Authenticate error: %s", err)
	}
}
