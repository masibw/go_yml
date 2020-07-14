package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Containers []struct {
	Container struct {
		Id         string `yaml:id`
		Name       string `yaml:name`
		Type string `yaml:type`
		Allow      struct {
			Process []struct {
				Name string `yaml:name`
				Path string `yaml:path`
				Exec string `yaml:exec`
			} `yaml:process`
			Socket []struct {
				Protocol string `yaml:protocol`
				Ip       string `yaml:ip`
				Port     int    `yaml:port`
			} `yaml:socket`
		} `yaml:allow`
	} `yaml:container`
}

func main() {
	//var m map[interface{}]interface{}
	buf, err := ioutil.ReadFile("main/test.yml")
	if err != nil {
		log.Fatalf("error:%v", err)
	}
	containers := Containers{}
	err = yaml.Unmarshal([]byte(buf), &containers)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//mapがkeyを持っていればifの中に入る
	fmt.Println(containers[0])
	fmt.Println(containers[0].Container.Allow.Process[1].Name)
}
