package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"beego/models"
	"fmt"
	//"io/ioutil"
	//"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
	//BaseController

}

func (c *LoginController) Login() {
	c.TplName = "login.html"
}

func (c *LoginController) DoLogin() {
	username := c.GetString("username")
	password := c.GetString("password")

	//user := models.User{Username:username, Password:password}
	user, err := models.UserGetByName(username)
	//fmt.Println(user)
	if(err != nil){
		fmt.Println(err)
		c.Ctx.WriteString("username is wrong")
		return 
	}

	if(user.Password != password){
		c.Ctx.WriteString("password is wrong")
		return
	}

	//登陆成功
	c.SetSession("user", user.Id)
	//session := c.GetSession("user")
	//fmt.Println(session)
	c.Ctx.WriteString("ok")
}




