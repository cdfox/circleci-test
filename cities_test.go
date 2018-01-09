package main

import (
	"log"
	"testing"
)

func TestFindLargestCity(t *testing.T) {
	cities := getCities()
	expected := cities[2]
	got := findLargestCity(cities)
	if got != expected {
		t.Fatalf("got: %v, expected: %v", got, expected)
	}
}

func TestFindLargestCityWithDB(t *testing.T) {
	db, err := getDB()
	if err != nil {
		log.Fatal(err)
	}

	cities := getCities()

	if err := insertCities(db, cities); err != nil {
		log.Fatal(err)
	}

	citiesFromDB, err := readCities(db)
	if err != nil {
		log.Fatal(err)
	}

	expected := cities[2]
	got := findLargestCity(citiesFromDB)

	if got != expected {
		t.Fatalf("got: %v, expected: %v", got, expected)
	}
}
