package main

import (
	"qqluck/req_handle"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const APPKEY = "dfd0e2c15b81499a422ea5d38eaeecae" //申请的APPKEY

func main(){
	Request1()

	http.HandleFunc("/",req_handle.Index)

	http.HandleFunc("/login",req_handle.Login)

	http.HandleFunc("/comment.html",req_handle.CommentPage)

	fmt.Println("大部分全网络开始监听")

	err := http.ListenAndServe(":9095",nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

//1.QQ号码测吉凶
func Request1(){
	//请求地址
	juheURL :="http://japi.juhe.cn/qqevaluate/qq"

	//初始化参数
	param:=url.Values{}

	//配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
	param.Set("key",APPKEY) //您申请的appKey
	param.Set("qq","1398078735") //需要测试的QQ号码


	//发送请求
	data,err:=Get(juheURL,param)
	if err!=nil{
		fmt.Errorf("请求失败,错误信息:\r\n%v",err)
	}else{
		var netReturn map[string]interface{}
		json.Unmarshal(data,&netReturn)
		if netReturn["error_code"].(float64)==0{
			fmt.Printf("接口返回result字段是:\r\n%v",netReturn["result"])
		}
	}
}

// get 网络请求
func Get(apiURL string,params url.Values)(rs[]byte ,err error){
	var Url *url.URL
	Url,err=url.Parse(apiURL)
	if err!=nil{
		fmt.Printf("解析url错误:\r\n%v",err)
		return nil,err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery=params.Encode()
	resp,err:=http.Get(Url.String())
	if err!=nil{
		fmt.Println("err:",err)
		return nil,err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)


}

// post 网络请求 ,params 是url.Values类型
func Post(apiURL string, params url.Values)(rs[]byte,err error){
	resp,err:=http.PostForm(apiURL, params)
	if err!=nil{
		return nil ,err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
func dbmysql(){
	//连接数据库
	database, err := sql.Open("mysql", "12345678@qqluck(127.0.0.1:3306)/stu?charset=utf8")
	defer database.Close()
	if err != nil {
		fmt.Println(err.Error)
		return
	}
}