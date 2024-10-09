package main

import (
	"hsse_go_homework/task2/test"
	"net/http"
)

func main() {
	client := &http.Client{}

	test.GetVersion(client)
	test.PostDecode(client)
	test.GetHardOp(client)

}
