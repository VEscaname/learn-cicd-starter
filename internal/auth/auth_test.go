package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://api.example.com/data", nil)
	req.Header.Set("Authorization", "ApiKey TheAPIKey")
	got, err := GetAPIKey(req.Header)
	if err != nil {
		t.Fatalf("Error from GetAPIKey: %v", err)
	}
	want := "TheAPIKey"

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}

}
