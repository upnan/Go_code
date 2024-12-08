package main

import "fmt"

type Income interface {
	calculate() int //用于计算并返回来源收入
	source() string //返回来源名称
}

type FixedBilling struct {
	projectName  string //项目名称
	biddedAmount int    //组织为该项目出的投标金额
}

type TimeAndMaterial struct { //表示时间和材料类型
	projectName string
	noOfHours   int
	hourlyRate  int
}

// 方法计算并返回实际收入和收入来源
// 不同项目类型利用不同的计算方式
func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

/*由于FixedBilling和TimeAndMaterial结构都提供了Income接口的calculate()和source()方法的定义，
因此这两个结构都实现了Income接口。*/

//该函数实现了多态性，不同接口对应着不同的实际类型
func calculateNetIncome(ic []Income) { //Income接口切片作为输入参数
	var netIncome int = 0
	for _, income := range ic { //遍历该接口切片，会自动识别实际类型，类似虚函数
		fmt.Printf("Income from %s = $%d\n", income.source(), income.calculate())
		netIncome += income.calculate()
	}
	fmt.Printf("Net income of organization = $%d", netIncome)
}

func main() {
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	incomeStreams := []Income{project1, project2, project3}
	calculateNetIncome(incomeStreams)
}
