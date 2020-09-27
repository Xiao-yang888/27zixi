package req_handle

import (
	"html/template"
	"net/http"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("./qqhtml/qqindex.html")
	if err != nil {
		writer.Write([]byte(err.Error()))
		errorTmp, _ := template.ParseFiles("./qqhtml/error.html")
		errorTmp.Execute(writer, err.Error())
		return
	}
	temp.Execute(writer, nil)
}
