type Spreadsheet struct {
	rows  int
	cells map[string]int 
}

func Constructor(rows int) Spreadsheet {
	return Spreadsheet{
		rows:  rows,
		cells: make(map[string]int),
	}
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	this.cells[cell] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	delete(this.cells, cell)
}

func (this *Spreadsheet) GetValue(formula string) int {
	if len(formula) == 0 || formula[0] != '=' {
		return 0
	}
	expr := formula[1:]
	parts := strings.Split(expr, "+")
	if len(parts) != 2 {
		return 0
	}
	parse := func(tok string) int {
		t := strings.TrimSpace(tok)
		if t == "" {
			return 0
		}
		first := t[0]
		if first >= 'A' && first <= 'Z' {
			if v, ok := this.cells[t]; ok {
				return v
			}
			return 0
		}
		val, _ := strconv.Atoi(t)
		return val
	}
	return parse(parts[0]) + parse(parts[1])
}
