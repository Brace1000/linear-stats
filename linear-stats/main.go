package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"linear/correlation"
	"linear/linearstat"
	"linear/regreation"
)

/*
------------------------------------------------------------
-----------------------------------------------------------
--------------------Main begins here -----------------------
-----------------------------------------------------------
*/

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run main.go data.txt")
		return
	}
	fileName := os.Args[1]
	if !strings.HasSuffix(fileName, ".txt") {
		fmt.Printf("the file should be a .txt %q\n", fileName)
		return
	}
	y_values := linearstat.ReadFile(fileName)

	m, c := regreation.CalculateLinearRegression(y_values)
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", m, c)

	r := correlation.CalculatePearsonCorrelation(y_values)
	fmt.Printf("Pearson Corellation: r = %.10f\n", r)
}
