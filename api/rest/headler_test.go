package rest

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestGetIp(t *testing.T) {
	// Create a request to pass to the handler
	request, err := http.NewRequest("POST", "/", strings.NewReader("ip=192.168.0.1"))
	if err != nil {
		t.Fatal(err)
	}
	request.PostForm = url.Values{"ip": {"192.168.0.1"}}

	// Create a ResponseRecorder to record the response
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(getIp)

	// Call the handler with the request and response recorder
	handler.ServeHTTP(responseRecorder, request)

	// Check the status code
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := `{"ok":true,"message":"Ip blocked 192.168.0.1"}`

	// Del /n
	if strings.TrimSpace(responseRecorder.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", responseRecorder.Body.String(), expected)
	}
}

func TestGetIpBadRequest(t *testing.T) {
	// Create a request to pass to the handler
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(getIp)

	// Call the handler with the request and response recorder
	handler.ServeHTTP(responseRecorder, request)

	// Check the status code
	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	// Check the response body
	expected := `{"ok":false,"message":"Bad Request"}`
	if strings.TrimSpace(responseRecorder.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", responseRecorder.Body.String(), expected)
	}
}
