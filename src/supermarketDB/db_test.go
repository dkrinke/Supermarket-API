package supermarketDB

import "testing"
import "supermarketProduce"

/********************************Test ReadAll***********************************/
func TestReadAll(t *testing.T) {

	//Expected result data
	expectedNameList := []string{"Lettuce", "Peach", "Green Pepper", "Gala Apple"}
	expectedCodeList := []string{"A12T-4GH7-QPL9-3N4M", "E5T6-9UI3-TH15-QR88", "YRT6-72AS-K736-L4AR", "TQ4C-VV6T-75ZX-1RMR"}
	expectedPriceList := []string{"$3.46", "$2.99", "$0.79", "$3.59"}

	result := ReadAll() //Read all produce from the db

	for i, produce := range result { //Loop through the resulting produce list
		if produce.Name != expectedNameList[i] { //Check if name does not match expected
			t.Errorf("Result[%d] was incorrect, got: %s, want: %s.", i, produce.Name, expectedNameList[i])
		}
		if produce.Code != expectedCodeList[i] { //Check if code does not match expected
			t.Errorf("Result[%d] was incorrect, got: %s, want: %s.", i, produce.Code, expectedCodeList[i])
		}
		if produce.Price != expectedPriceList[i] { //Check if price does not match expected
			t.Errorf("Result[%d] was incorrect, got: %s, want: %s.", i, produce.Price, expectedPriceList[i])
		}
	}
}

/********************************Test Read(Default Data)************************/
//Test Retrieving Lettuce
func TestReadLettuce(t *testing.T) {
	_, result := Read("A12T-4GH7-QPL9-3N4M") //Retrieve Lettuce from db

	if result.Name != "Lettuce" { //Check if name is "Lettuce"
		t.Errorf("Result was incorrect, got: %s, want: Lettuce.", result.Name)
	}
	if result.Code != "A12T-4GH7-QPL9-3N4M" { //Check if resulting code is "A12T-4GH7-QPL9-3N4M"
		t.Errorf("Result was incorrect, got: %s, want: YRT6-72AS-K736-L4AR.", result.Code)
	}
	if result.Price != "$3.46" { //Check if price is "$3.46"
		t.Errorf("Result was incorrect, got: %s, want: $3.46.", result.Price)
	}
}

//Test Retrieving Peach
func TestReadPeach(t *testing.T) {
	_, result := Read("E5T6-9UI3-TH15-QR88") //Retrieve Peach from db

	if result.Name != "Peach" { //Check if name is "Peach"
		t.Errorf("Result was incorrect, got: %s, want: Peach.", result.Name)
	}
	if result.Code != "E5T6-9UI3-TH15-QR88" { //Check if resulting code is "E5T6-9UI3-TH15-QR88"
		t.Errorf("Result was incorrect, got: %s, want: YRT6-72AS-K736-L4AR.", result.Code)
	}
	if result.Price != "$2.99" { //Check if price is "$2.99"
		t.Errorf("Result was incorrect, got: %s, want: $2.99.", result.Price)
	}
}

//Test Retrieving Green Pepper
func TestReadGreenPepper(t *testing.T) {
	_, result := Read("YRT6-72AS-K736-L4AR") //Retrieve Green Pepper from db

	if result.Name != "Green Pepper" { //Check if name is "Green Pepper"
		t.Errorf("Result was in0.79correct, got: %s, want: Green Pepper.", result.Name)
	}
	if result.Code != "YRT6-72AS-K736-L4AR" { //Check if resulting code is "YRT6-72AS-K736-L4AR"
		t.Errorf("Result was incorrect, got: %s, want: YRT6-72AS-K736-L4AR.", result.Code)
	}
	if result.Price != "$0.79" { //Check if price is "$0.79"
		t.Errorf("Result was incorrect, got: %s, want: $0.79.", result.Price)
	}
}

//Test Retrieving Green Apple
func TestReadGalaApple(t *testing.T) {
	_, result := Read("TQ4C-VV6T-75ZX-1RMR") //Retrieve Gala Apple from db

	if result.Name != "Gala Apple" { //Check if name is "Gala Apple"
		t.Errorf("Result was incorrect, got: %s, want: Gala Apple.", result.Name)
	}
	if result.Code != "TQ4C-VV6T-75ZX-1RMR" { //Check if resulting code is "TQ4C-VV6T-75ZX-1RMR"
		t.Errorf("Result was incorrect, got: %s, want: TQ4C-VV6T-75ZX-1RMR.", result.Code)
	}
	if result.Price != "$3.59" { //Check if price is "$3.59"
		t.Errorf("Result was incorrect, got: %s, want: $3.59.", result.Price)
	}
}

/********************************Test AddProduce********************************/
//Test Successfully adding produce
func TestAddProduce(t *testing.T) {

	//Build produce object with valid produce data
	var produce = supermarketProduce.Produce{Name: "Golden Apple", Code: "123A-456B-789C-DEFG", Price: "$99.99"}
	_, addProduceResult := AddProduce(produce) //Add produce and save the returned produce object

	if addProduceResult.Name != produce.Name { //Check if name matched the expected name
		t.Errorf("Result was incorrect, got: %s, want: %s.", addProduceResult.Name, produce.Name)
	}
	if addProduceResult.Code != produce.Code { //Check if code matched the expected code
		t.Errorf("Result was incorrect, got: %s, want: %s.", addProduceResult.Code, produce.Code)
	}
	if addProduceResult.Price != produce.Price { //Check if price matched the expected price
		t.Errorf("Result was incorrect, got: %s, want: %s.", addProduceResult.Price, produce.Price)
	}
}

//Test Failure in adding Produce
//The code is not unique
func TestAddProduceFailureNonUniqueCode(t *testing.T) {

	//Build produce object with valid produce data, but with a non-unique code
	var produce = supermarketProduce.Produce{Name: "Golden Apple", Code: "123A-456B-789C-DEFG", Price: "$99.99"}
	addProduceBoolean, _ := AddProduce(produce) //Add produce and save the returned boolean indicating Success/Failure

	if addProduceBoolean != false { //The boolean should be false indicating that the produce was rejected
		t.Errorf("Result was incorrect, got: %t, want: false.", addProduceBoolean)
	}
}

/********************************Test DeleteProduce*****************************/
//Test Successfully deleting produce
func TestDeleteProduceSuccess(t *testing.T) {

	firstReadResultBool, _ := Read("123A-456B-789C-DEFG") //Retreive produce from db
	if firstReadResultBool != true {                      //Confirm produce exists in db before test
		t.Errorf("Result was incorrect, got: %t, want: true.", firstReadResultBool)
	}

	deleteResultBool := DeleteProduce("123A-456B-789C-DEFG") //Delete produce and save the resulting boolean
	if deleteResultBool != true {                            //Check to see if the DeleteProduce function reported success
		t.Errorf("Result was incorrect, got: %t, want: true.", deleteResultBool)
	}

	secondReadResultBool, _ := Read("123A-456B-789C-DEFG") //Retreive produce from db
	if secondReadResultBool != false {                     //Confirm produce no longer exists in db
		t.Errorf("Result was incorrect, got: %t, want: false.", secondReadResultBool)
	}
}

//Test Failure in deleting Produce
//The code is not found in the db
func TestDeleteProduceFailure(t *testing.T) {

	deleteResult := DeleteProduce("123A-456B-789C-DEFG") //Delete produce and save the resulting boolean
	if deleteResult != false {                           //Check to see if the DeleteProduce function reported failure
		t.Errorf("Result was incorrect, got: %t, want: false.", deleteResult)
	}
}
