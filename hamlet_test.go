package hamlet

import (
	"strings"
	"testing"
)

func TestGetWords(t *testing.T) {
	words := GetWords(4)
	if len(strings.Split(words, " ")) != 4 {
		t.Errorf("Wrong word count")
	}
}

func TestGetWordArray(t *testing.T) {
	array := GetWordArray(2, 3)
	if len(array) != 2 {
		t.Errorf("Wrong item count")
	}
	for _, words := range array {
		if len(strings.Split(words, " ")) != 3 {
			t.Errorf("Wrong word count")
		}
	}
}

func TestObject(t *testing.T) {
	obj := GetObject()
	if len(obj.ID) == 0 {
		t.Errorf("ID is blank")
	}
}

func TestGetObjectArray(t *testing.T) {
	array := GetObjectArray(5)
	if len(array) != 5 {
		t.Errorf("Wrong object count")
	}
	for _, obj := range array {
		if len(obj.ID) == 0 {
			t.Errorf("ID is blank")
		}
	}
}
