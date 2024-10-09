package test

import (
	"fmt"
	"hsse_go_homework/task2/api/decode"
	"hsse_go_homework/task2/api/hard_op"
	"hsse_go_homework/task2/api/version"
	"log"
	"net/http"
)

func GetVersion(client *http.Client) {
	v, err := version.GetVersion(client)
	if err != nil {
		return
	}
	fmt.Println(v)
}

func GetHardOp(client *http.Client) {
	ok, text, _ := hard_op.GetHardOp(client)
	if ok {
		fmt.Printf("%v, %s", ok, text)
	} else {
		fmt.Printf("%v", ok)
	}
}

func PostDecode(client *http.Client) {
	// line was taken from: https://ru.wikipedia.org/wiki/Base64
	line := "TWFuIGlzIGRpc3Rpbmd1aXNoZWQsIG5vdCBvbmx5IGJ5IGhpcyByZWFzb24sIGJ1dCBieSB0\naGlzIHNpbmd1bGFyIHBhc3Npb24" +
		"gZnJvbSBvdGhlciBhbmltYWxzLCB3aGljaCBpcyBhIGx1\nc3Qgb2YgdGhlIG1pbmQsIHRoYXQgYnkgYSBwZXJzZXZlcmFuY2Ugb2YgZG" +
		"VsaWdodCBpbiB0\naGUgY29udGludWVkIGFuZCBpbmRlZmF0aWdhYmxlIGdlbmVyYXRpb24gb2Yga25vd2xlZGdl\nLCBleGNlZWRzIHR" +
		"oZSBzaG9ydCB2ZWhlbWVuY2Ugb2YgYW55IGNhcm5hbCBwbGVhc3VyZS4="
	response, err := decode.PostDecode(client, line)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(response)
}
