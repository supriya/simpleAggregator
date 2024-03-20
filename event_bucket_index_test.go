package main

import (
	"reflect"
	"testing"
	"time"
)

type TestCustomerEventStamp struct {
	customerID      string
	eventTimeString string
}

func TestCountEventsByCustomerIdWithCustomerOne(t *testing.T) {
	beginTime := StringToTime("2021-03-01 00:00:13.572+00")
	endTime := StringToTime("2021-03-01 04:00:40.3+00")
	begin := calculateBucketID(beginTime)
	bucketCount := []BucketCount{
		{
			bucketID: begin, count: 3},
		{
			bucketID: begin + 1, count: 3},
		{
			bucketID: begin + 2, count: 2},
		{
			bucketID: begin + 3, count: 4},
		{
			bucketID: begin + 4, count: 1},
	}

	want := EventsOutput{bucketCount}
	index := CreateIndex("test_events.csv")

	customerID := "b4f9279a0196e40632e947dd1a88e857"
	got := index.CountEventsByCustomerId(customerID, beginTime, endTime)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v is not equal to want %v", got, want)
	}
}

func TestCountEventsByCustomerIdWithNoCustomerEvents(t *testing.T) {
	beginTime := StringToTime("2021-03-01 00:00:13.572+00")
	endTime := StringToTime("2021-03-01 04:00:40.3+00")
	begin := calculateBucketID(beginTime)
	bucketCount := []BucketCount{
		{
			bucketID: begin, count: 0},
		{
			bucketID: begin + 1, count: 0},
		{
			bucketID: begin + 2, count: 0},
		{
			bucketID: begin + 3, count: 0},
		{
			bucketID: begin + 4, count: 0},
	}

	want := EventsOutput{bucketCount}
	index := CreateIndex("test_events.csv")

	customerID := "b4f9279a0196e40632e947dd1a88e858"
	got := index.CountEventsByCustomerId(customerID, beginTime, endTime)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v is not equal to want %v", got, want)
	}
}

func TestCountEventsByCustomerIdWithCustomerTwo(t *testing.T) {
	beginTime := StringToTime("2021-03-01 03:00:00.572+00")
	endTime := StringToTime("2021-03-01 04:00:40.3+00")
	begin := calculateBucketID(beginTime)
	bucketCount := []BucketCount{
		{
			bucketID: begin, count: 1},
		{
			bucketID: begin + 1, count: 0},
	}

	want := EventsOutput{bucketCount}
	index := CreateIndex("test_events.csv")

	customerID := "b4f9279a0196e40632e947dd1a88e856"
	got := index.CountEventsByCustomerId(customerID, beginTime, endTime)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v is not equal to want %v", got, want)
	}
}

func TestCountEventsByCustomerIdWithNoEventsForTheTimeRange(t *testing.T) {
	beginTime := StringToTime("2021-03-01 07:00:00.572+00")
	endTime := StringToTime("2021-03-01 08:00:40.3+00")
	begin := calculateBucketID(beginTime)
	bucketCount := []BucketCount{
		{
			bucketID: begin, count: 0},
		{
			bucketID: begin + 1, count: 0},
	}

	want := EventsOutput{bucketCount}
	index := CreateIndex("test_events.csv")

	customerID := "b4f9279a0196e40632e947dd1a88e856"
	got := index.CountEventsByCustomerId(customerID, beginTime, endTime)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v is not equal to want %v", got, want)
	}
}

func TestAddRecord(t *testing.T) {
	eventTime := StringToTime("2021-03-01 00:00:13.572+00")
	bucketID := calculateBucketID(eventTime)
	customerID := "b4f9279a0196e40632e947dd1a88e857"
	index := EventBucketIndex{index: make(map[int]map[string]TimeStampRecords)}
	index.AddRecord(customerID, eventTime)
	want := EventBucketIndex{index: make(map[int]map[string]TimeStampRecords)}
	if _, ok := want.index[bucketID]; !ok {
		want.index[bucketID] = make(map[string]TimeStampRecords)
	}
	if _, ok := want.index[bucketID][customerID]; !ok {
		want.index[bucketID][customerID] = TimeStampRecords{events: []time.Time{eventTime}}
	}

	got := index

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v is not equal to want %v", got, want)
	}
}
