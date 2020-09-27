package main

import (
	"encoding/hex"
	"fmt"
)

func main(){
	arr := []byte{'1','0','0','0','p','h','o','n','e'}
	fmt.Println(string(arr))
	str :=BytesToHexString(arr)
	fmt.Println(str)
	str = ReverseHexString(str)
	arr,_ = HexStringToBytes(str)
	fmt.Printf("%x\n",arr)
	//格式化打印
	ReverseBytes(arr)
	fmt.Println(string(arr))
}

func BytesToHexString(arr []byte)string{
	return hex.EncodeToString(arr)//编码
}

func HexStringToBytes(s string)([]byte, error){
	arr,err :=hex.DecodeString(s)//解码
	return arr,err
}

func ReverseHexString(hexStr string)string{
	arr,_:=hex.DecodeString(hexStr)
	ReverseBytes(arr)
	return hex.EncodeToString(arr)
}

func ReverseBytes(data []byte){
	for i , j := 0,len(data)-1;i < j;i , j = i+1,j-1{
		data[i],data[j] = data[j],data[i]
	}
}