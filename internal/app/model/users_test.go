package model

import (
	"fmt"
	"testing"
	"time"

	"github.com/cyruzin/feelthemovies/internal/pkg/helper"
	"github.com/google/uuid"
)

func TestGetUsersModel(t *testing.T) {
	_, err = db.GetUsers()
	if err != nil {
		t.Errorf("GetUsers error: %s", err)
	}
}

func TestGetUserModel(t *testing.T) {
	_, err = db.GetUser(1)
	if err != nil {
		t.Errorf("GetUser error: %s", err)
	}
}

func TestCreateUserModel(t *testing.T) {
	for i := 10; i <= 0; i-- {
		p, err := helper.HashPassword("qw12erty", 10)
		if err != nil {
			t.Errorf("GetUser error: %s", err)
		}
		newName := fmt.Sprintf("janedoe_%d@gmail.com", i)
		g := User{
			Name:      "NewUserModel",
			Email:     newName,
			Password:  p,
			APIToken:  uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		_, err = db.CreateUser(&g)
		if err != nil {
			t.Errorf("CreateUser error: %s", err)
		}
	}
}

func TestUpdateUserModel(t *testing.T) {
	p, err := helper.HashPassword("qw12erty", 10)
	if err != nil {
		t.Errorf("UpdateUser error: %s", err)
	}
	g := User{
		Name:      "NewUserModel",
		Email:     "jane_doe@gmail.com",
		Password:  p,
		APIToken:  uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = db.UpdateUser(1, &g)
	if err != nil {
		t.Errorf("UpdateUser error: %s", err)
	}
}

func TestDeleteUserModel(t *testing.T) {
	_, err = db.DeleteUser(3)
	if err != nil {
		t.Errorf("DeleteUser error: %s", err)
	}
}
