package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

type City struct {
	Name       string
	Population int
}

func getCities() []City {
	return []City{{"New York", 8537673}, {"Paris", 2206488}, {"Tokyo", 13617445}}
}

func getDB() (*sql.DB, error) {
	connStr := "user=root dbname=circle_test sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func insertCities(db *sql.DB, cities []City) error {
	var params []string
	var values []interface{}
	for i, city := range cities {
		params = append(params, fmt.Sprintf("($%v, $%v)", i*2+1, i*2+2))
		values = append(values, city.Name)
		values = append(values, city.Population)
	}
	query := fmt.Sprintf("INSERT INTO cities VALUES %s", strings.Join(params, ", "))
	_, err := db.Exec(
		query,
		values...,
	)
	return err
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

func deleteCities(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM cities")
	return err
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

	log.Printf("read cities: %v\n", citiesFromDB)

	largest := findLargestCity(citiesFromDB)

	log.Printf("largest city: %s\n", largest.Name)

	if err := deleteCities(db); err != nil {
		log.Fatal(err)
	}
}
