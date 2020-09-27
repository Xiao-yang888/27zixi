package req_handle

import (
	"html/template"
	"net/http"
)

func CommentPage(writer http.ResponseWriter, req *http.Request) {
	temp,err:=template.ParseFiles("./qqhtml/comment.html")
	if err != nil {
		writer.Write([]byte(err.Error()))
		errorTmp,_:=template.ParseFiles("./qqhtml/error.html")
		errorTmp.Execute(writer,err.Error())
		return
	}
	temp.Execute(writer,nil)
}
