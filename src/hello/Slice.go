// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	s1 := a[:5]
// 	s2 := a[0:5]
// 	s3 := a[5:10] //：号后面的数对应索引减1；前面的数正常
// 	fmt.Println(s1)
// 	fmt.Println(s2)
// 	fmt.Println(s3)

// 	s4 := make([]int, 3, 10) //3为已经存在数的数量，10为目前最大容量，若超过变为当前最大容量的2倍；若不设置最大容量，默认为当前存在的容量
// 	fmt.Println(s4)
// 	fmt.Println(len(s4), cap(s4))

// 	s5 := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
// 	sa := s5[2:5]
// 	fmt.Println(string(sa))
// 	sb := sa[1:3] //对于sa来说的索引为1~2个的元素，切除sa；
// 	fmt.Println(string(sb))
// 	//若超过sa大小,对于sa来说后面还有其他容量，即开始切片的数组的后面所有元素全部包含，故会取到后面对应的字符
// 	sc := sa[2:5]
// 	fmt.Println(string(sc))

// 	s7 := make([]int, 3, 6)
// 	fmt.Println(s7)
// 	s7 = append(s7, 1, 2, 3)
// 	fmt.Printf("%v %p\n", s7, s7)
// 	//超过最大容量，会重新创建一个数组，将原始数据拷贝到新数组中，所以得到的地址发生变化
// 	s7 = append(s7, 1, 2, 3)
// 	fmt.Printf("%v %p\n", s7, s7)

// 	b := []int{1, 2, 3, 4, 5}
// 	s8 := b[1:3]
// 	s9 := b[2:5]
// 	fmt.Println(s8, s9)
// 	//若改变其公共值，即b[2]
// 	s8[1] = 9
// 	fmt.Println(s8, s9, b) //其关联的所有数组或切片其值均改变

// 	s8 = append(s8, 1) //添加也相当于改变了其中的值，所有的相关值也会发生改变
// 	fmt.Println(s8, s9, b)

// 	s8 = append(s8, 2)
// 	fmt.Println(s8, s9, b)

// 	s8 = append(s8, 3) //若超出其最大容量，则会指向新数组，不会继续改变其原始数组了
// 	fmt.Println(s8, s9, b)

// 	//copy
// 	s10 := []int{1, 2, 3, 4}
// 	s11 := []int{5, 6, 7, 8, 9, 10, 11}
// 	// copy(s10, s11)
// 	// fmt.Println(s10, s11) //将s11的前s10个数的部分拷贝到s10
// 	copy(s11, s10)
// 	fmt.Println(s10, s11) //将所有s10拷贝到s11中前s10个数中,所有的变化中，数组长度是不变的

// 	c := []int{1, 2, 3, 4, 5}
// 	s12 := c[0:]
// 	fmt.Println(s12)
// }
