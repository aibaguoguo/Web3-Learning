package main

import "fmt"

// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
func test1() {
	r := Rectangle{}
	r.Perimeter()
	r.Area()
	c := Circle{}
	c.Perimeter()
	c.Area()
}

type Shape interface {
	Area()
	Perimeter()
}
type Rectangle struct {
}
type Circle struct {
}

func (re *Rectangle) Area() {
	fmt.Println("Rectangle 的Area方法")
}

func (re *Rectangle) Perimeter() {
	fmt.Println("Rectangle 的Perimeter方法")
}

func (ci *Circle) Area() {
	fmt.Println("Circle 的area方法")
}

func (re *Circle) Perimeter() {
	fmt.Println("Circle 的Perimeter方法")
}

// 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
// 再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
// 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
func test2() {
	e := Employee{
		Person{"peter", 33},
		"001",
	}

	e.PrntInfo()

}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (emp *Employee) PrntInfo() {
	fmt.Println(emp)
}
