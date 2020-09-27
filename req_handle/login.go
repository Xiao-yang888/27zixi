package req_handle

import (
	"html/template"
	"net/http"
)

func Login(writer http.ResponseWriter,requset *http.Request){
	qqname := requset.FormValue("qq")

	if qqname== ""{
		tmpt,_:=template.ParseFiles("./qqhtml/error.html")
		tmpt.Execute(writer,"qq为空")
		return
	}

		temp,_:=template.ParseFiles("./qqhtml/jiegou.html")

		temp.Execute(writer, nil)

}
