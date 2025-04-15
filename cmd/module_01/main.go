package main

import "fmt"

var (
	name = "Alice"
	s    = []int{2, 4, 7}
)

func main() {
	fmt.Printf("Привет, %s! Твои оценки: %v. Средняя оценка: %.2f.", name, s, middelNumber(s))
}

func middelNumber(s []int) float64 {
	var num float64
	for _, v := range s {
		num += float64(v)
	}
	return num / float64(len(s))
}
