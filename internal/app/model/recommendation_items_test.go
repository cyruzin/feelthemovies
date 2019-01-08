package model

import "testing"

func TestGetRecommendationItemssModel(t *testing.T) {

	_, err = db.GetRecommendationItems(1)

	if err != nil {
		t.Errorf("GetRecommendationItems error: %s", err)
	}
}

func TestGetRecommendationItemModel(t *testing.T) {
	_, err = db.GetRecommendationItems(1)

	if err != nil {
		t.Errorf("GetRecommendationItem error: %s", err)
	}
}

func TestDeleteRecommendationItemsModel(t *testing.T) {
	_, err = db.DeleteRecommendationItem(3)

	if err != nil {
		t.Errorf("DeleteRecommendationItems error: %s", err)
	}
}
