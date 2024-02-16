package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alvaroglvn/ravensfield-collection/initiate"
)

func TestGetDalleImg(t *testing.T) {

	router := initiate.RouterInit()

	testServer := httptest.NewServer(router)
	defer testServer.Close()

	response, err := http.Get(testServer.URL + "/img")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer response.Body.Close()

	got := response.StatusCode
	want := 201

	if got != want {
		t.Errorf("got status code %v, wanted %v", got, want)
	}
}
