package gocurl

import (
	"testing"
)


func TestGet(t *testing.T) {

	url := "https://mt-intgrapi.igp.cloud/api/v1/portal/affiliateFeed/goa/registrations/2020-02-13"

	headers := map[string]string{
		"Authorization" : "22c3b836-10b4-4729-9f78-f67f0b89e926",
	}

	want := 200

	if got, _:= Get(url, headers); got != want {
		t.Errorf("Get() = %q, want %q", got, want)
	}

}