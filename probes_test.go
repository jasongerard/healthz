package healthz

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	alive = "/healthz/alive"
	ready = "/healthz/ready"
)

func TestDefaults(t *testing.T) {

	mux := CreateMux()

	server := httptest.NewServer(mux)
	defer server.Close()

	res, err := http.Get(server.URL + alive)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected %v got %v", http.StatusOK, res.StatusCode)
	}

	res, err = http.Get(server.URL + ready)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected %v got %v", http.StatusOK, res.StatusCode)
	}
}

func TestCustomWithFailure(t *testing.T) {

	probe := func() bool {
		return false
	}
	LivenessCheck(probe)
	ReadinessCheck(probe)

	mux := CreateMux()

	server := httptest.NewServer(mux)
	defer server.Close()

	res, err := http.Get(server.URL + alive)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("expected %v got %v", http.StatusServiceUnavailable, res.StatusCode)
	}

	res, err = http.Get(server.URL + ready)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("expected %v got %v", http.StatusServiceUnavailable, res.StatusCode)
	}
}
