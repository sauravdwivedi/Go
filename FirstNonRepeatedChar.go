/* Write a method to find out first non repeated character in
 * a given string. Example:
 *
 * Input:
 * string = "abadbcd"
 *
 * Output:
 * result = 'c'
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func firstChar(inString string) string {
	var result string = ""
	for i := 0; i < len(inString); i++ {
		var newString string = inString[:i] + inString[i+1:]
		if !strings.Contains(newString, string(inString[i])) {
			result = string(inString[i])
			break
		}
	}
	return result
}

func main() {
	fmt.Print("Enter string: ")
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var inString string = strings.TrimSpace(scanner.Text())
	fmt.Println("First non repeated character in string:", firstChar(inString))
}
