package model

import (
	"testing"

	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
)

func TestCheckApiTokenModel(t *testing.T) {
	check, err := db.CheckAPIToken("ce3b81ee-0dc0-4133-8625-32007e64af7b")
	if err != nil {
		t.Errorf("CheckApiToken error: %s", err)
	}
	if !check {
		t.Errorf("CheckApiToken error: %s", err)
	}
}

func TestAuthenticateModel(t *testing.T) {
	dbPass, err := db.Authenticate("xorycx@gmail.com")
	if err != nil {
		t.Errorf("Authenticate error: %s", err)
	}
	checkPass := helper.CheckPasswordHash("-%O1r2y3c487-%", dbPass)
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
