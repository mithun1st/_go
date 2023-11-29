package main

import (
	"fmt"
)

// struct model
type Employee struct {
	eName  string
	salary float64
	bonus  float32
}
type Student struct {
	sName  string
	roll   int
	result float32
}

// function
func (e Employee) getEmpName() string {
	return e.eName
}
func (e Employee) getEmpTotSal() float64 {
	return e.salary + float64(e.bonus)
}

func (s Student) getStuName() string {
	return s.sName
}

// interface
type employeFnc interface {
	getEmpName()
	getStuName()
}

func main() {

	var emp Employee = Employee{
		eName:  "Mr X",
		salary: 12000.00,
		bonus:  200,
	}
	var stu Student = Student{
		sName:  "Alice",
		roll:   12,
		result: 4.21,
	}

	//emp
	fmt.Println(emp)
	fmt.Println(emp.getEmpName())
	fmt.Println(emp.getEmpTotSal())

	//stu
	fmt.Println(stu)
	fmt.Println(stu.getStuName())

}
