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
			Price: "$3.46",
		},
		produce.Produce{
			Name:  "Peach",
			Code:  "E5T6-9UI3-TH15-QR88",
			Price: "$2.99",
		},
		produce.Produce{
			Name:  "Green Pepper",
			Code:  "YRT6-72AS-K736-L4AR",
			Price: "$0.79",
		},
		produce.Produce{
			Name:  "Gala Apple",
			Code:  "TQ4C-VV6T-75ZX-1RMR",
			Price: "$3.59",
		},
	}
)

// Retrieve produce with matching code
func Read(code string) (bool, produce.Produce) {
	var locatedProduce produce.Produce
	var located = false
	var wg sync.WaitGroup

	// Start async task to read from db
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, produce := range database {
			if produce.Code == code {
				locatedProduce = produce
				located = true
			}
		}
	}()

	wg.Wait() // Wait for read from the db to complete

	return located, locatedProduce
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

func AddProduce(produce produce.Produce) (bool, produce.Produce) {

	var added = false
	var wg sync.WaitGroup

	// Start async task to write to db
	wg.Add(1)
	go func() {
		defer wg.Done()
		if isCodeUnique(produce.Code) {
			database = append(database, produce)
			added = true
		}
	}()

	wg.Wait() // Wait for read from the db to complete

	return added, produce
}

func isCodeUnique(code string) bool {
	var unique = true

	for _, produce := range database {
		if produce.Code == code {
			unique = false
			break
		}
	}
	return unique
}

func DeleteProduce(code string) bool {

	var wg sync.WaitGroup
	var produceFound bool = false

	// Start async task to delete from db
	wg.Add(1)
	go func() {
		defer wg.Done()
		var index int

		for i, produce := range database {
			if produce.Code == code {
				produceFound = true
				index = i
			}
		}

		if produceFound {
			copy(database[index:], database[index+1:]) // Shift a[i+1:] left one index
			database = database[:len(database)-1]      // Truncate slice
		}
	}()

	wg.Wait() // Wait for read from the db to complete

	return produceFound
}
