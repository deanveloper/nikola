package nikola

import "testing"

func AssertTrue(t *testing.T, expr bool, ifFalse func()) {
	t.Helper()
	if !expr {
		ifFalse()
	}
}

func AssertFalse(t *testing.T, expr bool, ifTrue func()) {
	t.Helper()
	AssertTrue(t, !expr, ifTrue)
}

func AssertNotEqual(t *testing.T, expected, actual interface{}, ifEq func()) {
	t.Helper()
	AssertFalse(t, expected == actual, ifEq)
}

func AssertEqual(t *testing.T, expected, actual interface{}, notEq func()) {
	t.Helper()
	AssertTrue(t, expected == actual, notEq)
}

func AssertTrueWithError(t *testing.T, expr bool) {
	t.Helper()
	AssertTrue(t, expr, func() { t.Errorf("expr is false, but should be true") })
}

func AssertFalseWithError(t *testing.T, expr bool) {
	t.Helper()
	AssertFalse(t, expr, func() { t.Errorf("expr is true, but should be false") })
}

func AssertEqualWithError(t *testing.T, expected, actual interface{}) {
	t.Helper()
	AssertIntEqual(t, expected, actual, func() { t.Errorf("Actual: %v, Expected: %v\n", expected, actual) })
}

func AssertEqualWithFatal(t *testing.T, expected, actual interface{}) {
	t.Helper()
	AssertIntEqual(t, expected, actual, func() { t.Errorf("Actual: %v, Expected: %v\n", expected, actual) })
}
