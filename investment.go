package main 

import (
	"math"
	"fmt"
)

func main() {
	var principalAmount float64
	var expectedReturnRate float64
	var years float64
	var futureRealValue float64

	fmt.Print("what's your Principal: ")
	fmt.Scan(&principalAmount)
	fmt.Print("The company's return rate is: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("For how many years do you plan on investing with us: ")
	fmt.Scan(&years)
	var futureValue = principalAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue = futureValue / 
	fmt.Printf("Your future value is %.2f \n", math.Round(futureValue))
}