package main

import (
	"testing"
)

func TestGetFilename(t *testing.T) {
	path := "/home/img/avatar.png"
	expected := "avatar.png"
	if getFilename(path) != expected {
		t.Fatal("Something is not ok!")
	}
}
