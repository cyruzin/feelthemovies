package logger

import (
	"testing"
)

func TestInitLogger(t *testing.T) {
	_, err := Init()
	if err != nil {
		t.Fatal("Could not initiate the logger: " + err.Error())
	}
}
