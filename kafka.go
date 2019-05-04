package main

import (
	"encoding/xml"
	"fmt"
)

type kafkaConfig struct {
	KakfaList []kafkanode `xml:"property" json:"node"`
}
type kafkanode struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

func getKafkaClaster(clsuter string) kafkaConfig {
	var etree = load("/conf/kafka/kafka-cluster.xml")
	var lcll kafkaConfig
	for k, v := range etree.PersonList {
		// var body = loadHTTPResource("/" + v.Path)
		// fmt.Println(v.Path, k)
		// fmt.Println(string(body))
		// var lcl kafkaConfig
		// xml.Unmarshal([]byte(body), &lcl)
		// fmt.Println(lcl)
		if v.Name == clsuter {
			fmt.Println(v.Path, k)
			var body = loadHTTPResource("/" + v.Path)
			xml.Unmarshal([]byte(body), &lcll)
			return lcll
		}
	}
	return lcll
}
