package models

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"time"
	"fmt"
)

type MongoDBLogger struct {
	session      *mgo.Session `json:"-"`
	Url          string       `json:"url"`
	DbName       string       `json:"db_name"`
	PreColleName string       `json:"colle_name"`
}

type logMsg struct {
	//Level int    `json:"level" bson:"level"`
	Msg   string `json:"msg" bson:"msg"`
	//When  int64  `json:"when" bson:"time"`
}

func (ml *MongoDBLogger) Init(config string) error {
	if config != "" {
		if err := json.Unmarshal([]byte(config), ml); err != nil {
			return err
		}
	}
	if ml.Url == "" {
		ml.Url = beego.AppConfig.String("mongodb::url")
	}
	if ml.DbName == "" {
		ml.DbName = beego.AppConfig.String("mongodb::dbName")
	}
	if ml.PreColleName == "" {
		ml.PreColleName = beego.AppConfig.String("log::preLogFileName")
	}

	s, err := mgo.Dial(ml.Url)
	if err != nil {
		return err
	}
	s.SetMode(mgo.Monotonic, true)
	ml.session = s
	return nil
}

func (ml *MongoDBLogger) WriteMsg(when time.Time, msg string, level int) error {
	if ml.session == nil {
		return fmt.Errorf("error not connect to host")
	}
	//start := time.Now()
	conn := ml.session.Copy()
	defer conn.Close()
	c := conn.DB(ml.DbName).C(ml.PreColleName + time.Now().Format("_2006_01_02"))
	err := c.Insert(&logMsg{msg})
	//fmt.Println(time.Since(start).String())
	return err
}

func (ml *MongoDBLogger) Destroy() {
	if ml.session != nil {
		ml.session.Close()
	}
}

func (ml *MongoDBLogger) Flush() {

}
