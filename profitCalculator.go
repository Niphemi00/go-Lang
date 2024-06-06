package main
import "fmt"
//PROFIT CALCULATOR APP
func main() {
	fmt.Println("Welcome to the Profit Calculator app")
	// Using revenue, expenses and tax rate
	var revenue float64
	var expense float64
	var taxRate float32
	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter expense: ")
	fmt.Scan(&expense)
	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)
	// Calculate earnings before tax
	var earningsBeforeTax float64 = revenue - expense
	fmt.Printf("Earnings before tax: $%.2f\n", earningsBeforeTax)
	// Calculate tax
	var tax float64 = earningsBeforeTax * float64(taxRate/100)
	fmt.Printf("Value to be deducted(tax): $%.2f\n", tax)
	// Calculate Profit 
	var profit float64 = earningsBeforeTax - tax
	fmt.Printf("Your Profit equals $%.2f\n", profit)
	var ratio float64 = earningsBeforeTax / tax
	fmt.Printf("Your earning ratio equals %.2f\n", ratio)

}