func convert(s string, numRows int) string {
	if numRows <= 1 || len(s) < numRows {
		return s
	}
	/*
	         0 1 2
	         - - -
	   row 0|
	       1|
	       2|
	*/
	rows := make([]strings.Builder, numRows)
	row := 0
	step := 1 // +1 is up, and -1 is down
	for _, c := range s {
		rows[row].WriteString(string(c))
		if row+step >= numRows || row+step < 0 {
			step *= -1
		}
		row += step
	}
	var result strings.Builder
	for _, r := range rows {
		result.WriteString(r.String())
	}
	return result.String()

}
