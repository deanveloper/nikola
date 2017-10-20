package nikola_test

import (
	"github.com/deanveloper/nikola"
	"time"
)

func ExampleAssertTrue() {
	// Example which shows that "Assert" will halt the test if
	// it fails. Does this by asserting several things, one of which is false.

	nikola.AssertTrue(t, 5*time.Hour < 1000*time.Minute)
	nikola.AssertTrue(t, time.Millisecond < time.Second)
	nikola.AssertTrue(t, time.Hour < 50*time.Minute)
	// Execution stops here as the test has failed.

	// Never gets run as the previous assertion failed.
	nikola.AssertTrue(t, time.Second > 4000*time.Nanosecond)

	// Output:
	// expected : true
	// actual   : false
	// <details about error from gotest>
}

func ExampleAssertEquals() {
	// Shows how AssertEquals can be used to provide more
	// detailed error messages within tests than AssertTrue.

	nikola.AssertEquals(t, "1ms", time.Millisecond.String())
	nikola.AssertEquals(t, "5h0m0s", (5 * time.Hour).String())
	nikola.AssertEquals(t, "10m0s", (20 * time.Minute).String())
	// Execution stops here as the test has failed.

	// Never gets run as the previous assertion failed.
	nikola.AssertEquals(t, "24h0m0s", (24 * time.Hour).String())

	// Output:
	// expected : 10m0s
	// actual   : 20m0s
	// <details about error from gotest>
}

func ExampleSuggestTrue() {
	// Example which shows that "Suggest" will not halt the test if
	// it fails.

	nikola.SuggestTrue(t, 5*time.Hour < 1000*time.Minute)
	nikola.SuggestTrue(t, time.Millisecond < time.Second)
	nikola.SuggestTrue(t, time.Hour < 50*time.Minute)
	// This suggestion is false, which is logged, and then execution continues..

	nikola.SuggestTrue(t, time.Second > 4000*time.Nanosecond)
	nikola.SuggestTrue(t, time.Second > 4000*time.Millisecond)
	// This suggestion is also false, and also gets logged.

	// Output:
	// expected : true
	// actual   : false
	// <details about error from gotest>
	// expected : true
	// actual   : false
	// <details about second error from gotest>
}

func ExampleSuggestEquals() {
	// Shows how AssertEquals can be used to provide more
	// detailed error messages within tests than AssertTrue.

	nikola.SuggestEquals(t, "1ms", time.Millisecond.String())
	nikola.SuggestEquals(t, "5h0m0s", (5 * time.Hour).String())
	nikola.SuggestEquals(t, "10m0s", (20 * time.Minute).String())
	// Error is logged, but execution still continues.

	nikola.SuggestEquals(t, "24h0m0s", (24 * time.Hour).String())
	nikola.SuggestEquals(t, "5s", (50000 * time.Millisecond).String())
	// This error is also logged.

	// Output:
	// expected : 10m0s
	// actual   : 20m0s
	// <details from gotest about error>
	// expected : 5s
	// actual   : 50s
	// <details about second error>
}
