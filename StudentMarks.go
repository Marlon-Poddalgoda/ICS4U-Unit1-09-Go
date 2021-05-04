// This program calculates the average of a class using 2D-arrays.
// Created By: Marlon Poddalgoda
// Created on: 2021-05-04

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func averageValue(slice [3][5]int) int {
	// this function finds the average from a 2D-array

	var totalValue int = 0
	var numOfValues int = 0
	var counter int = 0
	var counter1 int = 0
	var average int

	for counter < len(slice) {
		// loop goes through each value
		for counter1 < len(slice[counter]) {
			// add values to total
			totalValue += slice[counter][counter1]

			// counts up
			numOfValues++
			counter1++
		}
		// add 1 to counter
		counter++
		counter1 = 0
	}

	// calculating average
	average = totalValue / numOfValues

	return (average)
}

func main() {
	// this function generates random marks

	fmt.Println("This program generates 10 random numbers and locates the min/max value")
	fmt.Println("")

	// constants
	const numOfStudents int = 3
	const numOfAssignments = 5
	const randomRange int = 100

	// counter variable
	var counter int = 0
	var counter1 int = 0
	var average int

	// creating array
	var markSlice [numOfStudents][numOfAssignments]int

	// generates random number
	rand1 := rand.NewSource(time.Now().UnixNano())
	random := rand.New(rand1)

	// for loop to place random numbers in array
	for counter < numOfStudents {
		// print each student
		fmt.Print("Student ", counter, "'s marks: ")
		for counter1 < numOfAssignments {
			// generating random numbers
			markSlice[counter][counter1] = random.Intn(randomRange)

			fmt.Print(markSlice[counter][counter1], ", ")

			counter1++
		}
		counter++
		counter1 = 0
		fmt.Println("")
	}

	// call max function
	average = averageValue(markSlice)

	fmt.Println("")
	fmt.Println("The class average is", average)

	// program closes
	fmt.Println("")
	fmt.Println("Done.")
}
