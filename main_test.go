package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRedirectHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Redirect))
	defer server.Close()

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}

	resp, err := client.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusTemporaryRedirect {
		t.Fatalf("Received non-307 response: %d\n", resp.StatusCode)
	}
	url, _ := resp.Location()

	if url.Scheme != "https" {
		t.Fatalf("Redirect to the wrong HTTPS Scheme: %s\n", url.Scheme)
	}

	if url.Host != "127.0.0.1:443" {
		t.Fatalf("Redirect to the wrong HTTPS Host: %s\n", url.Host)
	}
}
