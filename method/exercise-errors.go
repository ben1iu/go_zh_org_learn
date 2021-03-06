package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	if(e < 0){
		return fmt.Sprintf("cannot Sqrt negative number: %f",float64(e))
	}
	return fmt.Sprintf("OK");
}

func Sqrt(x float64) (float64, error) {
	return math.Sqrt(x),ErrNegativeSqrt(x)
	
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

