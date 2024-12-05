package main

import "testing"

func TestMain(t *testing.T) {
	expected := "Hello, CI CD with Jenkins!"

	if expected != "Hello, CI CD with Jenkins!" {
		t.Fatalf("Expected %s, but got %s", expected, "Hello, CI/CD with Jenkins!")
	}

}
