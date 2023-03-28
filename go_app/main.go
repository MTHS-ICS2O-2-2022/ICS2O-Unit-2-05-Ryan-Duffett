// Created by: Ryan Duffett
// Created on: Sep 2020
//
// This program calculates tax and salary

package main

import (
	"fmt"
)

func main() {
	// This function finds the area and perimeter of a rectangle
	var pay int
	var hours_worked int
	var net_payment int
	var taxes int
	// This function displays currency
	//accountingFormater := accounting.Accounting{Symbol: "$", Precision: 2}
	//fmt.Println(accountingFormater.FormatMoney(123456789.213123)) 

	// input
	fmt.Println("This program finds the amount of taxes you have to pay.")
	fmt.Println()
	fmt.Print("Enter your pay (per hour): ")
	fmt.Scanln(&pay)
	fmt.Print("Enter your hours worked: ")
	fmt.Scanln(&hours_worked)

	// process
	net_payment = (hours_worked * pay) * (1.00 - 0.18)
	taxes = (hours_worked * pay) *  0.18

	// output
	fmt.Println()
	fmt.Println("Your total pay is:", net_payment)
	fmt.Println("The taxes deducted are:", taxes)
	fmt.Println("\nDone.")
}