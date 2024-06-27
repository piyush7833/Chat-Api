package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/piyush7833/Chat-Api/helpers"
)

func TestPostRequest(t *testing.T, data string, url string, expectedStatusCode int, expectedMessage string, isAuthorized bool) {
	r := InitRoutes()
	ts := httptest.NewServer(r)
	defer ts.Close()

	// fmt.Println(strings.NewReader(data), "post")

	req, err := http.NewRequest("POST", ts.URL+url, strings.NewReader(data))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if isAuthorized {
		req.Header.Set("Cookie", "token="+*token)
	} else {
		req.Header.Set("Cookie", "token=unauthorized token")
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer res.Body.Close()
	if url == "/api/signin" {
		for _, cookie := range res.Cookies() {
			if cookie.Name == "token" {
				token = &cookie.Value
				break
			}
		}
	}
	if res.StatusCode != expectedStatusCode {
		t.Errorf("expected status %d; got %d", expectedStatusCode, res.StatusCode)
		var errorResponse helpers.ErrorResponse
		err = json.NewDecoder(res.Body).Decode(&errorResponse)
		if err != nil {
			// fmt.Println(res.Body)
			t.Fatalf("failed to decode error response: %v", err)
		}
		t.Logf("Error message: %s", errorResponse.Message)
	}
}

func TestPatchRequest(t *testing.T, data string, url string, expectedStatusCode int, expectedMessage string, isAuthorized bool) {
	r := InitRoutes()
	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest("PATCH", ts.URL+url, strings.NewReader(data))

	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if isAuthorized {
		req.Header.Set("Cookie", "token="+*token)
	} else {
		req.Header.Set("Cookie", "token=unauthorized token")
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != expectedStatusCode {
		t.Errorf("expected status %d; got %d", expectedStatusCode, res.StatusCode)
		var errorResponse helpers.ErrorResponse
		err = json.NewDecoder(res.Body).Decode(&errorResponse)
		if err != nil {
			t.Fatalf("failed to decode error response: %v", err)
		}
		t.Logf("Error message: %s", errorResponse.Message)
	}
}

func TestGetRequest(t *testing.T, url string, expectedStatusCode int, expectedMessage string, isAuthorized bool) {
	r := InitRoutes()
	ts := httptest.NewServer(r)
	defer ts.Close()
	req, err := http.NewRequest("GET", ts.URL+url, nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if isAuthorized {
		req.Header.Set("Cookie", "token="+*token)
	} else {
		req.Header.Set("Cookie", "token=unauthorized token")
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != expectedStatusCode {
		t.Errorf("expected status %d; got %d", expectedStatusCode, res.StatusCode)
		var errorResponse helpers.ErrorResponse
		err = json.NewDecoder(res.Body).Decode(&errorResponse)
		if err != nil {
			t.Fatalf("failed to decode error response: %v", err)
		}
		t.Logf("Error message: %s", errorResponse.Message)
	}
}

func TestDeleteRequest(t *testing.T, url string, expectedStatusCode int, expectedMessage string, isAuthorized bool) {
	r := InitRoutes()
	ts := httptest.NewServer(r)
	defer ts.Close()
	req, err := http.NewRequest("DELETE", ts.URL+url, nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if isAuthorized {
		req.Header.Set("Cookie", "token="+*token)
	} else {
		req.Header.Set("Cookie", "token=unauthorized token")
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != expectedStatusCode {
		t.Errorf("expected status %d; got %d", expectedStatusCode, res.StatusCode)
		var errorResponse helpers.ErrorResponse
		err = json.NewDecoder(res.Body).Decode(&errorResponse)
		if err != nil {
			t.Fatalf("failed to decode error response: %v", err)
		}
		t.Logf("Error message: %s", errorResponse.Message)
	}
}
