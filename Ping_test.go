package main

import "testing"

func TestPing(t *testing.T) {
	s, err := Subscribe(addr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
	pi := Ping(s[0], 3, "https://cloudflare.com/cdn-cgi/trace")
	t.Log(pi.Link)
	t.Log(pi.Status.Value(0))
	t.Log(pi.Duration)
	t.Log(pi.Err)
}

func TestPingAndSort(t *testing.T) {
	s, err := Subscribe(addr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
	pipList := PingAndSort(s, 3, "https://cloudflare.com/cdn-cgi/trace", 8)
	for _, pip := range pipList {
		pi := *pip
		t.Log(pi.Link)
		t.Log(pi.Status.Value(0))
		t.Log(pi.Duration)
		t.Log(pi.Err)
	}
}
