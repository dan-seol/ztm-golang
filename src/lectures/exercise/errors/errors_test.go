package timeparse

import "testing"

func TestSuccessfulParsing(t *testing.T) {
	s := "14:07:33"
	timestamp, err := parse(s)
	if err != nil {
		t.Errorf("%v should not fail while converting", s)
	}
	if !((timestamp.hour == 14) && (timestamp.minute == 7) && (timestamp.second == 33)) {
		t.Errorf("%v has an incorrect value", s)
	}
}

func TestFailInvalidSplit(t *testing.T) {
	s := "14:07"
	_, err := parse(s)
	if err == nil {
		t.Errorf("%v should fail while converting", s)
	}
}

func TestFailInvalidComponent(t *testing.T) {
	s := "14:07:61"
	_, err := parse(s)
	if err == nil {
		t.Errorf("%v should fail while converting", s)
	}
}

func TestFailInvalidValue(t *testing.T) {
	s := "hi:hi:hi"
	_, err := parse(s)
	if err == nil {
		t.Errorf("%v should fail while converting", s)
	}
}
