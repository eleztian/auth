package models

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

type UserLogin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password,password,"`
}

type UserInfo struct {
	Id       string
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password,password,"`
	Salt     string `json:"salt" form:"sale,"`
	Info     Profile
}

type Profile struct {
	Email string
}

func AddUser(u UserInfo) string {
	
	return u.Id
}

func GetUser(userame string) (u *UserInfo, err error) {
	redis.Dial()
	return nil, nil
}

func GetAllUsers() map[string]*UserInfo {
	return nil
}

func UpdateUser(username string, uu *UserInfo) (a *UserInfo, err error) {
	return nil, nil
}

func Login(username, password string) bool {
	return false
}

func DeleteUser(username string) {
}
