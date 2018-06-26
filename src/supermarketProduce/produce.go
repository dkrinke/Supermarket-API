package supermarketProduce

import ()

// Produce Struct
type Produce struct {
	Name  string //Alphanumeric and case insensitive
	Code  string //Alphanumeric and case insensitive with sixteen characters long, with dashes separating each four character group
	Price string //Number(Represented by String) with up to 2 decimal places
}
