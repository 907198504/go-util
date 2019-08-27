package httputil

import "net/http"

func HelloController(writer http.ResponseWriter, request *http.Request) {
	WriteStringToResponse(writer, "hello")
}

/**
StartHttpServer
*/
func StartHttpServer() {
	// 绑定controller
	http.Handle("/", http.HandlerFunc(HelloController))
	// 启动
	err := http.ListenAndServe(":8188", nil)
	// 异常处理
	if err != nil {
		panic(err)
	}
}

/**
WriteBytesToResponse
*/
func WriteBytesToResponse(writer http.ResponseWriter, bytes []byte) {
	_, err := writer.Write(bytes)
	if err != nil {
		panic(err)
	}
}

/**
WriteBytesToResponse
*/
func WriteStringToResponse(writer http.ResponseWriter, str string) {
	WriteBytesToResponse(writer, []byte(str))
}
