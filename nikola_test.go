package nikola_test

import (
	"github.com/deanveloper/nikola"
	"time"
	"testing"
)

// Example which shows that "Assert" will halt the test if
// it fails. Does this by asserting several things, one of which is false.
func ExampleAssertTrue() {
	t := new(testing.T)

	nikola.AssertTrue(t, 5*time.Hour < 1000*time.Minute)
	nikola.AssertTrue(t, time.Millisecond < time.Second)
	nikola.AssertTrue(t, time.Hour < 50*time.Minute)
	// Execution stops here as the test has failed.

	// Never gets run as the previous assertion failed.
	nikola.AssertTrue(t, time.Second > 4000*time.Nanosecond)
}

// Shows how AssertEquals can be used to provide more
// detailed error messages within tests than AssertTrue.
func ExampleAssertEqual() {
	t := new(testing.T)

	nikola.AssertEqual(t, "1ms", time.Millisecond.String())
	nikola.AssertEqual(t, "5h0m0s", (5 * time.Hour).String())
	nikola.AssertEqual(t, "10m0s", (20 * time.Minute).String())
	// Execution stops here as the test has failed.

	// Never gets run as the previous assertion failed.
	nikola.AssertEqual(t, "24h0m0s", (24 * time.Hour).String())
}

// Example which shows that "Suggest" will not halt the test if
// it fails.
func ExampleSuggestTrue() {
	t := new(testing.T)

	nikola.SuggestTrue(t, 5*time.Hour < 1000*time.Minute)
	nikola.SuggestTrue(t, time.Millisecond < time.Second)
	nikola.SuggestTrue(t, time.Hour < 50*time.Minute)
	// This suggestion is false, which is logged, and then execution continues..

	nikola.SuggestTrue(t, time.Second > 4000*time.Nanosecond)
	nikola.SuggestTrue(t, time.Second > 4000*time.Millisecond)
	// This suggestion is also false, and also gets logged.
}

// Shows how AssertEquals can be used to provide more
// detailed error messages within tests than AssertTrue.
func ExampleSuggestEqual() {
	t := new(testing.T)

	nikola.SuggestEqual(t, "1ms", time.Millisecond.String())
	nikola.SuggestEqual(t, "5h0m0s", (5 * time.Hour).String())
	nikola.SuggestEqual(t, "10m0s", (20 * time.Minute).String())
	// Error is logged, but execution still continues.

	nikola.SuggestEqual(t, "24h0m0s", (24 * time.Hour).String())
	nikola.SuggestEqual(t, "5s", (50000 * time.Millisecond).String())
	// This error is also logged.
}
