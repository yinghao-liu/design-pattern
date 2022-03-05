package decoratorpattern

import (
	"fmt"
	"time"
)

// 抽象计算类
type Calculater interface {
	Add()
}

// --------------------------------------------
// 计算类
type Calculate struct {
}

// 加
func (c Calculate) Add() {
	//start := time.Now()

	var i int
	i += 1
	time.Sleep(2 * time.Second)
	fmt.Printf("i is %d\n", i)

	// duration := time.Since(start)
	// fmt.Printf("cost is %+v ms\n", duration.Milliseconds())

}

// -------------------------------------------------
// 装饰类
type CalculateDecorate struct {
	cal Calculater
}

func (d *CalculateDecorate) SetCalculater(cal Calculater) {
	d.cal = cal
}
func (d CalculateDecorate) Add() {
	start := time.Now()

	d.cal.Add()

	duration := time.Since(start)
	fmt.Printf("cost is %+v ms\n", duration.Milliseconds())
}
