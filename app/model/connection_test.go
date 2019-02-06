package model

import "testing"

func TestConnection(t *testing.T) {
	_, err := Connect()
	if err != nil {
		t.Errorf("MySQL Connection failed: %s", err)
	}
}
