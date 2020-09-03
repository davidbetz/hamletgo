package hamlet_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	hamlet "github.com/davidbetz/hamletgo"
)

func TestGetWords(t *testing.T) {
	words := hamlet.GetWords(4)
	if len(strings.Split(words, " ")) != 4 {
		t.Errorf("Wrong word count")
	}
}

func TestGetWordArray(t *testing.T) {
	array := hamlet.GetWordArray(2, 3)
	if len(array) != 2 {
		t.Errorf("Wrong item count")
	}
	for _, words := range array {
		if len(strings.Split(words, " ")) != 3 {
			t.Errorf("Wrong word count")
		}
	}
}

func TestGetObject(t *testing.T) {
	obj := hamlet.GetObject(true)
	if len(obj.ID) == 0 {
		t.Errorf("ID is blank")
	}
}

func TestGetObjectArray(t *testing.T) {
	array := hamlet.GetObjectArray(5)
	if len(array) != 5 {
		t.Errorf("Wrong object count")
	}
	for _, obj := range array {
		if len(obj.ID) == 0 {
			t.Errorf("ID is blank")
		}
	}
}

func ExampleGetWords() {
	words := hamlet.GetWords(4)
	buf, err := json.MarshalIndent(words, " ", "    ")
	if err != nil {
		fmt.Printf("Error serializing: %v", err)
	}
	fmt.Println(string(buf))
}

func ExampleGetWordArray() {
	array := hamlet.GetWordArray(2, 3)
	buf, err := json.MarshalIndent(array, " ", "    ")
	if err != nil {
		fmt.Printf("Error serializing: %v", err)
	}
	fmt.Println(string(buf))
}

func ExampleGetObject() {
	obj := hamlet.GetObject(true)
	buf, err := json.MarshalIndent(obj, " ", "    ")
	if err != nil {
		fmt.Printf("Error serializing: %v", err)
	}
	fmt.Println(string(buf))
}

func ExampleGetObjectArray() {
	array := hamlet.GetObjectArray(5)
	buf, err := json.MarshalIndent(array, " ", "    ")
	if err != nil {
		fmt.Printf("Error serializing: %v", err)
	}
	fmt.Println(string(buf))
}
