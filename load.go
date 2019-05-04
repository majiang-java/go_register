package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func initc(env string) interface{} {
	var jsonfile, err = os.Open("server_config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonfile.Close()
	byteValue, _ := ioutil.ReadAll(jsonfile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	serverConfig := result[env].(map[string]interface{})
	fmt.Println(serverConfig["url"])
	return serverConfig
}

type configuration struct {
	PersonList []cluster `xml:"cluster" json:"cluster"`
}

type cluster struct {
	Name string `xml:"name,attr"`
	Path string `xml:"path,attr"`
}

func loadHTTPResource(url string) []byte {
	fmt.Println("string:" + url)
	var serverConfigd = initc("default")
	serverConfig, ok := serverConfigd.(map[string]interface{})
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}
	req, err := http.NewRequest("GET", fmt.Sprint(serverConfig["url"])+url, nil)
	if err != nil {
		fmt.Println(err)
	}
	if ok != true {
		fmt.Println(serverConfig)
	}
	req.Header.Add("X-INFRA-TOKEN", fmt.Sprint(serverConfig["token"]))
	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)

	return body
}

func load(resource string) configuration {
	var body = loadHTTPResource(resource)
	fmt.Println("str" + string(body))
	var ll configuration
	xml.Unmarshal([]byte(body), &ll)
	return ll
}

func main() {
	//load()
	// var etree = load()
	// for k, v := range etree.PersonList {
	// 	if v.Name == "common-test" {
	// 		fmt.Println(v.Path, k)
	// 		var body = loadHTTPResource("/" + v.Path)
	// 		fmt.Println(string(body))
	// 	}
	// }
	// mes := getMysqlClaster("common-test")
	// fmt.Println(mes)
	// mes := getRedisClaster("common-test")
	// fmt.Println(mes)
	mes := getKafkaClaster("common-test")
	fmt.Println(mes)
	//loadMysql()
	//encryptMsg, _ := encrypt("common!next")
	// msg := DecryptAes128Ecb("3zlWX8Pc4gQI1sX0YSwpZg==")
	// fmt.Println(string(msg)) // Hello World
}
