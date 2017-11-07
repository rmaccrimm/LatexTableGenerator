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
	if len(s)+1 > t.ColWidth[j] {
		t.ColWidth[j] = len(s) + 1
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

func (t Table) GenTex() string {
	text := "\\begin{tabular}{|"
	for i := 0; i < t.Cols; i++ {
		text += "c|"
	}
	text += "} \\hline\n\t"
	for i, h := range t.Headers {
		text += h
		if i != t.Cols-1 {
			text += " & "
		}
	}
	text += " \\\\ \\hline\n\t"
	for i := 0; i < t.Rows; i++ {
		for j := 0; j < t.Cols; j++ {
			text += t.Contents[i][j]
			if j != t.Cols-1 {
				text += " & "
			}
		}
		text += " \\\\"
		if i != t.Rows-1 {
			text += "\n\t"
		}
	}
	text += " \\hline\n\\end{tabular}\n"
	return text
}
