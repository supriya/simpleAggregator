package main

import "time"

type TimeStampRecords struct {
	events []time.Time
}

func (t *TimeStampRecords) CountAllEvents() int {
	return len(t.events)
}

// CountEventsBefore counts the events happened
// before the given input timestamp.
func (t *TimeStampRecords) CountEventsBefore(input time.Time) int {
	bucketEventCount := 0
	for _, t := range t.events {
		if t.Before(input) {
			bucketEventCount += 1
		}
	}
	return bucketEventCount
}

// CountEventsBefore counts the events happened
// before the given input timestamp.
func (t *TimeStampRecords) CountEventsAfter(input time.Time) int {
	bucketEventCount := 0
	for _, t := range t.events {
		if t.After(input) {
			bucketEventCount += 1
		}
	}
	return bucketEventCount
}

func (t *TimeStampRecords) Append(event time.Time) {
	t.events = append(t.events, event)
}
