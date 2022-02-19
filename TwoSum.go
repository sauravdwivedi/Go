/* Write a method that checks if there is at least one pair of
 * numbers which sum equals target. arr=[1, 3, 4] and target=5
 * result is true because the pair (1,4) sums to five.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func twoSum(listOfNumbs []int, targetNum int) string {
	var result bool = false
	for _, p := range listOfNumbs {
		for _, q := range listOfNumbs {
			if (p != q) && (p+q == targetNum) {
				result = true
			}
		}
	}
	if result == true {
		return "At least one pair has sum equal to target!"
	} else {
		return "No pair has sum equal to target!"
	}
}

func main() {
	fmt.Println("Enter list of integers (space separated e.g. '1 2 3'): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var listOfNumbsTemp []string = strings.Split(strings.
		TrimSpace(scanner.Text()), " ")
	var listOfNumbs []int
	for i := 0; i < len(listOfNumbsTemp); i++ {
		element, err := strconv.ParseInt(listOfNumbsTemp[i], 10, 64)
		if err == nil {
			listOfNumbs = append(listOfNumbs, int(element))
		} else {
			panic(err)
		}
	}
	fmt.Println("Enter target integer: ")
	var targetNum int
	fmt.Scan(&targetNum)
	fmt.Printf("%s", twoSum(listOfNumbs, targetNum))
}
