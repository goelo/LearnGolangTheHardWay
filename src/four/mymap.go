package main

import (
	"github.com/k0kubun/pp"
)

//定义一个struct
type bookInfo struct {
	ID string
	Name string
	Price string
}
func main() {
	var bookDB map[string]bookInfo
	bookDB = make(map[string]bookInfo)

	//插入数据
	bookDB["1"] = bookInfo{"1", "Harry Potter", "$20"}
	bookDB["123"] = bookInfo{"123", "Steve Jobs:A Biography","$12"}

	//在map中查找
	book, ok := bookDB["1234"]

	if ok {
		pp.Println("find the book: ", book.Name, "the price is: ", book.Price)
	} else {
		pp.Println("sorry, didn't find the book.")
	}
}
