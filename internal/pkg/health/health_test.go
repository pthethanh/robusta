package health

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReadiness(t *testing.T) {
	ts := httptest.NewServer(Readiness())
	defer ts.Close()
	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("NotReady Response.StatusCode = %d; expected %d;", resp.StatusCode, http.StatusServiceUnavailable)
	}
	Ready()
	resp, err = http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Ready Response.StatusCode = %d; expected %d;", resp.StatusCode, http.StatusOK)
	}
}

func TestLivenessBasic(t *testing.T) {
	ts := httptest.NewServer(Liveness())
	defer ts.Close()
	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response.StatusCode = %d; expected %d;", resp.StatusCode, http.StatusOK)
	}
}

func TestLivenessChecks(t *testing.T) {
	var err1 error
	var err2 error
	check1 := func(context.Context) error {
		return err1
	}
	check2 := func(context.Context) error {
		return err2
	}
	ts := httptest.NewServer(Liveness(check1, check2))
	defer ts.Close()
	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response.StatusCode = %d; expected %d;", resp.StatusCode, http.StatusOK)
	}
	err1 = errors.New("health check 1 fail")
	resp, err = http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("Response.StatusCode = %d; expected %d;", resp.StatusCode, http.StatusServiceUnavailable)
	}
}
