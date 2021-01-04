// Copyright 2020 Hewlett Packard Enterprise Development LP

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	testBaseURL = "http://localhost:27778/boot/v1"
)

// This test case verifies that the base URL endpoint responds.  This endpoint
// is used by BOS to verify BSS is available.
func TestBaseEndpoint(t *testing.T) {
	testURL := testBaseURL + "/"
	req, err := http.NewRequest("GET", testURL, nil)
	if err != nil {
		t.Fatal("Cannot create http request:", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Index)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("HTTP handler \"Index\" for URL \"%s\" returned wrong error code, got %v, expected %v\n", testURL, rr.Code, http.StatusOK)
	}
}
