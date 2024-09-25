package main

import (
	"hsse_go_homework/task1/structs"
	"hsse_go_homework/task1/utils"
)

func main() {
	// lib init
	var lib structs.LibraryInterface = &structs.Library{}
	lib.SetIdGenerator(utils.Hash1)
	lib.SetStorage(structs.BookMap{})
	utils.Test1(lib)
	utils.Test2(lib)
	utils.Test3(lib)
	utils.Test4(lib)
}
