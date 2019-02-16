package helper

import (
	"testing"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/model"
)

func TestMarshalBinary(t *testing.T) {
	rec := model.Recommendation{
		ID:        1,
		UserID:    1,
		Title:     "Test title",
		Type:      0,
		Poster:    "huApZqTkkLç",
		Backdrop:  "ppkLiWUq",
		Status:    0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := MarshalBinary(rec)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnmarshalBinary(t *testing.T) {
	rec := []byte(`{"name": "John Doe"}`)
	var name struct {
		Name string
	}
	err := UnmarshalBinary(rec, &name)
	if err != nil {
		t.Fatal(err)
	}
}

func TestHashPassword(t *testing.T) {
	hashPass, err := HashPassword("teste", 10)
	if err != nil {
		t.Fatal(err)
	}
	if hashPass == "" {
		t.Error("Expected a string")
	}
}

func TestCheckPassword(t *testing.T) {
	pass := "teste"
	hashPass, err := HashPassword(pass, 10)
	if err != nil {
		t.Fatal(err)
	}
	match := CheckPasswordHash("teste", hashPass)
	if match == false {
		t.Errorf("Expected %t.\n Got %t", !match, match)
	}
}
