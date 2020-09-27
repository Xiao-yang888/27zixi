package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func main(){
	//1,准备要输入端的文件，即beego文件

	file,err:=os.Open("./bitcoin.exe")
	if err != nil{
		fmt.Println(err.Error())
		return
	}

	//2,打开具体的文件
	hash256 :=sha256.New()//新建一个sha256对象

	//用sha256对file进行读取
	length,err:=io.Copy(hash256,file)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println("文件的长度是",length)

	//使用sha256对输入的内容进行hash计算
	hashResult:=hash256.Sum(nil)
	fmt.Println(hashResult)
	//16进制表现格式
	fmt.Printf("%x\n",hashResult)

}
