package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"

	"bytes"
	"crypto/md5"
	"fmt"
	"io"
)

type AuthController struct {
	beego.Controller
}

//var globalSessions *session.Manager

func init() {
	//sessionConfig := &session.ManagerConfig{
	//	CookieName:      "gosessionid",
	//	EnableSetCookie: true,
	//	Gclifetime:      3600,
	//	Maxlifetime:     3600,
	//	Secure:          false,
	//	CookieLifeTime:  3600,
	//	ProviderConfig:  "111.231.215.178:6379,100,",
	//}
	//globalSessions, err := session.NewManager("redis", sessionConfig)
	//if err != nil {
	//	log.Println(err)
	//}
	//go globalSessions.GC()
}

type User struct {
	Name     string `json:"username" form:"username"`
	Password string `json:"password" form:"password,password,"`
}

func (c *AuthController) Get() {
	c.TplName = "login.html"
	c.Data["Form"] = &User{}
	c.Render()
}

const Passeord = "21232f297a57a5a743894a0e4a801fc3"

func (c *AuthController) Login() {
	userInfo := User{}
	c.ParseForm(&userInfo)
	md5Password := md5.New()
	io.WriteString(md5Password, userInfo.Password)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
	newPass := buffer.String()
	fmt.Println(newPass)
	//now := time.Now().Format("2006-01-02 15:04:05")
	if Passeord == newPass {
		//var users models.User
		//users.Last_logintime = now
		//models.UpdateUserInfo(users)

		//登录成功设置session
		c.SetSession("uname", userInfo.Name)
		//sess,err := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
		//if err != nil {
		//	log.Println("err ", err)
		//}
		//
		//sess.Set("uid", userInfo.Name)
		//sess.Set("uname", userInfo.Username)
		//c.Ctx.Output.Body([]byte("ok"))
		c.Redirect( "/index",302)
		return
	}
	c.Redirect("/", 302)
	return
}

func (c *AuthController) Index()  {
	if c.GetSession("uname") != "a" {
		c.Redirect("/", 302)
	}
}
