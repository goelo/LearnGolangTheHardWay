package main

import (
	"fmt"
	"github.com/k0kubun/pp"
)

func main() {
	//数组的声明
	var arr1 [2] int
	var arr2 [10] int
	var arr3 [2][3]int		//二维数组
	var arr4 [2]*int		//指针数组

	//数组赋值
	arr1[0] = 0
	arr1[1] = 1
	//指针数组赋值
	v0 := 1
	v1 := 2
	arr4[0] = &v0
	arr4[1] = &v1
	//上述赋值适应于数组元素较少时，当数组元素较多，可以用for循环来处理
	for i:=0; i<2; i++ {
		for j:=0; j<3; j++ {
			arr3[i][j] = i + j
		}
	}
	//数组初始化
	var arr5 = [5]float32{1.0, 2.0, 3.4, 5.6, 7.8}
	arr6 := [2]int{0,1}
	//[]中的数组可以忽略，编译器会根据元素个数来设置
	var arr7 = [...]string{"first name", "second name", "third name"}

	pp.Println(arr1, arr2, arr3, arr4, arr5, arr6, arr7)

	arr := [5]int{1,2,3,4,5}

	modifyArray(arr)

	fmt.Println("In main(), array values", arr)
}

func modifyArray(arr [5]int) {
	for i:=0; i<len(arr); i++ {
		arr[i] = i + 1
	}
	fmt.Println("In modifyArray(), array values: ", arr)
}
//output:
/*
[2]int{
0,
1,
}
[10]int{
0,
0,
0,
0,
0,
0,
0,
0,
0,
0,
}
[2][3]int{
[3]int{
0,
1,
2,
},
[3]int{
1,
2,
3,
},
}
[2]*int{
&1,
&2,
}
[5]float32{
1.000000,
2.000000,
3.400000,
5.600000,
7.800000,
}
[2]int{
0,
1,
}
[3]string{
"first name",
"second name",
"third name",
}
In modifyArray(), array values:  [1 2 3 4 5]
In main(), array values [1 2 3 4 5]
*/
