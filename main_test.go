// Copyright (c) avijit bhattacharjee 2024

package main

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortenURL(t *testing.T) {
	reqBody := URLShortenRequest{URL: "http://example.com"}

	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("POST", "/shorten", bytes.NewBuffer(reqJSON))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	shortenURL(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be 200")

	var resp map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "URL shortened successfully", resp["message"], "Unexpected message")
	assert.NotEmpty(t, resp["id"], "ID should not be empty")
	assert.Equal(t, extractDomain(reqBody.URL), resp["domain"], "Unexpected domain")
}

func TestGetTopDomains(t *testing.T) {
	// Initialize some domain counts for testing
	domainCounts["example.com"] = 5
	domainCounts["test.com"] = 3
	domainCounts["google.com"] = 7

	// Create a GET request
	req := httptest.NewRequest("GET", "/metrics/top-domains", nil)

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function directly
	getTopDomains(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be 200")

	// Decode the response body
	var resp []DomainCount
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	// Assert the response content
	assert.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, 3, len(resp), "Expected 3 domains")
	assert.Equal(t, "google.com", resp[0].Domain, "First domain should be google.com")
	assert.Equal(t, 7, resp[0].Count, "Count for google.com should be 7")
	// Add more assertions based on expected behavior
}

func TestRedirectURL(t *testing.T) {
	// Populate shortenedURLs with a test entry
	testID := "abc123"
	testURL := "http://example.com"
	shortenedURLs[testID] = testURL

	// Create a GET request with the mock ID
	req := httptest.NewRequest("GET", "/"+testID, nil)
	req = mux.SetURLVars(req, map[string]string{"id": testID})

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function directly
	redirectURL(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusFound, rr.Code, "Status code should be 302")

	// Check the location header for the redirected URL
	assert.Equal(t, testURL, rr.Header().Get("Location"), "Unexpected redirect location")
}
