package db

import (
	"supermarketProduce" //Provides the Produce Struct
	"sync"               //Provides basic synchronization primitives such as mutual exclusion locks
)

var (
	//db instantiated with default data
	database = []produce.Produce{
		produce.Produce{ //Required Lettuce
			Name:  "Lettuce",
			Code:  "A12T-4GH7-QPL9-3N4M",
			Price: "$3.46",
		},
		produce.Produce{ //Required Peach
			Name:  "Peach",
			Code:  "E5T6-9UI3-TH15-QR88",
			Price: "$2.99",
		},
		produce.Produce{ //Required Green Pepper
			Name:  "Green Pepper",
			Code:  "YRT6-72AS-K736-L4AR",
			Price: "$0.79",
		},
		produce.Produce{ //Required Gala Apple
			Name:  "Gala Apple",
			Code:  "TQ4C-VV6T-75ZX-1RMR",
			Price: "$3.59",
		},
	}
)

//Retrieve produce with matching code
func Read(code string) (bool, produce.Produce) {
	var locatedProduce produce.Produce //Create Produce Object
	var located = false                //Initialize located to false (Indicates if Produce was found)
	var wg sync.WaitGroup              //Create WaitGroup

	//Start async task to read from db
	wg.Add(1)   //Add one to WaitGroup
	go func() { //Start GoRoutine
		defer wg.Done()                    //Execute at the end of this Routine
		for _, produce := range database { //Loop through database
			if produce.Code == code { //If Produce code matches provided code
				locatedProduce = produce //Set locatedProduce to the current produce
				located = true           //Set located to true
			}
		}
	}()

	wg.Wait() // Wait for read from the db to complete

	return located, locatedProduce
}

//Retrieve all produce
func ReadAll() []produce.Produce {
	var locatedProduce []produce.Produce //Create Produce Object List
	var wg sync.WaitGroup                //Create WaitGroup

	//Start async task to read from db
	wg.Add(1)   //Add one to WaitGroup
	go func() { //Start GoRoutine
		defer wg.Done()           //Execute at the end of this Routine
		locatedProduce = database //Save Produce from database to the produce list
	}()

	wg.Wait() // Wait for read from the db to complete

	return locatedProduce
}

func ResetData() {
	var wg sync.WaitGroup //Create WaitGroup

	//Start async task to read from db
	wg.Add(1)   //Add one to WaitGroup
	go func() { //Start GoRoutine
		defer wg.Done() //Execute at the end of this Routine
		database = []produce.Produce{
			produce.Produce{ //Required Lettuce
				Name:  "Lettuce",
				Code:  "A12T-4GH7-QPL9-3N4M",
				Price: "$3.46",
			},
			produce.Produce{ //Required Peach
				Name:  "Peach",
				Code:  "E5T6-9UI3-TH15-QR88",
				Price: "$2.99",
			},
			produce.Produce{ //Required Green Pepper
				Name:  "Green Pepper",
				Code:  "YRT6-72AS-K736-L4AR",
				Price: "$0.79",
			},
			produce.Produce{ //Required Gala Apple
				Name:  "Gala Apple",
				Code:  "TQ4C-VV6T-75ZX-1RMR",
				Price: "$3.59",
			},
		}
	}()

	wg.Wait() // Wait for reset of the db to complete
}

//Add produce to the db
func AddProduce(produce produce.Produce) (bool, produce.Produce) {
	var added = false     //Initialize added to false (Indicates if Produce was added successfully)
	var wg sync.WaitGroup //Create WaitGroup

	//Start async task to write to db
	wg.Add(1)   //Add one to WaitGroup
	go func() { //Start GoRoutine
		defer wg.Done()                 //Execute at the end of this Routine
		if isCodeUnique(produce.Code) { //If code provided does not match any in the DB
			database = append(database, produce) //Add the produce to the DB
			added = true                         //Indicate that the produce was added successfully
		}
	}()

	wg.Wait() // Wait for add to the db to complete

	return added, produce
}

//Delete produce from the db
func DeleteProduce(code string) bool {
	var produceFound = false //Initialize produceFound to false (Indicates if Produce was found)
	var wg sync.WaitGroup    //Create WaitGroup

	// Start async task to delete from db
	wg.Add(1)   //Add one to WaitGroup
	go func() { //Start GoRoutine
		defer wg.Done()                    //Execute at the end of this Routine
		var index int                      //Create index to track location of produce that matches the code
		for i, produce := range database { //Loop through database
			if produce.Code == code { //If provided code matches one found in the DB
				produceFound = true //Set produceFound to true
				index = i           //Save index of the produce found
				break               //End search
			}
		}

		if produceFound { //If produce was found, remove it
			copy(database[index:], database[index+1:]) // Shift database[index+1:] left one index
			database = database[:len(database)-1]      // Truncate slice
		}
	}()

	wg.Wait() // Wait for delete from the db to complete

	return produceFound
}

//Check if the provided code matches any in the DB
//Does not start it's own GoRoutine since it is called from a db GoRoutine
func isCodeUnique(code string) bool {
	var unique = true //Initialize unique to true(Set to false if code found in DB)

	for _, produce := range database { //Loop through database
		if produce.Code == code { //If provided code matches one found in the DB
			unique = false //Set unique to false
			break          //End search
		}
	}
	return unique
}
