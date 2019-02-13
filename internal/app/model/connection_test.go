package model

import "testing"

func TestConnection(t *testing.T) {
	if err := Connect(); err.Ping() != nil {
		t.Error("MySQL Connection failed")
	}
}
