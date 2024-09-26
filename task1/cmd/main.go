package main

import (
	"hsse_go_homework/task1/pkg/library"
	"hsse_go_homework/task1/test"
	"hsse_go_homework/task1/tools"
)

func main() {
	// lib init
	var lib library.LibraryInterface = &library.Library{}
	lib.SetIdGenerator(tools.HashGen1)
	lib.SetStorage(library.BookMap{})

	test.BasicUsage(lib)
	test.SetNewIDGenerator(lib)
	test.SetNewStorage(lib)
	test.SetNewNonEmptyStorage(lib)
}
