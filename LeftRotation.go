/* A left rotation operation on an array shifts each of the array's
 * elements 1 unit to the left. For example, if 2 left rotations are
 * performed on array [1,2,3,4,5], then the array would become
 * [3,4,5,1,2]. Note that the lowest index item moves to the highest
 * index in a rotation. This is called a circular array.
 *
 * Given an array a of n integers and a number, d, perform d left
 * rotations on the array. Return the updated array to be printed
 * as a single line of space-separated integers.
 */

package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func leftRot(intArray []int, numRot int) []int {
    for i := 0; i < numRot; i++ {
        var rotArray []int
        for j := 0; j < len(intArray)-1; j++ {
            rotArray = append(rotArray, int(intArray[j+1]))
        }
        rotArray = append(rotArray, int(intArray[0]))
        intArray = rotArray
    }
    return intArray
}

func main() {
    fmt.Println("Enter array of integers (space separated e.g. '1 2 3'): ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    var intArrayTemp []string = strings.Split(strings.
        TrimSpace(scanner.Text()), " ")
    var intArray []int
    for i := 0; i < len(intArrayTemp); i++ {
        element, err := strconv.ParseInt(intArrayTemp[i], 10, 64)
        if err == nil {
            intArray = append(intArray, int(element))
        } else {
            panic(err)
        }
    }
    fmt.Println("Enter no. of rotations: ")
    var numRot int
    fmt.Scan(&numRot)
    fmt.Printf("Array after %d left rotations: %d", numRot, leftRot(intArray, numRot))
}
