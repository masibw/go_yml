package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func main() {
	var m map[interface{}]interface{}
	buf, err := ioutil.ReadFile("main/test.yml")
	if err !=nil{
		log.Fatalf("error:%v",err)
	}

	//mに読み込んだyamlをマッピングする
	err = yaml.Unmarshal([]byte(buf), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//mapがkeyを持っていればifの中に入る
	if content,ok:=m["container"].(map[interface{}]interface{})["allow"].([]interface{})[0].(map[interface{}]interface{})["process"].(map[interface{}]interface{})["name"]; ok{
		fmt.Print(content)
	}


}