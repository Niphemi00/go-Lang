package main
import "fmt"
//PROFIT CALCULATOR APP

func main() {
	fmt.Println("Welcome to the Profit Calculator app")
	calculateProfit()
}
func outputText(text string) {
	fmt.Print(text)
}

func calculateEarningsBeforeTax() float64 {
		// Calculate earnings before tax
		var revenue float64
		var expense float64
			// Using revenue, expenses
		outputText("Enter revenue: ")
		fmt.Scan(&revenue)
		outputText("Enter expense: ")
		fmt.Scan(&expense)
		var earningsBeforeTax float64 = revenue - expense
		return earningsBeforeTax
}
func calculateTax() float64 {
		// Calculate tax
		var taxRate float64	
		outputText("Enter tax rate: ")
		fmt.Scan(&taxRate)
		earningsBeforeTax := calculateEarningsBeforeTax()
		fmt.Printf("Earnings before tax: $%.2f\n", earningsBeforeTax)
		var tax float64 = earningsBeforeTax * float64(taxRate/100)
		return tax
}
func calculateProfit() {
		// Calculate Profit
		earningsBeforeTax := calculateEarningsBeforeTax() 
		tax := calculateTax()
		fmt.Printf("Value to be deducted(tax): $%.2f\n", tax)
		var profit float64 = earningsBeforeTax - tax
		fmt.Printf("Your Profit equals $%.2f\n", profit)
		var ratio float64 = earningsBeforeTax / tax
		fmt.Printf("Your earning ratio equals %.2f\n", ratio)
}