package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := "My favorite quote by Mark Twain: Twenty Years From Now You Will Be More Disappointed By The Things You Didn’t Do Than By The Ones You Did Do "
	s64 := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(s64)
	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatalln("something is wrong", err)
	}
	fmt.Println(string(bs))
}
