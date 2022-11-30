package errstack

import (
	"os"
	"strings"
	"testing"
)

func TestBasic(t *testing.T) {
	// 1. First error
	e := New("password field is missing")

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
	e := New("password field is missing")

	// 2. And then the second error occured.
	e.Append("company name is missing")

	e.PopAll()

	if len(e.stack) != 0 {
		t.Fatal("after PopAll, the stack should be empty")
	}
}

func TestNoMetadata(t *testing.T) {
	// 1. First error
	e := New("password field is missing")

	// 2. And then the second error occured.
	e.Append("company name is missing")

	// 3. And then the third error occured.
	e.Append("username is too short")

	if len(e.stack) != 3 {
		t.Fatalf("There should be 3 errors. Got: %v", len(e.stack))
	}

	e.SetShowMetadata(false)

	if e.showMetadata != false {
		t.Fatalf("Failed to set showMetadata")
	}

	expected := "username is too short, company name is missing, password field is missing"

	if e.Error() != expected {
		t.Fatalf("Error string incorrect.\nexpected: %v,\ngot: %v", expected, e.Error())
	}
}

func TestTrimFilename(t *testing.T) {
	// 1. First error
	e := New("password field is missing")

	// 2. And then the second error occured.
	e.Append("company name is missing")

	// 3. And then the third error occured.
	e.Append("username is too short")

	if len(e.stack) != 3 {
		t.Fatalf("There should be 3 errors. Got: %v", len(e.stack))
	}

	e.SetTrimFilename(true)

	if e.trimFilename != true {
		t.Fatalf("Failed to set trimFilename")
	}

	expected1 := `github.com/didip/errstack/errstack_test.go:96="username is too short" github.com/didip/errstack/errstack_test.go:93="company name is missing" github.com/didip/errstack/errstack.go:24="password field is missing"`
	expected2 := `/home/runner/work/errstack/errstack/errstack_test.go:96="username is too short" /home/runner/work/errstack/errstack/errstack_test.go:93="company name is missing" /home/runner/work/errstack/errstack/errstack.go:24="password field is missing"`
	expected3 := `/Users/runner/work/errstack/errstack/errstack_test.go:96="username is too short" /Users/runner/work/errstack/errstack/errstack_test.go:93="company name is missing" /Users/runner/work/errstack/errstack/errstack.go:24="password field is missing"`

	// expected2 is when Github Action is running
	if (e.Error() != expected1) && (e.Error() != expected2) && (e.Error() != expected3) {
		t.Fatalf("Error string incorrect.\nexpected: %v,\nor: %v,\nor: %v,\ngot: %v", expected1, expected2, expected3, e.Error())
	}

	e.SetTrimFilename(false)

	if !strings.HasPrefix(e.Error(), string(os.PathSeparator)) {
		t.Fatalf("Error string incorrect. Expected e.Error() to start with %v", string(os.PathSeparator))
	}
}
