package main

const (
	dirty     = "💨"
	normal    = "😃"
	annoying  = "😁"
	available = "🆓"
)

type (
	desks [][]string
)

func main() {
	printFactored(normalDesks)
}
