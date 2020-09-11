package pagerduty

import (
	"net/http"
	"testing"
)

func TestEventV2_ManageEvent(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/enqueue", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.Write([]byte(`{"status": "ok", "dedup_key": "yes", "message": "ok"}`))
	})
	var client = &Client{apiEndpoint: server.URL, authToken: "foo", HTTPClient: defaultHTTPClient}
	evt := &V2Event{
		RoutingKey: "abc123",
	}
	res, err := client.ManageEvent(evt)
	if err != nil {
		t.Fatal(err)
	}
	want := &V2EventResponse{
		Status:   "ok",
		DedupKey: "yes",
		Message:  "ok",
	}
	testEqual(t, want, res)
}
