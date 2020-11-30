package service

import "net/http"

func Service(writer http.ResponseWriter, request *http.Request)  {
	tmp := []byte("hello world")
	writer.Write(tmp)
}
