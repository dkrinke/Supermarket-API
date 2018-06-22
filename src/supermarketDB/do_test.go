package db

import "testing"
import "supermarketProduce"

// Test Add Produce
func TestAddProduce(t *testing.T) {

	var produce = produce.Produce{Name: "Added Produce", Code: "123A-456B-789C-DEFG", Price: "$99.99"}
	var addProduceResult = AddProduce(produce)

	t.Log(addProduceResult.Name)
	t.Log(addProduceResult.Code)
	t.Log(addProduceResult.Price)

	if addProduceResult.Name != produce.Name {
		t.Errorf("Result was incorrect, got: %s, want: %s.", addProduceResult.Name, produce.Name)
	}
	if addProduceResult.Code != produce.Code {
		t.Errorf("Result was incorrect, got: %s, want: %s.", addProduceResult.Code, produce.Code)
	}
	if addProduceResult.Price != produce.Price {
		t.Errorf("Result was incorrect, got: %s, want: %s.", addProduceResult.Price, produce.Price)
	}
}

// Test Delete Produce
func TestDeleteProduce(t *testing.T) {

	firstReadResultBool, _ := Read("123A-456B-789C-DEFG")
	if firstReadResultBool != true {
		t.Errorf("Result was incorrect, got: %t, want: true.", firstReadResultBool)
	}

	deleteResult := DeleteProduce("123A-456B-789C-DEFG")
	if deleteResult != true {
		t.Errorf("Result was incorrect, got: %s, want: true.", deleteResult)
	}

	secondReadResultBool, _ := Read("123A-456B-789C-DEFG")
	if secondReadResultBool != false {
		t.Errorf("Result was incorrect, got: %t, want: false.", secondReadResultBool)
	}
}

// Test Read All
func TestReadAll(t *testing.T) {

	// Test Data
	testNameList := []string{"Lettuce", "Peach", "Green Pepper", "Gala Apple"}
	testCodeList := []string{"A12T-4GH7-QPL9-3N4M", "E5T6-9UI3-TH15-QR88", "YRT6-72AS-K736-L4AR", "TQ4C-VV6T-75ZX-1RMR"}
	testPriceList := []string{"$3.46", "$2.99", "$0.79", "$3.59"}

	result := ReadAll()

	for i, produce := range result {
		if produce.Name != testNameList[i] {
			t.Errorf("Result[%d] was incorrect, got: %s, want: %s.", i, produce.Name, testNameList[i])
		}
		if produce.Code != testCodeList[i] {
			t.Errorf("Result[%d] was incorrect, got: %s, want: %s.", i, produce.Code, testCodeList[i])
		}
		if produce.Price != testPriceList[i] {
			t.Errorf("Result[%d] was incorrect, got: %s, want: %s.", i, produce.Price, testPriceList[i])
		}
	}
}

// Test Lettuce
func TestReadLettuceName(t *testing.T) {
	_, result := Read("A12T-4GH7-QPL9-3N4M")
	if result.Name != "Lettuce" {
		t.Errorf("Result was incorrect, got: %s, want: Lettuce.", result.Name)
	}
}

func TestReadLettuceCode(t *testing.T) {
	_, result := Read("A12T-4GH7-QPL9-3N4M")
	if result.Code != "A12T-4GH7-QPL9-3N4M" {
		t.Errorf("Result was incorrect, got: %s, want: YRT6-72AS-K736-L4AR.", result.Code)
	}
}

func TestReadLettucePrice(t *testing.T) {
	_, result := Read("A12T-4GH7-QPL9-3N4M")
	if result.Price != "$3.46" {
		t.Errorf("Result was incorrect, got: %s, want: 3.46.", result.Price)
	}
}

// Test Peach
func TestReadPeachName(t *testing.T) {
	_, result := Read("E5T6-9UI3-TH15-QR88")
	if result.Name != "Peach" {
		t.Errorf("Result was incorrect, got: %s, want: Peach.", result.Name)
	}
}

func TestReadPeachCode(t *testing.T) {
	_, result := Read("E5T6-9UI3-TH15-QR88")
	if result.Code != "E5T6-9UI3-TH15-QR88" {
		t.Errorf("Result was incorrect, got: %s, want: YRT6-72AS-K736-L4AR.", result.Code)
	}
}

func TestReadPeachPrice(t *testing.T) {
	_, result := Read("E5T6-9UI3-TH15-QR88")
	if result.Price != "$2.99" {
		t.Errorf("Result was incorrect, got: %s, want: 2.99.", result.Price)
	}
}

// Test Green Pepper
func TestReadGreenPepperName(t *testing.T) {
	_, result := Read("YRT6-72AS-K736-L4AR")
	if result.Name != "Green Pepper" {
		t.Errorf("Result was in0.79correct, got: %s, want: Green Pepper.", result.Name)
	}
}

func TestReadGreenPepperCode(t *testing.T) {
	_, result := Read("YRT6-72AS-K736-L4AR")
	if result.Code != "YRT6-72AS-K736-L4AR" {
		t.Errorf("Result was incorrect, got: %s, want: YRT6-72AS-K736-L4AR.", result.Code)
	}
}

func TestReadGreenPepperPrice(t *testing.T) {
	_, result := Read("YRT6-72AS-K736-L4AR")
	if result.Price != "$0.79" {
		t.Errorf("Result was incorrect, got: %s, want: 0.79.", result.Price)
	}
}

// Test Green Apple
func TestReadGalaAppleName(t *testing.T) {
	_, result := Read("TQ4C-VV6T-75ZX-1RMR")
	if result.Name != "Gala Apple" {
		t.Errorf("Result was incorrect, got: %s, want: Gala Apple.", result.Name)
	}
}

func TestReadGalaAppleCode(t *testing.T) {
	_, result := Read("TQ4C-VV6T-75ZX-1RMR")
	if result.Code != "TQ4C-VV6T-75ZX-1RMR" {
		t.Errorf("Result was incorrect, got: %s, want: TQ4C-VV6T-75ZX-1RMR.", result.Code)
	}
}

func TestReadGalaApplePrice(t *testing.T) {
	_, result := Read("TQ4C-VV6T-75ZX-1RMR")
	if result.Price != "$3.59" {
		t.Errorf("Result was incorrect, got: %s, want: 3.59.", result.Price)
	}
}
