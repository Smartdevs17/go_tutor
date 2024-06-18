package sample

import (
	"fmt"
	"strconv"
)

// EvenEndedNumber -
func EvenEndedNumber() {

	count := 0
	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ {
			n := a * b
			s := strconv.Itoa(n)
			if s[0] == s[len(s)-1] {
				count++
			}
		}
	}
	fmt.Println("Total Count", count)
}
