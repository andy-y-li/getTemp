package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/liamylian/values"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strings"
)

//模拟curl命令发送http GET|POST|PUT|DELETE 请求
func main() {
	url := flag.String("url", "http://localhost:8080/api/cpu/T", "request url")
	method := flag.String("method", "GET|POST|PUT|DELETE", "request method")
	reqBody := flag.String("reqBody", "{a:x}", "json data")
	flag.Parse()

	if strings.EqualFold(*method, "GET") {
		get(*url)
		return
	}
	if strings.EqualFold(*method, "POST") {
		post(*url, "application/json", *reqBody)
		return
	}
	if strings.EqualFold(*method, "PUT") {
		put(*url, *reqBody)
		return
	}
	if strings.EqualFold(*method, "DELETE") {
		delete(*url, *reqBody)
		return
	}
}

func get(url string) {
	//fmt.Println("GET REQ...")
	client := http.Client{}
	rsp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
	}

	result := gjson.ParseBytes(body)

	m := map[string]values.Value{"foo": "bar"}
	vs := values.FromMap(m)

	fmt.Printf("Time,Temperature,Fan\n")
	for _, obj := range result.Array() {
		//	println(obj.String())
		json.Unmarshal([]byte(obj.String()), &vs)

		//idNum, _ := vs.Get("id")
		td, _ := vs.Get("tdatetime")
		t, _ := vs.Get("temperature")
		fan, _ := vs.Get("fan")
		fmt.Printf("%s,%-4s,%s\n", td, t, fan)
	}

}

func post(url string, contentType string, reqBody string) {
	fmt.Println("POST REQ...")
	fmt.Println("REQ:", reqBody)
	client := http.Client{}
	rsp, err := client.Post(url, contentType, strings.NewReader(reqBody))
	if err != nil {
		fmt.Println(err)
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("RSP:", string(body))
}

func put(url string, reqBody string) {
	fmt.Println("PUT REQ...")
	fmt.Println("REQ:", reqBody)
	req, err := http.NewRequest("PUT", url, strings.NewReader(reqBody))
	if err != nil {
		fmt.Println(err)
	}

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("RSP:", string(body))
}

func delete(url string, reqBody string) {
	fmt.Println("DELETE REQ...")
	fmt.Println("REQ:", reqBody)
	req, err := http.NewRequest("DELETE", url, strings.NewReader(reqBody))
	if err != nil {
		fmt.Println(err)
	}

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("RSP:", string(body))
}
