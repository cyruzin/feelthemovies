package model

import (
	"testing"
	"time"
)

func TestGetKeywordsModel(t *testing.T) {
	_, err = db.GetKeywords()
	if err != nil {
		t.Errorf("GetKeywords error: %s", err)
	}
}

func TestGetKeywordModel(t *testing.T) {
	_, err = db.GetKeyword(1)
	if err != nil {
		t.Errorf("GetKeyword error: %s", err)
	}
}

func TestCreateKeywordModel(t *testing.T) {
	g := Keyword{
		Name:      "NewKeywordModel",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = db.CreateKeyword(&g)
	if err != nil {
		t.Errorf("CreateKeyword error: %s", err)
	}
}

func TestUpdateKeywordModel(t *testing.T) {
	g := Keyword{
		Name:      "UpKeywordModel",
		UpdatedAt: time.Now(),
	}
	_, err = db.UpdateKeyword(1, &g)
	if err != nil {
		t.Errorf("UpdateKeyword error: %s", err)
	}
}

func TestDeleteKeywordModel(t *testing.T) {
	_, err = db.DeleteKeyword(3)
	if err != nil {
		t.Errorf("DeleteKeyword error: %s", err)
	}
}
