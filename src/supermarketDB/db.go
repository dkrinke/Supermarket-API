package db

import (
	"supermarketProduce"
	"sync"
)

var (
  // db instantiated with default data
	database = []produce.Produce{
		produce.Produce{
			Name:  "Lettuce",
			Code:  "A12T-4GH7-QPL9-3N4M",
			Price: 3.46,
		},
		produce.Produce{
			Name:  "Peach",
			Code:  "E5T6-9UI3-TH15-QR88",
			Price: 2.99,
		},
		produce.Produce{
			Name:  "Green Pepper",
			Code:  "YRT6-72AS-K736-L4AR",
			Price: 0.79,
		},
		produce.Produce{
			Name:  "Gala Apple",
			Code:  "TQ4C-VV6T-75ZX-1RMR",
			Price: 3.59,
		},
	}
)

// Retrieve produce with matching code
func Read(code string) produce.Produce {
	var locatedProduce produce.Produce
	var wg sync.WaitGroup

  // Start async task to read from db
  wg.Add(1)
	go func() {
		defer wg.Done()
		for _, produce := range database {
			if produce.Code == code {
				locatedProduce = produce
			}
		}
	}()

	wg.Wait() // Wait for read from the db to complete

	return locatedProduce
}

// Retrieve all produce
func ReadAll() []produce.Produce {
	var wg sync.WaitGroup
	var locatedProduce []produce.Produce

  // Start async task to read from db
  wg.Add(1)
	go func() {
		defer wg.Done()
		locatedProduce = database
	}()

	wg.Wait() // Wait for read from the db to complete

	return locatedProduce
}
