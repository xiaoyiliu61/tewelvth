package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client:=http.Client{}

	request,err:=http.NewRequest("GET","http://www.baidu.com",nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	reponse,err:=client.Do(request)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer reponse.Body.Close()

	if reponse.StatusCode==200{
		body,_:=ioutil.ReadAll(reponse.Body)
		fmt.Println(string(body))
	}
}
