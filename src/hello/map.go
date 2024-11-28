// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	var m map[int]map[int]string     //键值为int型,存储的值为string
// 	m = make(map[int]map[int]string) //该处只是创建了第一层map，需要进行第二层的创建
// 	a, ok := m[1][1]
// 	if !ok {
// 		m[1] = make(map[int]string)
// 	}
// 	m[1][2] = "GOOD"
// 	a, ok = m[1][2]
// 	fmt.Println(a, ok)

// 	//map迭代
// 	sm := make([]map[int]string, 5) //大小为5的map切片
// 	//sm:=make([]int,5)	//这样是错的，类型不一致
// 	for i := range sm {
// 		sm[i] = make(map[int]string) //创建map
// 		sm[i][1] = "ok"
// 		fmt.Println(sm[i])
// 	}
// 	fmt.Println(sm)

// 	//map本身是无序的，通过slice进行map中键进行排序，以此来排序map
// 	m1 := map[int]string{1: "a", 3: "c", 2: "b", 5: "d", 7: "e"}
// 	sm1 := make([]int, len(m1))
// 	i := 0
// 	for k, _ := range m1 { //range可以取到键和值，可通过下划线舍弃不需要的
// 		sm1[i] = k
// 		i++
// 	}
// 	fmt.Println(sm1)
// 	sort.Ints(sm1)
// 	fmt.Println(sm1)
// 	for i := 0; i < len(sm1); i++ {
// 		fmt.Println(m1[sm1[i]])
// 	}

// 	ma := map[int]string{1: "a", 3: "c", 2: "b"}
// 	mb := make(map[string]int, 3)
// 	for k, v := range ma {
// 		mb[v] = k
// 	}
// 	fmt.Println(mb)
// }
