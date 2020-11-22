package main

import (
	"fmt"
	"net/http"
)

func main() {
	client:=http.DefaultClient

	request,_:=http.NewRequest("GET","http://www.baidu.com",nil)
   resp,_:=client.Do(request)

   fmt.Println(resp.StatusCode)
}
