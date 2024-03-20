# Simple Aggregator

Attached, find the file `events.csv`, which contains a log of events with the
format customer\_id, event\_type, transaction\_id, timestamp.

Simple Aggregator program answers the following question:

How many events did customer X send in the one hour buckets between arbitrary timestamps A and B?

So, for example, let's say you have the following usage events (this is just example data -- see the csv file for test data):

// - 2022-03-01T03:01:00Z event_1 customer_id_1
// - 2022-03-01T04:29:00Z event_2 customer_id_1
// - 2022-03-01T04:15:00Z event_3 customer_id_1
// - 2022-03-01T05:08:00Z event_4 customer_id_1

If you requested counts for customer_id_1 with start and end timestamps of Mar 1, 2022 at 3:00 am - Mar 1, 2022 at 6:00 am, we’d expect to see these hourly counts (the format of the output is up to you):

// - 2022-03-01T03:00:00Z bucket -> 1
// - 2022-03-01T04:00:00Z bucket -> 2
// - 2022-03-01T05:00:00Z bucket -> 1
