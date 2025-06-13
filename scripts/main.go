package main

import (
	"fmt"
)

// 接口：任何类型只要有 Drive 方法，就是 Driver
type Driver interface {
	Drive()
}

type Car struct {
	Name string
}

// 用指针接收者实现 Drive 方法
func (c Car) Drive() {
	fmt.Println(c.Name, "is driving")
}

func main() {
	var c1 Car = Car{Name: "BMW"}
	var c2 *Car = &Car{Name: "Tesla"}

	var d Driver
	d = c1
	d.Drive()
	d = c2 // ✅ OK，*Car 实现了 Driver 接口

}
