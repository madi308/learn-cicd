package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey1(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey test_key")
	got, err := GetAPIKey(headers)
	if err != nil {
		t.Fatal(err)
	}
	want := "test_key"
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestGetAPIKey2(t *testing.T) {
	headers := http.Header{}
	_, err := GetAPIKey(headers)
	if !reflect.DeepEqual(ErrNoAuthHeaderIncluded.Error(), err.Error()) {
		t.Fatalf("expected: %v, got: %v", ErrNoAuthHeaderIncluded, err)
	}
}

func TestGetAPIKey3(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "notapikey test_key")
	_, err := GetAPIKey(headers)
	want := "malforme authorization header"
	if !reflect.DeepEqual(want, err.Error()) {
		t.Fatalf("expected: %v, got: %v", want, err)
	}
}
