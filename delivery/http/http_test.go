package main_test

import (
	"bytes"
	"github.com/derrickwilliams/go-clean-architecture/delivery/http"
	"go"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

type MockRoundTripper struct {
	originalTripper http.RoundTripper
	fakeBodyText    string
	targetURL       string
	*testing.T
}

func makeResponseBody(b string) io.ReadCloser {
	bodybytes := []byte(b)
	return ioutil.NopCloser(bytes.NewReader(bodybytes))
}

func (m MockRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	var body io.ReadCloser
	m.Logf("WHAT!!! %s", r.Host)

	if r.Host == m.targetURL {
		body = makeResponseBody(m.fakeBodyText)
		return &http.Response{Body: body}, nil
	} else {
		return m.originalTripper.RoundTrip(r)
	}
}

func TestHttpDeliveryBootstrap(t *testing.T) {
	mockTripper := MockRoundTripper{originalTripper: http.DefaultTransport, targetURL: "www.googasdfsle.com", fakeBodyText: "no google for you"}
	http.DefaultTransport = mockTripper

	body := myhttp.RequestGoogle()
	bodystring := string(body[:len(body)])
	t.Logf("%+v", bodystring)
}
