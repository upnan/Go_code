package main

import (
	"fmt"
	"math"
)

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		//return 0, errors.New("radius<0") //此处新建了一种错误
		//利用fmt包中函数实现，显示半径的值
		return 0, fmt.Errorf("%0.2f<0", radius)
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20
	area, error := circleArea(float64(radius))
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Printf("%.2f", area)
}
