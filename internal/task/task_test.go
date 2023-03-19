package task

import (
	"testing"
)

func TestNewTask(t *testing.T) {
	var id string = "34"
	var name string = "Leer libro"
	var content string = "Clean Code"

	_, err := NewTask(id, name, content)

	if err != nil {
		t.Errorf(err.Error())
	}
}
