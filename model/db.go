package model

import (
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
)

var db = NewDBConn()

func NewDBConn() *gorm.DB {
	db, err := gorm.Open(GetDBConfig("dbconfig.yml","development"))
	if err != nil {
		panic(err)
	}
	return db
}

func GetDBConn() *gorm.DB {
	return db
}

func GetDBConfig(configPath string, dbname string) (string, string) {
	var buf, err = ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		panic(err)
	}
	driver := m[dbname].(map[interface{}]interface{})["dialect"].(string)
	source := m[dbname].(map[interface{}]interface{})["datasource"].(string)
	return driver, source
}