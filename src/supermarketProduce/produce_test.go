package supermarketProduce

import "testing"

const (
	name  = "Green Apples"        //Name test data
	code  = "123A-123A-123A-123A" //Code test data
	price = "$1.20"               //Price test data
)

//Test Validate Produce
func TestProduceStruct(t *testing.T) {
	//Build produce object
	var produce = Produce{Name: name, Code: code, Price: price}
	if produce.Name != name { //Check if name does not match expected
		t.Errorf("Result was incorrect, got: %s, want: %s.", produce.Name, name)
	}
	if produce.Code != code { //Check if code does not match expected
		t.Errorf("Result was incorrect, got: %s, want: %s.", produce.Code, code)
	}
	if produce.Price != price { //Check if price does not match expected
		t.Errorf("Result was incorrect, got: %s, want: %s.", produce.Price, price)
	}
}
