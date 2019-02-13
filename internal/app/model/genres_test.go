package model

import (
	"testing"
	"time"
)

func TestGetGenresModel(t *testing.T) {
	_, err := db.GetGenres()
	if err != nil {
		t.Errorf("GetGenres error: %s", err)
	}
}

func TestGetGenreModel(t *testing.T) {
	_, err := db.GetGenre(1)
	if err != nil {
		t.Errorf("GetGenre error: %s", err)
	}
}

func TestCreateGenreModel(t *testing.T) {
	g := Genre{
		Name:      "NewGenreModel",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := db.CreateGenre(&g)
	if err != nil {
		t.Errorf("CreateGenre error: %s", err)
	}
}

func TestUpdateGenreModel(t *testing.T) {
	g := Genre{
		Name:      "UpGenreModel",
		UpdatedAt: time.Now(),
	}
	_, err := db.UpdateGenre(1, &g)
	if err != nil {
		t.Errorf("UpdateGenre error: %s", err)
	}
}

func TestDeleteGenreModel(t *testing.T) {
	if err := db.DeleteGenre(3); err != nil {
		t.Errorf("DeleteGenre error: %s", err)
	}
}
