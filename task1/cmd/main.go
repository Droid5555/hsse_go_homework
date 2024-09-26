package main

import (
	"hsse_go_homework/task1/pkg/library"
	"hsse_go_homework/task1/test"
	"hsse_go_homework/task1/tools"
)

func main() {
	// lib init
	var lib library.LibraryInterface = &library.Library{}
	lib.SetIdGenerator(tools.Hash1)
	lib.SetStorage(library.BookMap{})
	test.Test1(lib)
	test.Test2(lib)
	test.Test3(lib)
	test.Test4(lib)
}
