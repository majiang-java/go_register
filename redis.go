package main

import (
	"encoding/xml"
	"fmt"
)

type redisConfig struct {
	RedisList []redisnode `xml:"node" json:"node"`
}

type redisnode struct {
	Host     string `xml:"host,attr"`
	Port     string `xml:"port,attr"`
	Password string `xml:"password,attr"`
}

func getRedisClaster(clsuter string) redisConfig {
	var etree = load("/conf/redis/redis-cluster.xml")
	var lcl redisConfig
	for k, v := range etree.PersonList {
		if v.Name == clsuter {
			fmt.Println(v.Path, k)
			var body = loadHTTPResource("/" + v.Path)
			xml.Unmarshal([]byte(body), &lcl)
			covertRedis(&lcl)
			return lcl
		}
	}
	return lcl
}

func covertRedis(l *redisConfig) {
	node := l.RedisList[0]
	password := DecryptAes128Ecb(node.Password)
	l.RedisList[0].Password = string(password)
}
