package main

import (
	"fmt"

	"github.com/MrAndiw/go-calculator/math"
)

func main() {
	math := math.NewMath()

	fmt.Println(math.SumNumber(5, 5))
}
