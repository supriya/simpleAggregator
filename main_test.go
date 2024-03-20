package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestParseFile(t *testing.T) {
	testCustomerEvents := []TestCustomerEventStamp{
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 00:01:11.055+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 00:01:36.259+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 00:03:21.74+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 01:01:11.055+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 01:02:06.397+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 01:02:13.572+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 02:03:22.795+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 02:04:19.217+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 04:06:09.67+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 04:07:14.277+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 03:07:17.652+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 04:00:00.165+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 03:07:29.02+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 03:07:31.628+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 03:07:40.3+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e856", eventTimeString: "2021-03-01 02:07:31.628+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e856", eventTimeString: "2021-03-01 03:07:40.3+00"},
	}
	want := EventBucketIndex{index: make(map[int]map[string]TimeStampRecords)}
	for _, v := range testCustomerEvents {
		eventTime := StringToTime(v.eventTimeString)
		want.AddRecord(v.customerID, eventTime)
	}
	got := CreateIndex("test_events.csv")

	if len(got.index) != len(want.index) {
		t.Errorf("got %v is not equal to want %v", got, want)
	}
	for k, v := range got.index {
		if _, ok := want.index[k]; !ok {
			t.Errorf("got %v is not equal to want %v", got, want)
		}

		for customerid, val := range v {
			wantEvents := want.index[k][customerid].events
			gotEvents := val.events
			sort.Slice(wantEvents, func(i, j int) bool {
				return wantEvents[i].Before(wantEvents[j])
			})
			sort.Slice(gotEvents, func(i, j int) bool {
				return gotEvents[i].Before(gotEvents[j])
			})

			if !reflect.DeepEqual(wantEvents, gotEvents) {
				t.Errorf("got %v is not equal to want %v", gotEvents, wantEvents)
			}
		}
	}
}

func TestParseFileWithMissingRecords(t *testing.T) {
	testCustomerEvents := []TestCustomerEventStamp{
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 00:01:11.055+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 00:01:36.259+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 00:03:21.74+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 01:01:11.055+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 01:02:06.397+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 01:02:13.572+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 02:03:22.795+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 02:04:19.217+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 04:06:09.67+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 04:07:14.277+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 03:07:17.652+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 04:00:00.165+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 03:07:29.02+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 03:07:31.628+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e857", eventTimeString: "2021-03-01 03:07:40.3+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e856", eventTimeString: "2021-03-01 02:07:31.628+00"},
		{customerID: "b4f9279a0196e40632e947dd1a88e856", eventTimeString: "2021-03-01 03:07:40.3+00"},
	}
	want := EventBucketIndex{index: make(map[int]map[string]TimeStampRecords)}
	for _, v := range testCustomerEvents {
		eventTime := StringToTime(v.eventTimeString)
		want.AddRecord(v.customerID, eventTime)
	}
	got := CreateIndex("test_events.csv")

	if len(got.index) != len(want.index) {
		t.Errorf("got %v is not equal to want %v", got, want)
	}
	for k, v := range got.index {
		if _, ok := want.index[k]; !ok {
			t.Errorf("got %v is not equal to want %v", got, want)
		}

		for customerid, val := range v {
			wantEvents := want.index[k][customerid].events
			gotEvents := val.events
			sort.Slice(wantEvents, func(i, j int) bool {
				return wantEvents[i].Before(wantEvents[j])
			})
			sort.Slice(gotEvents, func(i, j int) bool {
				return gotEvents[i].Before(gotEvents[j])
			})

			if !reflect.DeepEqual(wantEvents, gotEvents) {
				t.Errorf("got %v is not equal to want %v", gotEvents, wantEvents)
			}
		}
	}
}
