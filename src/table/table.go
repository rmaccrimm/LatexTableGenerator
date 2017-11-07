package table

import "fmt"
import "strconv"

const CENTER = 0
const RIGHT = 1
const LEFT = 2

type Table struct {
	Rows, Cols                      int
	ColDiv, RowDiv                  bool
	TopRule, BottomRule, HeaderRule bool
	Contents                        [][]string
	Headers                         []string
	ColWidth                        []int
}

func (t Table) SetHeader(i int, s string) {
	t.Headers[i] = s
	if len(s)+1 > t.ColWidth[i] {
		t.ColWidth[i] = len(s) + 1
	}

}

func (t Table) SetElement(i, j int, s string) {
	t.Contents[i][j] = s
	if len(s)+1 > t.ColWidth[i] {
		t.ColWidth[i] = len(s) + 1
		fmt.Println("colwidth: ", t.ColWidth[i])
	}
}

func (t Table) Display() {
	for i, s := range t.Headers {
		w := strconv.Itoa(t.ColWidth[i])
		f := "%" + w + "s"
		fmt.Printf(f, s)
	}
	fmt.Println()
	for _, r := range t.Contents {
		for i, e := range r {
			w := strconv.Itoa(t.ColWidth[i])
			f := "%" + w + "s"
			fmt.Printf(f, e)
		}
		fmt.Println()
	}
}

func (t Table) GenTex() {
	text := "\\begin{tabular}"

	text += "\\end{tabular}"
}
