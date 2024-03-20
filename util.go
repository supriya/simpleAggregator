package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

// StringToTime as the name suggests converts
// timestamp string to time. Golang time library does
// not support the time parsing with this particular
// format, hence the string manipulation.
func StringToTime(eventTimeStamp string) time.Time {
	dateTime := strings.Split(eventTimeStamp, " ")
	eventDate := dateTime[0]
	eventTime := dateTime[1]

	hour := strings.Split(eventTime, ":")[0]
	hourInt, err := strconv.Atoi(hour)
	if err != nil {
		log.Fatal("Error converting hour", err)
	}
	minute := strings.Split(eventTime, ":")[1]
	minuteInt, err := strconv.Atoi(minute)
	if err != nil {
		log.Fatal("Error converting Minute", err)
	}

	second := strings.Split(eventTime, ":")[2]
	secondInt, err := strconv.ParseFloat(strings.Split(second, "+")[0], 64)
	if err != nil {
		log.Fatal("Error converting Second", err, eventTime, eventTimeStamp)
	}
	dateParts := strings.Split(eventDate, "-")
	year, err := strconv.Atoi(dateParts[0])
	if err != nil {
		log.Fatal("Error converting year", err)
	}
	month, err := strconv.Atoi(dateParts[1])
	if err != nil {
		log.Fatal("Error converting year", err)
	}

	day, err := strconv.Atoi(dateParts[2])
	if err != nil {
		log.Fatal("Error converting year", err)
	}

	return time.Date(year, time.Month(month), day, hourInt, minuteInt, int(secondInt), 0, time.UTC)
}
