package main

import (
	"fmt"
	"math"
)

func MySqrt(x float64) float64 {
	z := 1.0
	var prv_z float64
	for i:=0; i < 100; i++ {
		prv_z = z
		z -= (z*z - x) / (2*z)
		i += 1
		fmt.Printf("The z value is %v over the iteration #%v \n", z, i)
		
		if math.Abs(z-prv_z)<0.0001{
			fmt.Printf("Stopping loop on iteration: #%v \n",i)
			return z
		}
	}
	
	return z
}

func main() {
	fmt.Println(MySqrt(3025))
}
