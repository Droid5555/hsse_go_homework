package main

import (
	"hsse_go_homework/task2/api/client"
	"hsse_go_homework/task2/test"
)

func main() {
	c := client.New("http://localhost:8080")

	test.GetVersion(c)
	test.PostDecode(c)
	test.GetHardOp(c)

}
