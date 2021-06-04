package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(demo(12))
}
/**
0-100
二分法求平方根
 */
func demo(num float64) (gen float64) {
	start := 0.0
	end := num
	for  {
		s := (end-start)/2
		v := start + math.Floor(s)
		if start == v{
			return 0.0
		}
		v1 := v * v
		if v1 == num {
			return v
		}else if v1 > num{
			end = v
		}else{
			start = v
		}
	}
}