package main

import "oop/employee"

func main() {
	// e := employee.Employee{
	// 	FirstName:   "Sam",
	// 	LastName:    "Adolf",
	// 	TotalLeaves: 30,
	// 	LeavesTaken: 20,
	// }
	//var e employee.employee //零值创建的变量不可用
	e := employee.New("Sam", "Adolf", 30, 20) //结构体有效完成了类的实现，New方法代替构造函数
	e.LeavesRemaining()
}
