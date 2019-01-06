package helper

import (
	"testing"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/model"
)

func TestToJSON(t *testing.T) {

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

	data, err := ToJSON(rec)

	if err != nil {
		t.Fatal(err)
	}

	if data == "" {
		t.Error("Expected a string")
	}
}

func TestToJSONIndent(t *testing.T) {

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

	data, err := ToJSONIndent(rec)

	if err != nil {
		t.Fatal(err)
	}

	if data == "" {
		t.Error("Expected a string")
	}

}

func TestIsEmptyTrue(t *testing.T) {

	m := make(map[int64][]int)

	empty, err := IsEmpty(m)

	if err != nil {
		t.Fatal(err)
	}

	if !empty {
		t.Errorf("Expected %t.\n Got %t.", empty, !empty)
	}
}

func TestIsEmptyFalse(t *testing.T) {

	m := make(map[int64][]int)

	m[0] = []int{1, 2, 3}

	empty, err := IsEmpty(m)

	if err != nil {
		t.Fatal(err)
	}

	if empty {
		t.Errorf("Expected %t.\n Got %t.", !empty, empty)
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
