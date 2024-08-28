package main

import (
	"fmt"
	"os"
)

// Bill excluding the taxes
var subTotalBillExcludetaxes float64

// Creating a map of type string - int , string the food , int the number of plates ordered
var customerOrder = make(map[string]uint, 0)

func main() {
	var customerName string

	fmt.Println("Hey, Welcome to the food-world !")
	fmt.Println("Can you help me with your name ?")

	_, err := fmt.Scanf("%s", &customerName)

	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	greet(customerName)
	orderItems()
	displayGeneratingBill() //just displays that "generating bill" in a fancy manner.
	generateBill()
	modifyOrder()
	printFinalBill()
	sayTata(customerName)

}
