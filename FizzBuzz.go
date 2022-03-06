/* Write a program to print "Fizz" if a number is divisible by 3, print
 * "Buzz" if a number is divisible by 5, and print "FizzBuzz" if a number
 * is divisible by both 3 and 5 for numbers in a given range. If none of
 * these conditions are met, print the number.
 */
 
package main
 
import (
    "fmt"
)
 
func fizzBuzz(targetNum int) {
    for i := 1; i < targetNum; i++ {
        if (i % 3 == 0) && (i % 5 == 0) {
            fmt.Println("FizzBuzz")
        } else if i % 3 == 0 {
            fmt.Println("Fizz")
        } else if i % 5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
    }
}

func main() {
    fmt.Print("Enter range of numbers: ")
    var targetNum int
    _, err := fmt.Scan(&targetNum)
    if err == nil {
        fizzBuzz(targetNum)
    } else {
        fmt.Println(err)
    }
}
