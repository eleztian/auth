package rediesDB

import (
	"gopkg.in/mgo.v2"
	"github.com/astaxie/beego"
)

var MDBConfig struct {
	DBUrl  string
	DBName string
	CName  string
}

var session *mgo.Session

// Conn return mongodb session.
func Conn() *mgo.Session {
	return session.Copy()
}

func Close() {
	session.Close()
}

func init() {
	MDBConfig.DBUrl = beego.AppConfig.String("mongodb::url")
	MDBConfig.DBName = beego.AppConfig.String("mongodb::name")

	sess, err := mgo.Dial(MDBConfig.DBUrl)
	if err != nil {
		panic(err)
	}

	session = sess
	session.SetMode(mgo.Monotonic, true)
}
