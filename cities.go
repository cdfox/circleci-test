package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type City struct {
	Name       string
	Population int
}

func readCities(db *sql.DB) ([]City, error) {
	var cities []City
	rows, err := db.Query("SELECT * FROM cities")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var city City
		if err := rows.Scan(&city.Name, &city.Population); err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cities, nil
}

func findLargestCity(cities []City) City {
	var largest City
	for _, city := range cities {
		if city.Population > largest.Population {
			largest = city
		}
	}
	return largest
}

func main() {
	connStr := "user=root dbname=circle_test sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	cities, err := readCities(db)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("read cities: %v\n", cities)

	largest := findLargestCity(cities)

	log.Printf("largest city: %s\n", largest.Name)
}
