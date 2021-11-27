package main

import (
	"testing"
)

func AsertMoveEquals(t *testing.T, actual, expected Move) {
	if actual.from != expected.from {
		t.Errorf("Starting squares not equal " + actual.String() + " != " + expected.String())
	}
	if actual.capture != expected.capture {
		t.Errorf("Captures not equal " + actual.String() + " != " + expected.String())
	}
	if actual.promote != expected.promote {
		t.Errorf("Promotions not  " + actual.String() + " != " + expected.String())
	}
}
