package test

import (
	"fmt"
	"hsse_go_homework/task2/api/client"
	"log"
)

func GetVersion(c *client.Client) {
	v, err := c.GetVersion()
	if err != nil {
		return
	}
	fmt.Println(v)
}

func GetHardOp(c *client.Client) {
	ok, text, _ := c.GetHardOp()
	if ok {
		fmt.Printf("%v, %s", ok, text)
	} else {
		fmt.Printf("%v", ok)
	}
}

func PostDecode(c *client.Client) {
	// line was taken from: https://ru.wikipedia.org/wiki/Base64
	line := "TWFuIGlzIGRpc3Rpbmd1aXNoZWQsIG5vdCBvbmx5IGJ5IGhpcyByZWFzb24sIGJ1dCBieSB0\naGlzIHNpbmd1bGFyIHBhc3Npb24" +
		"gZnJvbSBvdGhlciBhbmltYWxzLCB3aGljaCBpcyBhIGx1\nc3Qgb2YgdGhlIG1pbmQsIHRoYXQgYnkgYSBwZXJzZXZlcmFuY2Ugb2YgZG" +
		"VsaWdodCBpbiB0\naGUgY29udGludWVkIGFuZCBpbmRlZmF0aWdhYmxlIGdlbmVyYXRpb24gb2Yga25vd2xlZGdl\nLCBleGNlZWRzIHR" +
		"oZSBzaG9ydCB2ZWhlbWVuY2Ugb2YgYW55IGNhcm5hbCBwbGVhc3VyZS4="
	response, err := c.PostDecode(line)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(response)
}
