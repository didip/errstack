package errstack

import (
	"strings"
	"testing"
)

func TestBasic(t *testing.T) {
	// 1. First error
	e := NewString("password field is missing")

	// 2. And then the second error occured.
	e.Append("company name is missing")

	// 3. And then the third error occured.
	e.Append("username is too short")

	if len(e.stack) != 3 {
		t.Fatalf("There should be 3 errors. Got: %v", len(e.stack))
	}

	if e.Error() == "" {
		t.Fatalf("error string should not be empty")
	}

	for i, err := range e.GetAll() {
		if err.filename == "" {
			t.Fatalf("The filename metadata should not be empty")
		}
		if err.line == 0 {
			t.Fatalf("The line metadata should not be empty")
		}
		if err.err == "" {
			t.Fatalf("The error string should not be empty")
		}

		if i == 0 && !strings.Contains(err.Error(), "username is too short") {
			t.Fatalf("The last error should be printed first")
		}

		if i == 2 && !strings.Contains(err.Error(), "password field is missing") {
			t.Fatalf("The first error should be printed last")
		}
	}
}

func TestPopAll(t *testing.T) {
	e := NewString("password field is missing")

	// 2. And then the second error occured.
	e.Append("company name is missing")

	e.PopAll()

	if len(e.stack) != 0 {
		t.Fatal("after PopAll, the stack should be empty")
	}
}

func TestNoMetadata(t *testing.T) {
	// 1. First error
	e := NewString("password field is missing")

	// 2. And then the second error occured.
	e.Append("company name is missing")

	// 3. And then the third error occured.
	e.Append("username is too short")

	if len(e.stack) != 3 {
		t.Fatalf("There should be 3 errors. Got: %v", len(e.stack))
	}

	e.SetShowMetadata(true)

	if e.showMetadata != true {
		t.Fatalf("Failed to set showMetadata")
	}

	expected := "username is too short, company name is missing, password field is missing"

	if e.Error() != expected {
		t.Fatalf("Error string incorrect. expected: %v, got: %v", expected, e.Error())
	}
}
