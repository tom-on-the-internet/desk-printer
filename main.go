package main

const (
	sick      = "😷"
	normal    = "😃"
	annoying  = "🤡"
	available = "🆓"
)

type (
	desks [][]string
)

func main() {
	printFactored(normalDesks)
}
