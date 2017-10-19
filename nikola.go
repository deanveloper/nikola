// A package to make testing easier. Requires Go 1.9.
package nikola

import "testing"

// The template that is used when two things should equal each other.
var ExpectEqualsTemplate string = "expected : %v\nactual   : %v"

// The template that is used when a value should not be the value specified.
var ExpectNotEqualsTemplate string = "expected : !%v\nactual   : %v"

// Asserts that a value should be true. Calls testing.Fatalf if false.
func AssertTrue(t *testing.T, expr bool) {
	t.Helper()
	if !expr {
		t.Fatalf(ExpectedEqualsTemplate, true, expr)
	}
}

// Asserts that a value should be false. Calls testing.Fatalf if true.
func AssertFalse(t *testing.T, expr bool) {
	t.Helper()
	if expr {
		t.Fatalf(ExpectedEqualsTemplate, false, expr)
	}
}

// Asserts that two values should be equal. Calls testing.Fatalf if not equal.
func AssertEqual(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if expected != actual {
		t.Fatalf(ExpectedEqualsTemplate, expected, actual)
	}
}

// Asserts that two values should not be equal. Calls testing.Fatalf if equal.
func AssertNotEqual(t *testing.T, expectedNot, actual interface{}) {
	t.Helper()
	if expected == actual {
		t.Fatalf(ExpectedNotEqualsTemplate, expectedNot, actual)
	}
}

// Suggests that the provided value should be true. Calls testing.Errorf if false.
func SuggestTrue(t *testing.T, expr bool) {
	t.Helper()
	if !expr {
		t.Errorf(ExpectedEqualsTemplate, true, expr)
	}
}

// Suggests that the provided value should be false. Calls testing.Errorf if true.
func SuggestFalse(t *testing.T, expr bool) {
	t.Helper()
	if expr {
		t.Errorf(ExpectedEqualsTemplate, false, expr)
	}
}

// Suggests that two values should be equal. Calls testing.Errorf if not equal.
func SuggestEqual(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if expected != actual {
		t.Errorf(ExpectedEqualsTemplate, expected, actual)
	}
}

// Suggests that two values should not be equal. Calls testing.Errorf if equal.
func SuggestNotEqual(t *testing.T, expectedNot, actual interface{}) {
	t.Helper()
	if expected == actual {
		t.Errorf(ExpectedNotEqualsTemplate, expectedNot, actual)
	}
}
