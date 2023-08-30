package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKeyEmpty(t *testing.T) {
	headers := http.Header{}
	_, err := GetAPIKey(headers)

	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Error("GetAPIKey did not return the correct error for an empty header")
	}
}

func TestGetAPIKeyMalformed(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "this is not right")

	_, err := GetAPIKey(headers)

	if err == nil {
		t.Error("GetAPIKey did not return the correct error for a malformed header")
	}
}

func TestGetAPIKeyValid(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey thisisadummyapikey")

	apiKey, err := GetAPIKey(headers)

	if err != nil {
		t.Error("GetAPIKey returned an error when it shouldn't have")
	}

	if apiKey != "thisisadummyapikey" {
		t.Errorf(`apiKey was "%v" when it should have been "%v"`, apiKey, "thisisadummyapikey")
	}
}
