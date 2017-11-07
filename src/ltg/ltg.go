package main

import (
	"fmt"
	"table"
)

func main() {
	var nrows, ncols int
	fmt.Print("Number of rows: ")
	fmt.Scan(&nrows)
	fmt.Print("Number of columns: ")
	fmt.Scan(&ncols)

	tbl := table.Table{
		Rows: nrows, Cols: ncols,
		ColDiv:  true,
		TopRule: true, BottomRule: true, HeaderRule: true,
		Contents: make([][]string, nrows),
		Headers:  make([]string, ncols),
		ColWidth: make([]int, ncols)}

	for i := range tbl.Contents {
		tbl.Contents[i] = make([]string, ncols)
	}

	fmt.Print("Enter column headers: ")
	var header string
	for i := range tbl.Headers {
		fmt.Scan(&header)
		tbl.SetHeader(i, header)
	}
	getInput(&tbl)
	//tbl.Display()
	fmt.Print(tbl.GenTex())
}

func getInput(t *table.Table) {
	var input string
	for i := 0; i < t.Cols; i++ {
		for j := 0; j < t.Rows; j++ {
			fmt.Print("r", j, ", c", i, ": ")
			fmt.Scan(&input)
			t.SetElement(j, i, input)
		}
	}
}
