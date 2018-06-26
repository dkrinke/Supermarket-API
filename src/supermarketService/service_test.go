package supermarketService

import "testing"
import "supermarketProduce"

/********************************Test Validate Produce**************************/
func TestValidateProducePass(t *testing.T) {
	//Build produce object with valid data
	var produce = supermarketProduce.Produce{Name: "Green Apples", Code: "123A-123A-123A-123A", Price: "$1.20"}
	_, _, resultBoolean := validateProduce(produce) //Validate and save the boolean result
	if resultBoolean != true {                      //Check if validation passed
		t.Errorf("Result was incorrect, got: %t, want: true.", resultBoolean)
	}
}

func TestValidateProduceFailureInvalidName(t *testing.T) {
	//Build produce object with an invalid name
	var produce = supermarketProduce.Produce{Name: "Green @pples", Code: "123A-123A-123A-123A", Price: "$1.20"}
	_, _, resultBoolean := validateProduce(produce) //Validate and save the boolean result
	if resultBoolean != false {                     //Check if validation failed
		t.Errorf("Result was incorrect, got: %t, want: false.", resultBoolean)
	}
}

func TestValidateProduceFailureInvalidCode_1(t *testing.T) {
	//Build produce object with an invalid code
	var produce = supermarketProduce.Produce{Name: "Green Apples", Code: "123@-123A-123A-123A", Price: "$1.20"}
	_, _, resultBoolean := validateProduce(produce) //Validate and save the boolean result
	if resultBoolean != false {                     //Check if validation failed
		t.Errorf("Result was incorrect, got: %t, want: false.", resultBoolean)
	}
}

func TestValidateProduceFailureInvalidCode_2(t *testing.T) {
	//Build produce object with an invalid code
	var produce = supermarketProduce.Produce{Name: "Green Apples", Code: "123A-123A-123A", Price: "$1.20"}
	_, _, resultBoolean := validateProduce(produce) //Validate and save the boolean result
	if resultBoolean != false {                     //Check if validation failed
		t.Errorf("Result was incorrect, got: %t, want: false.", resultBoolean)
	}
}

func TestValidateProduceFailureInvalidCode_3(t *testing.T) {
	//Build produce object with an invalid code
	var produce = supermarketProduce.Produce{Name: "Green Apples", Code: "123A123A123A123A", Price: "$1.20"}
	_, _, resultBoolean := validateProduce(produce) //Validate and save the boolean result
	if resultBoolean != false {                     //Check if validation failed
		t.Errorf("Result was incorrect, got: %t, want: false.", resultBoolean)
	}
}

func TestValidateProduceFailureInvalidPrice_1(t *testing.T) {
	//Build produce object with an invalid price
	var produce = supermarketProduce.Produce{Name: "Green Apples", Code: "123A-123A-123A-123A", Price: "$1.201"}
	_, _, resultBoolean := validateProduce(produce) //Validate and save the boolean result
	if resultBoolean != false {                     //Check if validation failed
		t.Errorf("Result was incorrect, got: %t, want: false.", resultBoolean)
	}
}

func TestValidateProduceFailureInvalidPrice_2(t *testing.T) {
	//Build produce object with an invalid price
	var produce = supermarketProduce.Produce{Name: "Green Apples", Code: "123A-123A-123A-123A", Price: "$1.2"}
	_, _, resultBoolean := validateProduce(produce) //Validate and save the boolean result
	if resultBoolean != false {                     //Check if validation failed
		t.Errorf("Result was incorrect, got: %t, want: false.", resultBoolean)
	}
}

func TestValidateProduceFailureInvalidPrice_3(t *testing.T) {
	//Build produce object with an invalid price
	var produce = supermarketProduce.Produce{Name: "Green Apples", Code: "123A-123A-123A-123A", Price: "$1."}
	_, _, resultBoolean := validateProduce(produce) //Validate and save the boolean result
	if resultBoolean != false {                     //Check if validation failed
		t.Errorf("Result was incorrect, got: %t, want: false.", resultBoolean)
	}
}

func TestValidateProduceFailureInvalidPrice_4(t *testing.T) {
	//Build produce object with an invalid price
	var produce = supermarketProduce.Produce{Name: "Green Apples", Code: "123A-123A-123A-123A", Price: "$1"}
	_, _, resultBoolean := validateProduce(produce) //Validate and save the boolean result
	if resultBoolean != false {                     //Check if validation failed
		t.Errorf("Result was incorrect, got: %t, want: false.", resultBoolean)
	}
}

func TestValidateProduceFailureInvalidPrice_5(t *testing.T) {
	//Build produce object with an invalid price
	var produce = supermarketProduce.Produce{Name: "Green Apples", Code: "123A-123A-123A-123A", Price: "1.20"}
	_, _, resultBoolean := validateProduce(produce) //Validate and save the boolean result
	if resultBoolean != false {                     //Check if validation failed
		t.Errorf("Result was incorrect, got: %t, want: false.", resultBoolean)
	}
}

/********************************Test Validate Names****************************/
func TestValidateNamePass(t *testing.T) {
	result := validateName("Green Apples") //Call validateNames with valid name
	if result != true {                    //Check if validation succeeded
		t.Errorf("Result was incorrect, got: %t, want: true.", result)
	}
}

func TestValidateNameFailure(t *testing.T) {
	result := validateName("Green @pples") //Call validateNames with invalid name
	if result != false {                   //Check if validation failed
		t.Errorf("Result was incorrect, got: %t, want: false.", result)
	}
}

/********************************Test Validate Codes**************************/
func TestValidateCodePass_1(t *testing.T) {
	result := validateCode("123A-123A-123A-123A") //Call validateCode with valid code
	if result != true {                           //Check if validation succeeded
		t.Errorf("Result was incorrect, got: %t, want: true.", result)
	}
}

func TestValidateCodePass_2(t *testing.T) {
	result := validateCode("1234-1234-1234-1234") //Call validateCode with valid code
	if result != true {                           //Check if validation succeeded
		t.Errorf("Result was incorrect, got: %t, want: true.", result)
	}
}

func TestValidateCodePass_3(t *testing.T) {
	result := validateCode("ABCD-ABCD-ABCD-ABCD") //Call validateCode with valid code
	if result != true {                           //Check if validation succeeded
		t.Errorf("Result was incorrect, got: %t, want: true.", result)
	}
}

func TestValidateCodeFailure_1(t *testing.T) {
	result := validateCode("123@-123A-123A-123A") //Call validateCode with invalid code
	if result != false {                          //Check if validation failed
		t.Errorf("Result was incorrect, got: %t, want: false.", result)
	}
}

func TestValidateCodeFailure_2(t *testing.T) {
	result := validateCode("123A-123A-123A-123") //Call validateCode with invalid code
	if result != false {                         //Check if validation failed
		t.Errorf("Result was incorrect, got: %t, want: false.", result)
	}
}

func TestValidateCodeFailure_3(t *testing.T) {
	result := validateCode("123@-123A-123A") //Call validateCode with invalid code
	if result != false {                     //Check if validation failed
		t.Errorf("Result was incorrect, got: %t, want: false.", result)
	}
}

func TestValidateCodeFailure_4(t *testing.T) {
	result := validateCode("123A123A123A123A") //Call validateCode with invalid code
	if result != false {                       //Check if validation failed
		t.Errorf("Result was incorrect, got: %t, want: false.", result)
	}
}

/********************************Test Validate Price**************************/
func TestValidatePricePass_1(t *testing.T) {
	result := validatePrice("$1.21") //Call validatePrice with valid price
	if result != true {              //Check if validation passed
		t.Errorf("Result was incorrect, got: %t, want: true.", result)
	}
}

func TestValidatePriceFailure_1(t *testing.T) {
	// Invalid Price
	result := validatePrice("$1.212") //Call validatePrice with invalid price
	if result != false {              //Check if validation failed
		t.Errorf("Result was incorrect, got: %t, want: false.", result)
	}
}

func TestValidatePriceFailure_2(t *testing.T) {
	result := validatePrice("$1.2") //Call validatePrice with invalid price
	if result != false {            //Check if validation failed
		t.Errorf("Result was incorrect, got: %t, want: false.", result)
	}
}

func TestValidatePriceFailure_3(t *testing.T) {
	result := validatePrice("$1") //Call validatePrice with invalid price
	if result != false {          //Check if validation failed
		t.Errorf("Result was incorrect, got: %t, want: false.", result)
	}
}

func TestValidatePriceFailure_4(t *testing.T) {
	result := validatePrice("1.21") //Call validatePrice with invalid price
	if result != false {            //Check if validation failed
		t.Errorf("Result was incorrect, got: %t, want: false.", result)
	}
}
