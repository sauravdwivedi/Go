/**Write a method to read two integer matrices from StdIn and print
 * their product matrix.
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type matrix struct {
	name   string
	rows   int
	cols   int
	matrix [][]int
}

func (m *matrix) setMatrix() *matrix {
	fmt.Println(fmt.Sprintf("Matrix %s rows: ", m.name))
	var rows int
	_, errOne := fmt.Scan(&rows)
	if errOne != nil {
		log.Print("Can only accept integers. Please try again!")
		return m.setMatrix()
	}
	fmt.Println(fmt.Sprintf("Matrix %s cols: ", m.name))
	var cols int
	_, errTwo := fmt.Scan(&cols)
	if errTwo != nil {
		log.Print("Can only accept integers. Please try again!")
		return m.setMatrix()
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(fmt.Sprintf("Space separated Matrix %s entries (e.g. '1 2 3'): ", m.name))
	scanner.Scan()
	var errThree error
	var matRaw []string = strings.Split(strings.TrimSpace(scanner.Text()), " ")
	var mat [][]int
	for i := 0; i < rows; i++ {
		var rowTemp []int
		for j := 0; j < cols; j++ {
			var element int64
			element, errThree = strconv.ParseInt(matRaw[i*cols+j], 10, 64)
			if errThree != nil {
				log.Print("Can only accept integers. Please try again!")
				return m.setMatrix()
			}
			rowTemp = append(rowTemp, int(element))
		}
		mat = append(mat, rowTemp)
	}
	m.rows = rows
	m.cols = cols
	m.matrix = mat
	return m
}

func matMult(matA matrix, matB matrix) [][]int {
	var prodMat [][]int
	for i := 0; i < matA.rows; i++ {
		var rowTemp []int
		for j := 0; j < matB.cols; j++ {
			var prodMatIJ int = 0
			for k := 0; k < matA.cols; k++ {
				prodMatIJ += matA.matrix[i][k] * matB.matrix[k][j]
			}
			rowTemp = append(rowTemp, prodMatIJ)
		}
		prodMat = append(prodMat, rowTemp)
	}
	return prodMat
}

func main() {
	var matA = matrix{name: "A"}
	matA.setMatrix()
	fmt.Println(fmt.Sprintf("Matrix %s: ", matA.name))
	for i := 0; i < matA.rows; i++ {
		for j := 0; j < matA.cols; j++ {
			fmt.Printf("%d ", matA.matrix[i][j])
		}
		fmt.Println("")
	}
	var matB = matrix{name: "B"}
	matB.setMatrix()
	fmt.Println(fmt.Sprintf("Matrix %s: ", matB.name))
	for i := 0; i < matB.rows; i++ {
		for j := 0; j < matB.cols; j++ {
			fmt.Printf("%d ", matB.matrix[i][j])
		}
		fmt.Println("")
	}
	if matA.cols == matB.rows {
		fmt.Println("The product Matrix is: ")
		var prodMat = matMult(matA, matB)
		for i := 0; i < matA.rows; i++ {
			for j := 0; j < matB.cols; j++ {
				fmt.Printf("%d ", prodMat[i][j])
			}
			fmt.Println("")
		}
	} else {
		fmt.Println("Matrices can't be multiplied!")
	}
}
