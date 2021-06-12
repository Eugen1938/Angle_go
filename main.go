package main

import (
	"fmt"
	"math"
)

func main() {
	printTime123(12, 30)
}
func printTime123(hour, minute int) {

	fmt.Printf("Время: %d:%d \n", hour, minute)

	angle :=calculateAngle(hour, minute)
	fmt.Printf("Угол: %.1f \n", angle)

}

func calculateAngle(hour, minute int) float64 {
	return math.Abs(30*float64(hour)-5.5*float64(minute))
}
