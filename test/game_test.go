package tests

import (
	. "lovelace/source"
	"testing"
)

func AsertMoveEquals(t *testing.T, actual, expected Move) {
	if actual.From != expected.From {
		t.Errorf("Starting squares not equal " + actual.String() + " != " + expected.String())
	}
	if actual.Capture != expected.Capture {
		t.Errorf("Captures not equal " + actual.String() + " != " + expected.String())
	}
	if actual.Promote != expected.Promote {
		t.Errorf("Promotions not  " + actual.String() + " != " + expected.String())
	}
}
