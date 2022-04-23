/** A staircase of size (say n = 5) is combination of white spaces and #:
 *
 *      #
 *     ##
 *    ###
 *   ####
 *  #####
 *
 *  where number of #'s at bottom is equal to n. Whrite a method to print
 *  a staircase for a given integer n as input.
 */
 

package main


import (
    "fmt"
)


func stairCase(stairCaseSize int) string {
    var stairCase string = ""
    
    for i := 1; i < stairCaseSize + 1; i++ {
        for j := 0; j < stairCaseSize - i; j++ {
            stairCase += " "
        }
        
        for k := 0; k < i; k++ {
            stairCase += "#"
        }
        
        stairCase += "\n"
    }
    
    return stairCase
}


func main() {
    fmt.Print("Size of staircase: ")
    var stairCaseSize int
    _, err := fmt.Scan(&stairCaseSize)
    
    if err == nil {
        var result string = stairCase(stairCaseSize)
        fmt.Printf("The staircase of size %d! \n", stairCaseSize)
        fmt.Print(result)
    } else {
        panic(err)
    }

}
