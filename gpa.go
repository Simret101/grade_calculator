package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Function to calculate average grade
func gradeCalculator(grade int, res map[string]int, subjectCount int, name string) int {
	total := 0
	for _, value := range res {
		total += value
	}

	average := total / subjectCount

	//To display the all the result including the grades and the average
	fmt.Printf("Welcome %s \n", name)
	fmt.Println("Your Grades Are:")
	for id, value := range res {
		fmt.Println(id, value)
	}
	fmt.Println("Average:", average)
	return average
}

func main() {
	//Average grade based on different subjects
	var name string
	var subjectCount int
	var grade int
	var subjectName string

	//To enter the name and number of subject they have taken
	for {
		fmt.Println("Enter your full name:")
		name, _ = readLine()

		// To Validate the name
		if isValidName(name) {
			break
		} else {
			fmt.Println(" Name must contain only letters and cannot be empty.")
			continue
		}
	}

	for {
		fmt.Println("How many subjects have you taken?")
		subjectCountStr, _ := readLine()
		//Atoi:- converts string to number
		subjectCount, _ = strconv.Atoi(subjectCountStr)
		if subjectCount > 0 {
			break
		} else {
			fmt.Println("Subject count must be a positive integer.")
			continue
		}
	}

	//Using map to store all the grades the person enters
	res := make(map[string]int)
	for i := 0; i < subjectCount; i++ {
		for {
			fmt.Println("Enter Subject name:")
			subjectName, _ = readLine()

			// Validate the subject name
			if isValidName(subjectName) {
				break
			} else {
				fmt.Println("Subject name must contain only letters and cannot be empty.")
				continue
			}
		}

		//Validate the user grade input
		for {
			fmt.Println("Please enter your grade:")
			gradeStr, _ := readLine()
			grade, _ = strconv.Atoi(gradeStr)
			if grade >= 0 && grade <= 100 {
				res[subjectName] = grade
				break
			} else {
				fmt.Println("Subject count must be a positive integer.")
				continue
			}

		}
	}
	gradeCalculator(grade, res, subjectCount, name)
}

// Instead of writing the reader over and over I used readLine() function to simplify and reuse in the code
func readLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	ans, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(ans), nil
}

// To validate the person's name
func isValidName(name string) bool {
	// To Check if the name is empty
	if name == "" {
		return false
	}

	// To Check if the name contains only letters and spaces
	//always use _ if you don't want to specify
	for _, r := range name {
		if !unicode.IsLetter(r) && !unicode.IsSpace(r) {
			return false
		}
	}

	return true
}

//Clean code style
/*Write clear comments
Use meaningful names
Avoid duplicating code
use abstraction
reuse the piece of code
functions short
Use consistent formatting*/

//to test a go file
/*cd the file
go mod init some name(usually the project description __this will create you a go.mod )
go test
go run ___to run use the first letter and tab (shortcut)*/
