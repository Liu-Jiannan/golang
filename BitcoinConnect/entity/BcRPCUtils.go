package entity

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

/*
1.准备一个json数据
2.发送一个post请求
 */

type Prepare struct {
	id int64
	jsonrpc float32
	method string
	params []byte
}

func PrepareJSON(method string,params ... byte) []byte {
	object := Prepare{
		id:      time.Now().UnixNano(),
		jsonrpc: 2.0,
		method:  method,
	}
	if params != nil {
		object.params = params
	}
	perpare_json,error:= json.Marshal(object)
	if error!=nil {
		fmt.Println("json数据准备失败")
	}
	return perpare_json
}


/*
发送一个post请求
 */

func SendPost(headers map[string]string,jsonStr []byte)  {
	client := &http.Client{}

	url := "http://127.0.0.1:8332"




}
