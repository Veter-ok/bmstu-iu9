package main

import "fmt"

type Fraction struct {
	num, den int
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func multiplyFrac(a, b Fraction) Fraction {
	return Fraction{a.num * b.num, a.den * b.den}
}

func subtractionFrac(a, b Fraction) Fraction {
	return Fraction{(a.num * b.den) - (a.den * b.num), a.den * b.den}
}

func (f Fraction) simplifyFrac() Fraction {
	g := gcd(abs(f.num), abs(f.den))
	f.num /= g
	f.den /= g
	if f.den < 0 {
		f.num = -f.num
		f.den = -f.den
	}
	return f
}

func subtractionRows(matrix [][]Fraction, N, startRow, startCol int) [][]Fraction {
	for i := startRow + 1; i < N; i++ {
		k := Fraction{
			matrix[i][startCol].num * matrix[startRow][startCol].den,
			matrix[i][startCol].den * matrix[startRow][startCol].num,
		}.simplifyFrac()
		for j := 0; j < N+1; j++ {
			a := multiplyFrac(k, matrix[startRow][j])
			matrix[i][j] = subtractionFrac(matrix[i][j], a)
		}
	}
	return matrix
}

func swapRows(matrix [][]Fraction, rowA, rowB, N int) {
	for i := 0; i < N+1; i++ {
		matrix[rowA][i], matrix[rowB][i] = matrix[rowB][i], matrix[rowA][i]
	}
}

func triangularMatrix(matrix [][]Fraction, N int) [][]Fraction {
	for A, B := 0, 0; A < N && B < N+1; A, B = A+1, B+1 {
		isSwap := false
		for j := A; j < N+1; j++ {
			for i := B; i < N; i++ {
				if matrix[i][j].num != 0 {
					swapRows(matrix, A, i, N)
					isSwap = true
					break
				}
			}
			if isSwap {
				break
			}
			B++
		}
		matrix = subtractionRows(matrix, N, A, B)
	}
	return matrix
}

func isNoSolution(matrix [][]Fraction, N int) bool {
	for i := 0; i < N; i++ {
		allZero := true
		for j := 0; j < N; j++ {
			if matrix[i][j].num != 0 {
				allZero = false
				break
			}
		}
		if allZero && matrix[i][N].num != 0 {
			return true
		}
	}
	return false
}

func gauss(matrix [][]Fraction, N int) []Fraction {
	matrix = triangularMatrix(matrix, N)
	if isNoSolution(matrix, N) {
		return nil
	}
	roots := make([]Fraction, N)
	for i := N - 1; i >= 0; i-- {
		sum := Fraction{0, 1}
		for j := i + 1; j < N; j++ {
			sum = Fraction{
				sum.num*matrix[i][j].den*roots[j].den + sum.den*matrix[i][j].num*roots[j].num,
				sum.den * matrix[i][j].den * roots[j].den,
			}
		}
		roots[i] = Fraction{
			matrix[i][N].num*sum.den - sum.num*matrix[i][N].den,
			matrix[i][i].num * sum.den,
		}.simplifyFrac()
	}
	return roots
}

func main() {
	var N, n int
	fmt.Scan(&N)
	matrix := make([][]Fraction, N)
	for i := 0; i < N; i++ {
		matrix[i] = make([]Fraction, N+1)
		for j := 0; j < N+1; j++ {
			fmt.Scan(&n)
			matrix[i][j] = Fraction{n, 1}
		}
	}
	roots := gauss(matrix, N)
	if roots == nil {
		fmt.Println("No solution")
	} else {
		for i := 0; i < N; i++ {
			fmt.Printf("%d/%d\n", roots[i].num, roots[i].den)
		}
	}
}
