package main

import (
	"fmt"
	"net/http"
)

func main() {
	client:=http.Client{}

	response,_:=client.Get("http://www.baidu.com")

	fmt.Println(response.StatusCode)
}
