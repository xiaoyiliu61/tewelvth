package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName?theCityName=地名
	urlStr := "http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName?theCityName=%E5%8D%97%E6%98%8C"
	response, err := http.Get(urlStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer response.Body.Close()

	statusCode:=response.StatusCode
	if statusCode == 200 {
			body,err:=ioutil.ReadAll(response.Body)
			if err!=nil {
				log.Fatal(err)
				return
			}
			fmt.Println(string(body))

	 fmt.Println("status:")
	 fmt.Println(response.Status)
	 header := response.Header
	 for key, value := range header {
		fmt.Println(key, value)
	}
	 request := response.Request
	 fmt.Println("method", request.Method)
	 fmt.Println("url", request.URL)
	 fmt.Println("path", &request.URL)

	 fmt.Println("cookies", response.Cookies())
     }else{
	  fmt.Println("此次请求不正常，请求状态码",response.StatusCode)
 }
}
