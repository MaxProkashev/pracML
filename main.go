package main

import (
	"fmt"
	"strconv"

	"github.com/sajari/regression"
	"github.com/tealeg/xlsx"
)

var (
	nameTable = "./presidentUSA.xlsx"
	sheetName = "list"
	ageInCell = 5

	setNum = func(n int) string {
		return "(Возраст -" + strconv.Itoa(n) + " президента)"
	}
)

func main() {
	var (
		x       []float64
		buf     float64
		predict float64
		deep    int = 2
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

	// for p = 2; p <= len(x)/2; p++ {
	// 	predict = 0
	// 	_, a := lineFactors(x, len(x), p)

	// 	for i := 0; i < len(a); i++ {
	// 		predict += a[i] * x[len(x)-len(a)+i]
	// 	}
	// 	fmt.Printf("предполагаемый возраст по предыдущим %d президентам: %.2f\n",
	// 		p, predict)
	// }

	fmt.Println("deep?")
	fmt.Scan(&deep)
	if deep < 2 {
		deep = 2
		fmt.Println("min deep = 2")
	}

	r := new(regression.Regression)
	r.SetObserved("Возраст нового президента")

	for i := 0; i < deep; i++ {
		r.SetVar(i, setNum(i+1))
	}

	for i := 0; i < len(x)-deep; i++ {
		tr := make([]float64, deep)
		for j := 0; j < deep; j++ {
			tr[j] = x[i+j]
		}
		r.Train(
			regression.DataPoint(x[i+deep], tr),
		)
	}
	r.Run()

	fmt.Println("Полученные кофициенты:")
	fmt.Printf("Свободный коэфициент = %.10f\n", r.GetCoeffs()[0])
	for i := 1; i <= deep; i++ {
		fmt.Printf("%d коэфициент = %.10f\n", i, r.GetCoeffs()[i])
	}
	fmt.Printf("\n\n")

	tr := make([]float64, deep)
	for j := 0; j < deep; j++ {
		tr[j] = x[len(x)-j-1]
	}
	predict, _ = r.Predict(tr)
	fmt.Println("Предсказанный возраст: ", predict)
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
