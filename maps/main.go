// Map are like dictionary in python

package main

import "fmt"

// Product struct containing short name
type Product struct {
	shortName   string
	Description string
}

// ReferenceEntity description
type ReferenceEntity struct {
	description string
}

// define p which is a map and maps a string to product
// e.g. Swaps to SW
var p map[string]Product

//method 1 to inituialize
var r = map[string]ReferenceEntity{
	"1": {"google"},
	"2": {"DTCC"},
}

func main() {
	//method 2 to initialize
	p = make(map[string]Product)
	// Insert a new entry
	p["Swaps"] = Product{"SW", "Swaps"}
	//Insert a new entry
	p["Credit Derivatives"] = Product{"CDS", "Credit Default Swaps"}
	//insert a new entry
	p["Equity"] = Product{"EQ", "Equity"}
	//enter twice, it keps it still unique
	p["Equity"] = Product{"EQ", "Equity"}
	//new entry
	p["Interest Rate"] = Product{"IR", "Interest rate"}
	fmt.Println(p)
	fmt.Println(p["SWAPS"]) //case sensitive
	fmt.Println(p["Swaps"])
	fmt.Println(r)
}
