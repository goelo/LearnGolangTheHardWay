package main

import (
	"fmt"
	"github.com/k0kubun/pp"
)

func main() {
	mySlice := make([]int, 5, 10)

	fmt.Println("len(mySlice): ", len(mySlice))
	fmt.Println("cap(mySlice): ", cap(mySlice))

	//往mySlice元素后追加新元素，形成新数组切片
	mySlice = append(mySlice, 1,2,3)

	mySlice2 := []int{8,9,10}

	mySlice = append(mySlice, mySlice2...)
	/* 上面这行代码mySlice2后面的三个点如果缺少，会编译错误。*/
	/* 原因是append的函数定义从第二个参数起，都是可增加参数。*/
	/* mySlice的元素类型是init, 直接传递的mySlice2是数组切片，类型错误。*/
	/* 加上省略号即三个点相当于把mySlice2包含的所有元素打散后传入 */
	/* 相当于 mySlice = append(mySlice, 8,9,10) */
	fmt.Println("mySlice: ", mySlice)
	/*上面追加的元素超过了原来的10个元素容量，此时数组切片会重新自动处理存储空间不足，自动分配一块足够大的内存*/
	slice1 := []int{1,2,3,4,5}
	slice2 := []int{5,4,3}

	copy(slice1, slice2)    //只复制slice2的3个元素到slice1的前三个位置
	copy(slice2, slice1) 	//只复制slice1的前三个元素


	pp.Println(slice1, slice2)

}
