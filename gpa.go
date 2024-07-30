package main

import (
	"fmt"
)

// Function to calulate average grade
func gradeCalculator(grade int, res map[string]int, subjectCount int, name string) int {
	total := 0
	for _, value := range res {
		total += value

	}

	average := total / subjectCount

	//To display the all  the result including the grades and the average
	fmt.Println("Welcome %v ", name)
	fmt.Println("v,Your Grades Are:", name)
	for id, value := range res {

		fmt.Println(id, value)

	}
	fmt.Println(average)
	return average
}

func main() {

	//Average grade based on different subjects
	var name string
	var subjectCount int
	var grade int
	var subjectName string

	//To enter the name and number of subject they have taken
	fmt.Println("Enter your name!")
	fmt.Scan(&name)
	fmt.Println("How many subjects have you taken?")
	fmt.Scan(&subjectCount)

	//Using map to store all the grades the person enters
	res := make(map[string]int)
	for i := 0; i < subjectCount; i++ {
		fmt.Println("Enter Subject name")
		fmt.Scan(&subjectName)

		//validating the user grade input
		for {
			fmt.Println("Please enter your grade:")
			if _, err := fmt.Scan(&grade); err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}

			if grade > 100 || grade < 0 {
				fmt.Println("Please enter a value between 0 and 100.")
				continue
			}

			break
		}
		res[subjectName] = grade
	}
	gradeCalculator(grade, res, subjectCount, name)

}

//Clean code style
/*Write clear comments
Use meaningful names
Avoid duplicating code
use abstraction
reuse the piece of code
functions short
Use consistent formatting*/
