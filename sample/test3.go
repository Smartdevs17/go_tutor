// Calculate maxmal value in a slice
package sample

import (
	"fmt"
	"math"
)

func GetMaxValue(slice []int) {
	max := math.MinInt32
	for _, value := range slice {
		if value > max {
			max = value
		}
	}
	fmt.Println("max:", max)
}
