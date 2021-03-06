package main

import (
	"strconv"
	"testing"
	"time"
)

func TestGetSleepTime(t *testing.T) {
	fakeNow := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)

	// Zero
	result := getSleepTime(getTimeAsString(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)), fakeNow)
	expected := 0 * time.Second
	if result != expected {
		t.Errorf("Expected %s got %s", expected, result)
	}

	// Default
	result = getSleepTime(getTimeAsString(time.Date(2021, 1, 1, 0, 0, 10, 0, time.UTC)), fakeNow)
	expected = 10 * time.Second
	if result != expected {
		t.Errorf("Expected %s got %s", expected, result)
	}

	// Random junk
	result = getSleepTime("random-string-of-rubbish", fakeNow)
	expected = 60 * time.Second
	if result != expected {
		t.Errorf("Expected %s got %s", expected, result)
	}

	// Negative
	result = getSleepTime(getTimeAsString(time.Date(2020, 12, 30, 0, 0, 10, 0, time.UTC)), fakeNow)
	expected = 0 * time.Second
	if result != expected {
		t.Errorf("Expected %s got %s", expected, result)
	}

}

func getTimeAsString(date time.Time) string {
	return strconv.FormatInt(date.Unix(), 10)
}
