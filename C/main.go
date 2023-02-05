package main

import (
	"fmt"
	"log"
	"strings"
)

func validString(newString string, oldString string) bool {
	newString = strings.ReplaceAll(newString, " ", "")
	return newString == oldString
}
func printValues(autoNumbers []string) {

	for _, value := range autoNumbers {
		checkMe := ""
		newString := ""
		newNum := ""
		for _, symb := range value {
			if symb >= 65 && symb <= 90 {
				checkMe += "s"
				newNum += string(symb)
			}
			if symb >= 48 && symb <= 57 {
				checkMe += "n"
				newNum += string(symb)
			}
			if checkMe == "snnss" || checkMe == "snss" {
				newString += newNum + " "
				checkMe = ""
				newNum = ""
			}
		}
		if validString(newString, value) {
			fmt.Println(newString)
		} else {
			fmt.Println("-")
		}
	}

}

func main() {
	var lines int
	var autoNumbers []string
	if _, err := fmt.Scanln(&lines); err != nil {
		log.Printf("failed to scan num of lines")
		return
	}
	if lines < 1 || lines > 1000 {
		log.Printf("wrong number of lines")
		return
	}
	for i := 0; i < lines; i++ {
		var newString string
		fmt.Scanln(&newString)
		if (len(newString) > 50) || (len(newString) < 1) {
			log.Printf("line should be from 1 to 50 symbols")
			return
		}
		autoNumbers = append(autoNumbers, newString)

	}
	printValues(autoNumbers)
}
