package bsonid

import (
	"reflect"
	"testing"
)

func TestGenerateMachineID(t *testing.T) {
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

func TestGetRandomCounter(t *testing.T) {
	objectIDCounter := getRandomCounter()
	if reflect.TypeOf(objectIDCounter).Kind() != reflect.Uint32 {
		t.Errorf("Expected getRandomCounter to return of  type [%d] Got [%d]", reflect.Uint32, reflect.TypeOf(objectIDCounter).Kind())
	}

}
func BenchmarkNew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New()
	}
}
