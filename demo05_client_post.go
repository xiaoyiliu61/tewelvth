package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	client:=http.Client{}

	urlStr:="http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName"
	params:=&url.Values{
		"theCityName":{"南昌"},
	}
	bufferstring:=bytes.NewBufferString(params.Encode())
	response,err:=client.Post(urlStr,"application/x-www-form-urlencoded",bufferstring)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer response.Body.Close()
	fmt.Println(response.StatusCode  )
}
