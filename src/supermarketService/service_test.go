package service

import "testing"
import "supermarketProduce"

// Test Validate Produce
func TestValidateProducePass(t *testing.T) {
  var produce = produce.Produce{Name: "Green Apples", Code: "123A-123A-123A-123A", Price: "$1.20"}
	_, _, resultBoolean := validateProduce(produce)
	if resultBoolean != true {
		t.Errorf("Result was incorrect, got: %t, want: true.", resultBoolean)
	}
}

func TestValidateProduceFailure(t *testing.T) {
  var produce = produce.Produce{Name: "Green Apples!", Code: "123A-123A-123A-123A", Price: "$1.20"}
	_, _, resultBoolean := validateProduce(produce)
	if resultBoolean != false {
		t.Errorf("Result was incorrect, got: %t, want: true.", resultBoolean)
	}
}

// Test Validation of Nmaes
func TestValidateNamePass(t *testing.T) {
	result := validateName("Green Apples")
	if result != true {
		t.Errorf("Result was incorrect, got: %t, want: true.", result)
	}
}

func TestValidateNameFailure(t *testing.T) {
	result := validateName("Green @pples")
	if result != false {
		t.Errorf("Result was incorrect, got: %t, want: false.", result)
	}
}

// Test Validation of Codes
func TestValidateCodePass_1(t *testing.T) {
	result := validateCode("123A-123A-123A-123A")
	if result != true {
		t.Errorf("Result was incorrect, got: %t, want: true.", result)
	}
}

func TestValidateCodePass_2(t *testing.T) {
	result := validateCode("1234-1234-1234-1234")
	if result != true {
		t.Errorf("Result was incorrect, got: %t, want: true.", result)
	}
}

func TestValidateCodePass_3(t *testing.T) {
	result := validateCode("ABCD-ABCD-ABCD-ABCD")
	if result != true {
		t.Errorf("Result was incorrect, got: %t, want: true.", result)
	}
}

func TestValidateCodeFailure_1(t *testing.T) {
	result := validateCode("123A-123A-123A-123")
	if result != false {
		t.Errorf("Result was incorrect, got: %t, want: false.", result)
	}
}

func TestValidateCodeFailure_2(t *testing.T) {
	result := validateCode("123A-123A-123A-123!")
	if result != false {
		t.Errorf("Result was incorrect, got: %t, want: false.", result)
	}
}

func TestValidateCodeFailure_3(t *testing.T) {
	result := validateCode("123A123A123A123A")
	if result != false {
		t.Errorf("Result was incorrect, got: %t, want: false.", result)
	}
}

// Test Validation of Prices
func TestValidatePricePass_1(t *testing.T) {
	result := validatePrice("$1.21")
	if result != true {
		t.Errorf("Result was incorrect, got: %t, want: true.", result)
	}
}

func TestValidatePriceFailure_1(t *testing.T) {
  // Invalid Price
	result := validatePrice("$1.212")
	if result != false {
		t.Errorf("Result was incorrect, got: %t, want: false.", result)
	}
}

func TestValidatePriceFailure_2(t *testing.T) {
	result := validatePrice("$1.2")
	if result != false {
		t.Errorf("Result was incorrect, got: %t, want: false.", result)
	}
}


func TestValidatePriceFailure_3(t *testing.T) {
	result := validatePrice("$1")
	if result != false {
		t.Errorf("Result was incorrect, got: %t, want: false.", result)
	}
}
