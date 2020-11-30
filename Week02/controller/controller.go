package controller

import (
	"fmt"
	"net/http"
	servcie "week02/service"
)

func Init()  {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		ret, err := servcie.Service()
		if err != nil {
			writer.Write([]byte(fmt.Sprintf("%+v\n", err)))
			return
		}
		writer.Write(ret)
	})
	err := http.ListenAndServe("127.0.0.1:12000", nil)
	if err != nil {
		fmt.Println("server start fail.")
	}
}
