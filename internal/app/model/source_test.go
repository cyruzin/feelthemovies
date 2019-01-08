package model

import (
	"testing"
	"time"
)

func TestGetSourcesModel(t *testing.T) {

	_, err = db.GetSources()

	if err != nil {
		t.Errorf("GetSources error: %s", err)
	}
}

func TestGetSourceModel(t *testing.T) {
	_, err = db.GetSource(1)

	if err != nil {
		t.Errorf("GetSource error: %s", err)
	}
}

func TestCreateSourceModel(t *testing.T) {
	g := Source{
		Name:      "NewSourceModel",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = db.CreateSource(&g)

	if err != nil {
		t.Errorf("CreateSource error: %s", err)
	}
}

func TestUpdateSourceModel(t *testing.T) {
	g := Source{
		Name:      "UpSourceModel",
		UpdatedAt: time.Now(),
	}
	_, err = db.UpdateSource(1, &g)

	if err != nil {
		t.Errorf("UpdateSource error: %s", err)
	}
}

func TestDeleteSourceModel(t *testing.T) {
	_, err = db.DeleteSource(3)

	if err != nil {
		t.Errorf("DeleteSource error: %s", err)
	}
}
