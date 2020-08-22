package main

import "fmt"

const (
	line       = "======================="
	tempFormat = "%0.1f"
	rowFormat  = "| %8s | %8s |\n"
)

type fahrenheit float64
type celsius float64

//fahrenheit method converts celsius temperature to fahrenheit
func (c celsius) fahrenheit() fahrenheit {
	f := 9*c/5 + 32
	return fahrenheit(f)
}

//celsius method converts fahreneheit temperature to celsius
func (f fahrenheit) celsius() celsius {
	c := (f - 32) * 5 / 9
	return celsius(c)
}

//getRowFn type functions returns cells for a particular row
type getRowFn func(rowNumber int) (cell1, cell2 string)

//row cells containing F->C conversions
func getFToC(rowNumber int) (cell1, cell2 string) {
	
	var f fahrenheit = fahrenheit(rowNumber*5 - 40)
	var c celsius = f.celsius()

	//calculating cell1, cell2
	cell1 = fmt.Sprintf(tempFormat, f)
	cell2 = fmt.Sprintf(tempFormat, c)

	return cell1, cell2
}

//row cells containing C->F conversions
func getCToF(rowNumber int) (cell1, cell2 string) {
	var c celsius = celsius(rowNumber*5 - 40)
	var f fahrenheit = c.fahrenheit()

	//calculating cell1, cell2
	cell1 = fmt.Sprintf(tempFormat, c)
	cell2 = fmt.Sprintf(tempFormat, f)

	return cell1, cell2
}

/*drawTable takes Header's to display,
getRow func for getting appropriate row cells and
number of rows to display
*/
func drawTable(hdr1, hdr2 string, getRow getRowFn, rows int) {
	fmt.Println(line)
	fmt.Printf(rowFormat, hdr1,hdr2)
	fmt.Println(line)
	for currentRow := 0; currentRow < rows; currentRow++ {
		cell1, cell2 := getRow(currentRow)
		fmt.Printf(rowFormat, cell1, cell2)
	}
	fmt.Println(line)
}

func main() {

	// Total values b/w -40 and 100 with step of 5=> (100-(-40))/5 + 1 => 29
	drawTable("°C", "°F", getFToC, 29)
	drawTable("°F", "°C", getCToF, 29)


}
