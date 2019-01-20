package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {

	var v1, v2 bool
	v1 = true
	// v2 = 1 error
	v3 := false
	v4 := (1 != 2)

	fmt.Printf("%v %v %v %v\n", v1, v2, v3, v4)

	var v5 int16
	v5 = 15
	v6 := 16
	//Go语言中查看数据属于哪种类型，使用reflect.TypeOf()方法
	fmt.Println(reflect.TypeOf(v5))
	fmt.Println(reflect.TypeOf(v6))

	var v7 float32
	v7 = 1.0
	v8 := 1.0
	fmt.Println(reflect.TypeOf(v7),reflect.TypeOf(v8))

	var cpx1 complex64
	cpx1 = 1.0 + 2i
	cpx2 := 1.0 + 2i
	cpx3 := complex(1.0, 2)
	fmt.Println("cpx1 is", cpx1, "and type is", reflect.TypeOf(cpx1))
	fmt.Println("cpx2 is", cpx1, "and type is", reflect.TypeOf(cpx2))
	fmt.Println("cpx3 is", cpx1, "and type is", reflect.TypeOf(cpx3))

	s := "LearnGolangTheHardWay"
	//get the length of s
	fmt.Println(len(s))
	//find if the substring in the string
	fmt.Println(strings.Contains(s, "Hard"))
	fmt.Println(strings.ContainsAny(s, "L & z"))
	fmt.Println(strings.ContainsRune(s, 'H'))

	//contact the strings
	f := "find"
	sf:= "find"
	fmt.Println(f + s)

	//loop the strings
	for i:=0; i<len(s);i++ {
		ch := s[i]
		fmt.Println(i, ch)
	}

	fmt.Println("========")
	//another way for loop the strings
	for j, ca := range s {

		fmt.Println(j, ca)
	}

	//compare the strings
	fmt.Println(strings.Compare("Go", "go"))
	fmt.Println(strings.Compare(f, s))
	fmt.Println(strings.Compare(f, sf))
	fmt.Println(strings.EqualFold(f, s))
	fmt.Println(strings.EqualFold(f, sf))
	//Upper and Lower
	fmt.Println(strings.ToUpper(f))
	fmt.Println(strings.ToLower(strings.ToUpper(f)))
}