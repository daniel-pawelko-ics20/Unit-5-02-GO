// Copyright (c) 2021 Daniel Pawelko All rights reserved
//
// Created by: Daniel Pawelko
// Created on: May 2021
// This file contains GO program that returns if integer is positive or negative

package main

import (
	"fmt"
)

// This main function will ask user for dimensions and output calculations
func main() {
	// Defining variables
	var number int

	fmt.Println("Check if number is negative or positive")
	fmt.Println()

	// User Input
	fmt.Println("Please enter Integer:")
	fmt.Print("Integer: ")
	fmt.Scanln(&number)
	fmt.Println()

	// If number is positive
	if(number > 0) {
    fmt.Println(number, "is positive")
  } 
	// If number is negative  
  if(number < 0) {
    fmt.Println(number, "is negative")
  }
	// If number is 0
  if(number == 0) {
    fmt.Println("0 is neither negative nor positive")
  }
}
