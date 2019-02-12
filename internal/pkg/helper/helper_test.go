package helper

import (
	"testing"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/model"
)

var db, err = model.Connect()

func TestAttach(t *testing.T) {
	m := make(map[int64][]int)
	m[1] = []int{1, 2}
	err := Attach(m, "genre_recommendation", db.DB)
	if err != nil {
		t.Errorf("Attach error: %s", err)
	}
}

func TestDetach(t *testing.T) {
	m := make(map[int64][]int)
	m[1] = []int{2, 3}
	err := Detach(m, "genre_recommendation", "recommendation_id", db.DB)
	if err != nil {
		t.Errorf("Detach error: %s", err)
	}
}

func TestSync(t *testing.T) {
	m := make(map[int64][]int)
	m[1] = []int{2, 3}
	err := Sync(m, "genre_recommendation", "recommendation_id", db.DB)
	if err != nil {
		t.Errorf("Sync error: %s", err)
	}
}

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

func TestIsEmptyTrue(t *testing.T) {
	m := make(map[int64][]int)
	empty := IsEmpty(m)
	if !empty {
		t.Errorf("Expected %t.\n Got %t.", empty, !empty)
	}
}

func TestIsEmptyFalse(t *testing.T) {
	m := make(map[int64][]int)
	m[0] = []int{1, 2, 3}
	empty := IsEmpty(m)
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