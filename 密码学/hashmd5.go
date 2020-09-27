package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	str := "c190604"
	fmt.Println(str)
	hashMD5 := md5.New()
	_, err := hashMD5.Write([]byte(str))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	md5Bytes := hashMD5.Sum(nil)
	fmt.Printf("%x",md5Bytes)
}
