package main

import "fmt"

func printFactored(desks [][]string) {
	printTitle()
	printLegend()
	printDesks(desks)
	printFooter()
}

func printTitle() {
	fmt.Printf("\n~DESK PRINTER~\n\n")
}

func printLegend() {
	fmt.Println(available + " = Available Desk")
	fmt.Println(normal + " = Normal Person")
	fmt.Println(annoying + " = Annoying Person")
	fmt.Println(sick + " = Sick Person")
	fmt.Println("")
}

func printDesks(desks [][]string) {
	for _, row := range desks {
		rowStr := ""

		for _, desk := range row {
			rowStr += desk + " "
		}

		fmt.Println(rowStr)
	}
}

func printFooter() {
	fmt.Println("\n© copyright tomontheinternet")
}
