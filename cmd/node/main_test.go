package main

import "testing"

func TestBanner(t *testing.T) {
	want := "handmade-blockchain node: ready"

	if got := banner(); got != want {
		t.Fatalf("banner() = %q, want %q", got, want)
	}
}
