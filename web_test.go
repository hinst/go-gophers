package gophers

import (
	"io"
	"net"
	"net/http"
	"strings"
	"testing"
)

// Simulates a transport that fails with a network error after consuming the
// request body, like a connection reset mid-transfer.
type failingTransport struct {
	failures int
	inner    http.RoundTripper
}

func (me *failingTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	if me.failures > 0 {
		me.failures--
		if request.Body != nil {
			_, _ = io.Copy(io.Discard, request.Body)
		}
		return nil, &net.DNSError{Err: "timeout", Name: "example.invalid", IsTimeout: true}
	}
	return me.inner.RoundTrip(request)
}

// Captures the request body and returns a successful response.
type capturingTransport struct {
	body *string
}

func (me *capturingTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	var body, _ = io.ReadAll(request.Body)
	*me.body = string(body)
	return &http.Response{StatusCode: http.StatusOK, Status: "200 OK", Body: io.NopCloser(strings.NewReader("ok"))}, nil
}

func TestWebRetryResendsBody(t *testing.T) {
	var captured string
	var transport = &failingTransport{failures: 2, inner: &capturingTransport{body: &captured}}
	var request = AssertResultError(http.NewRequest(http.MethodPost, "http://example.invalid", strings.NewReader("the-body")))
	var response = AssertResultError(WebRetry{AttemptLimit: 3}.Run(&http.Client{Transport: transport}, request))
	defer IoCloseSilently(response.Body)
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200, got %v", response.StatusCode)
	}
	if captured != "the-body" {
		t.Fatalf("Expected body 'the-body' to be resent on retry, got %q", captured)
	}
}

func TestWebRetryWithoutGetBody(t *testing.T) {
	var captured string
	var transport = &failingTransport{failures: 10, inner: &capturingTransport{body: &captured}}
	var request = AssertResultError(http.NewRequest(http.MethodPost, "http://example.invalid", strings.NewReader("the-body")))
	request.GetBody = nil
	var _, err = WebRetry{AttemptLimit: 3}.Run(&http.Client{Transport: transport}, request)
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
	if captured != "" {
		t.Fatalf("Expected the transport not to be called with a non-restorable body, got %q", captured)
	}
}
