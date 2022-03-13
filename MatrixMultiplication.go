/**Write a method to read two matrices from StdIn and print
 * their product matrix.
 */
 
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func matMult(matA [][]int, matB [][]int, rowsA int, colsA int, colsB int) [][]int {
    var prodMat [][]int
    for i := 0; i < rowsA; i++ {
        var rowTemp []int
        for j := 0; j < colsB ; j++ {
            var prodMatIJ int = 0
            for k := 0; k < colsA; k++ {
                prodMatIJ += matA[i][k] * matB[k][j]
            }
            rowTemp = append(rowTemp, prodMatIJ)
        }
        prodMat = append(prodMat, rowTemp)
    }
    return prodMat
}

func main() {
    fmt.Println("Matrix A rows: ")
    var rowsA int
    fmt.Scan(&rowsA)
    fmt.Println("Matrix A cols: ")
    var colsA int
    fmt.Scan(&colsA)
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Space separated Matrix A entries (e.g. '1 2 3'): ")
    scanner.Scan()
    var matARaw []string = strings.Split(strings.TrimSpace(scanner.Text()), " ")
    var matA [][]int
    for i := 0; i < rowsA; i++ {
        var rowTemp []int
        for j := 0; j < colsA; j++ {
            element, err := strconv.ParseInt(matARaw[i*colsA + j], 10, 64)
            if err == nil {
                rowTemp = append(rowTemp, int(element))
            } else {
                panic(err)
            }
        }
        matA = append(matA, rowTemp)
    }
    fmt.Println("Matrix B rows: ")
    var rowsB int
    fmt.Scan(&rowsB)
    fmt.Println("Matrix B cols: ")
    var colsB int
    fmt.Scan(&colsB)
    fmt.Println("Space separated Matrix B entries (e.g. '1 2 3'): ")
    scanner.Scan()
    var matBRaw []string = strings.Split(strings.TrimSpace(scanner.Text()), " ")
    var matB [][]int
    for i := 0; i < rowsB; i++ {
        var rowTemp []int
        for j := 0; j < colsB; j++ {
            element, err := strconv.ParseInt(matBRaw[i*colsB + j], 10, 64)
            if err == nil {
                rowTemp = append(rowTemp, int(element))
            } else {
                panic(err)
            }
        }
        matB = append(matB, rowTemp)
    }
    if colsA == rowsB {
        fmt.Println("The product Matrix is: ")
        var prodMat = matMult(matA, matB, rowsA, colsA, colsB)
        for i := 0; i < rowsA; i++ {
            for j := 0; j < colsB; j++ {
                fmt.Printf("%d ", prodMat[i][j])
            }
            fmt.Println("")
        }
    } else {
        fmt.Println("Matrices can't be multiplied!")
    }
}
