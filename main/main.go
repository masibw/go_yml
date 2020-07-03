package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)
// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func main() {
	var m map[interface{}]interface{}
	buf, err := ioutil.ReadFile("main/test.yml")
	if err !=nil{
		log.Fatalf("error:%v",err)
	}
	err = yaml.Unmarshal([]byte(buf), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)
}