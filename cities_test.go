package main

import "testing"

func TestFindLargestCity(t *testing.T) {
	cities := []City{{"New York", 8537673}, {"Paris", 2206488}, {"Tokyo", 13617445}}
	expected := cities[2]
	got := findLargestCity(cities)
	if got != expected {
		t.Fatalf("got: %v, expected: %v", got, expected)
	}
}
