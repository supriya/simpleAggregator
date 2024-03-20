package main

import (
	"testing"
	"time"
)

func TestCountEventsBefore(t *testing.T) {
	timeStampStrings := []string{"2021-03-01 00:03:21.74+00", "2021-03-01 01:02:06.397+00", "2021-03-01 02:02:06.397+00"}
	timeStamps := []time.Time{}
	for _, t := range timeStampStrings {
		timeStamps = append(timeStamps, StringToTime(t))
	}
	timeStampRecord := TimeStampRecords{events: timeStamps}

	inputTimeStamp := StringToTime("2021-03-01 02:02:06.397+00")
	got := timeStampRecord.CountEventsBefore(inputTimeStamp)
	want := 2
	if got != want {
		t.Errorf("got %v is not equal to want %v", got, want)
	}
}

func TestCountEventsBeforeWithNoEvents(t *testing.T) {
	timeStampStrings := []string{"2021-03-01 01:03:21.74+00", "2021-03-01 01:20:06.397+00", "2021-03-01 02:02:06.397+00"}
	timeStamps := []time.Time{}
	for _, t := range timeStampStrings {
		timeStamps = append(timeStamps, StringToTime(t))
	}
	timeStampRecord := TimeStampRecords{events: timeStamps}

	inputTimeStamp := StringToTime("2021-03-01 00:02:06.397+00")
	got := timeStampRecord.CountEventsBefore(inputTimeStamp)
	want := 0
	if got != want {
		t.Errorf("got %v is not equal to want %v", got, want)
	}
}

func TestCountEventsAfter(t *testing.T) {
	timeStampStrings := []string{"2021-03-01 00:03:21.74+00", "2021-03-01 01:02:06.397+00", "2021-03-01 02:02:06.397+00"}
	timeStamps := []time.Time{}
	for _, t := range timeStampStrings {
		timeStamps = append(timeStamps, StringToTime(t))
	}
	timeStampRecord := TimeStampRecords{events: timeStamps}

	inputTimeStamp := StringToTime("2021-03-01 00:02:06.397+00")
	got := timeStampRecord.CountEventsAfter(inputTimeStamp)
	want := 3
	if got != want {
		t.Errorf("got %v is not equal to want %v", got, want)
	}
}

func TestCountEventsAfterWithNoEvents(t *testing.T) {
	timeStampStrings := []string{"2021-03-01 01:03:21.74+00", "2021-03-01 01:20:06.397+00", "2021-03-01 02:02:06.397+00"}
	timeStamps := []time.Time{}
	for _, t := range timeStampStrings {
		timeStamps = append(timeStamps, StringToTime(t))
	}
	timeStampRecord := TimeStampRecords{events: timeStamps}

	inputTimeStamp := StringToTime("2021-03-01 02:03:06.397+00")
	got := timeStampRecord.CountEventsAfter(inputTimeStamp)
	want := 0
	if got != want {
		t.Errorf("got %v is not equal to want %v", got, want)
	}
}
