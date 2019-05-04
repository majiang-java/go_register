package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlConfig struct {
	PersonList []node `xml:"node" json:"node"`
}

type node struct {
	Name     string `xml:"name,attr"`
	Host     string `xml:"host,attr"`
	Port     string `xml:"port,attr"`
	Password string `xml:"password,attr"`
}

func getConnection() *sql.DB {
	db, err := sql.Open("mysql", "username:password@tcp(118.24.3.198:3306)/test")
	if err != nil {
		fmt.Println("err happend")
	}
	defer db.Close()
	return db
}

func getMysqlClaster(clsuter string) mysqlConfig {
	var etree = load("/conf/mysql/mysql-cluster.xml")
	var lcl mysqlConfig
	for k, v := range etree.PersonList {
		if v.Name == clsuter {
			fmt.Println(v.Path, k)
			var body = loadHTTPResource("/" + v.Path)
			xml.Unmarshal([]byte(body), &lcl)
			covert(&lcl)
			return lcl
		}
	}
	return lcl
}

func covert(l *mysqlConfig) {
	node := l.PersonList[0]
	password := DecryptAes128Ecb(node.Password)
	l.PersonList[0].Password = string(password)
}

func loadMysql() {
	var db = getConnection()
	rows, err := db.Query("select id, name from test.location limit 5")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(id, name)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
