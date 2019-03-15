package main

import "fmt"

var arr = []int{1, 2, 3}

func main() {
	c := merge([]int{1, 3, 5}, []int{2, 4, 6})
	fmt.Println(c)

	//fmt.Println(del([]int{1, 2, 3, 4, 5, 6}, 4))
}

//实现一个支持动态扩容的数组
func array(r []int) {

}

/*
	假设数组为整型数组
	实现一个大小固定的有序数组，支持动态增删改操作
*/

func add(i int) {
	arr = append(arr, i)
}

func del(c []int, index int) []int {
	for i := index; i < len(c)-1; i++ {
		c[i] = c[i+1]
	}
	//c = append(c[:index], c[index+1:]...)
	return c
}

/*
	假设数组为整型数组
	实现两个有序数组合并为一个有序数组
	假设数组为升序排列
*/
func merge(a, b []int) []int {
	l := len(a) + len(b)
	var c = make([]int, 0, l)

	i, j := 0, 0
	for {
		if a[i] <= b[j] {
			c = append(c, a[i])
			i++ //数组a右移
			if i >= len(a) {
				i = -1
				break
			}
		} else {
			c = append(c, b[j])
			j++ //数组b右移
			if j >= len(b) {
				j = -1
				break
			}
		}
	}

	if i == -1 { //数组a读完了
		c = append(c, b[j:]...)
	} else { //数组b读完了
		c = append(c, a[i:]...)
	}

	fmt.Println(len(c), cap(c))
	fmt.Println(del(c, 2))
	fmt.Println(len(c), cap(c))
	fmt.Println(del(del(c, 2), 2))
	fmt.Println(len(c), cap(c))

	return c
}
