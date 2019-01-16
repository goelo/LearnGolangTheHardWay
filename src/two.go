package main
import "fmt"

const myname string = "Morgan"
const zero = 0.0
const (
	number int64 = 1024
	result = "success"
)
const a, b, c = 1, 2.0, "good"

const offset = 3 >> 2

const (
	a0 = iota
	a1
	a2
)

func main() {

	var v1 int
	var v2 string
	var (
		v3 float64
		v4 *int
	)
	var v5, v6, v7 int

	v1 = 1
	v2 = "hello"
	v3 = 2.0
	v4 = &v1
	v5, v6, v7 = 3,4,5

	var i, j int = 1,2
	i, j = j, i
	v8 := "world"
	//v4 = &v1

	fmt.Println("v1 = ", v1)
	fmt.Println(v2 + "," + v8)
	fmt.Println("v3 = ", v3)
	fmt.Println("v4 = ", v4)
	fmt.Println(v5,v6,v7)

	fmt.Println("v8 = ", v8)
	fmt.Println("i, j = ", i, j)

	fmt.Printf("%v\n", myname)
	fmt.Printf("%v\n", zero)
	fmt.Printf("number is %v, result is %v\n", number, result)
	fmt.Printf("a= %v, b=%v, c=%v\n", a,b,c)
	fmt.Printf("offset is %v\n", offset)
	fmt.Printf("a0 = %v, a1 = %v, a2 = %v", a0, a1, a2)
}

