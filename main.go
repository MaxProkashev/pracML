package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

var (
	nameTable = "./presidentUSA.xlsx"
	sheetName = "list"
	ageInCell = 5
)

func main() {
	var (
		x       []float64
		buf     float64
		predict float64
		p       int
	)

	wb, _ := xlsx.OpenFile(nameTable)
	sh, ok := wb.Sheet[sheetName]
	if !ok {
		fmt.Println("Sheet does not exist")
		return
	}

	for i := 1; i < sh.MaxRow; i++ {
		buf, _ = sh.Rows[i].Cells[ageInCell].Float()
		x = append(x, buf)
	}

	for p = 2; p <= len(x)/2; p++ {
		predict = 0
		_, a := lineFactors(x, len(x), p)

		for i := 0; i < len(a); i++ {
			predict += a[i] * x[len(x)-len(a)+i]
		}
		fmt.Printf("предполагаемый возраст по предыдущим %d президентам: %.2f\n",
			p, predict)
	}
}

func lineFactors(x []float64, N int, p int) (E float64, a []float64) {
	R := make([]float64, p, p)
	k := make([]float64, p, p)
	a = make([]float64, p, p)
	for i := 0; i < p; i++ {
		for n := 0; n < N-i; n++ {
			R[i] += x[n] * x[n+i]
		}
		R[i] /= float64(N)
	}

	E = R[0]

	for l := 1; l < p; l++ {
		for i := 1; i < l; i++ {
			k[l] += a[i] * R[l-i]
		}

		k[l] = (k[l] - R[l]) / E
		a[l] = -k[l]

		for j := 1; j < l; j++ {
			a[j] += k[l] * a[l-j]
		}
		E = E * (1 - k[l]*k[l])
	}

	// fmt.Println("коэфициенты a[i]: ")
	// for i := range a {
	// 	fmt.Printf("a[%d] = %f\n", i+1, a[i])
	// }
	// fmt.Printf("оценка ошибки: %.2f\n", E)
	return E, a
}
