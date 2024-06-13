package main
import "fmt"

// WEIGHT CONVERSION SYSTEM

var userOption uint
var weight float32
func read_weight_conversion() {
	weightConversion()
}
func weightConversion() {
	for {
		fmt.Println("welcome to my weight conversion system")
		fmt.Println("Pick an opion below")
		fmt.Println("1 for Kilogram to Pounds")
		fmt.Println("2 for Pounds to Kilogram")
		fmt.Scan(&userOption)
	
		if userOption == 1 {
			fmt.Println("Enter weight in kilogram")
			fmt.Scan(&weight)
			weight = weight * 2.2
			fmt.Printf("Your weight in pounds is %vlb\n", weight)
		} else if userOption == 2 {
			fmt.Println("Enter weight in Pounds")
			fmt.Scan(&weight)
			weight = weight / 2.2
			fmt.Printf("Your weight in kilograms is %vkg\n", weight)
		}
	}
}