package controller

import (
	"net/http"
	servcie "week02/service"
)

func Init()  {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		servcie.Service(writer, request)
	})
	http.ListenAndServe("127.0.0.1:12000", nil)
}
