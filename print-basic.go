package main

import "fmt"

func printBasic(desks [][]string) {
	// title
	fmt.Printf("\n~DESK PRINTER~\n\n")

	// legend
	fmt.Println(available + " = Available Desk")
	fmt.Println(dirty + " = Dirty Desk")
	fmt.Println(normal + " = Normal Person")
	fmt.Println(annoying + " = Annoying Person")
	fmt.Println("")

	// desks
	for _, row := range desks {
		rowStr := ""

		for _, desk := range row {
			rowStr += desk + " "
		}

		fmt.Println(rowStr)
	}

	// footer
	fmt.Println("\n© copyright tomontheinternet")
}
