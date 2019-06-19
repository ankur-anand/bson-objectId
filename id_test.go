package id

import (
	"testing"
)

func TestGenerateMachineID(t  *testing.T) {
	id := generateMachineID()
	if len(id) != 3 {
		t.Errorf("Expected len of generated machine Id [%d] Got [%d]", 3, len(id))
	}
}

func TestNew(t *testing.T) {
	objectID := New()
	if len(objectID) != 24 {
		t.Errorf("Expected len of generated Object Id [%d] Got [%d]", 24, len(objectID))
	}
}