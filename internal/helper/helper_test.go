package helper

import (
	"testing"
	"time"

	"github.com/cyruzin/feelthemovies/internal/model"
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

	res, err := ToJSON(rec)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
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

	res, err := ToJSONIndent(rec)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}
