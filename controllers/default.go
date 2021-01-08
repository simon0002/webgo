package controllers

import (
	//beego "github.com/beego/beego/v2/server/web"
	"beego/models"
	"fmt"
	"time"
	//"io/ioutil"
	"os"
)

type MainController struct {
	//beego.Controller
	BaseController

}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Abc() {
	c.Data["Website"] = "beego.me.abc"
	c.Data["Email"] = "astaxieabc@gmail.com"
	c.TplName = "index.tpl"

	//this.Ctx.WriteString("hello")
}

func (c *MainController) Home() {
	c.Data["Users"] = models.List()
	c.TplName = "home.html"
}

func (c *MainController) Add() {
	//c.Data["Users"] = models.List()
	c.TplName = "add.html"
}

func (c *MainController) Save() {
	username := c.GetString("username")
	password := c.GetString("password")
	email := c.GetString("email")

	user := models.User{Username:username, Password:password, Email:email}
	models.Save(&user)

	c.Ctx.WriteString("hello")
}

func (c *MainController) Test() {
	a := 1
	
	go func (a int) {
		for i,n:=0,9; i<n; i++ {
			a := i
			//c.Ctx.WriteString(a)
			fmt.Println(a)
		} 
	}(a)

	//c.Ctx.WriteString(a)
	fmt.Println(a)
	c.Ctx.WriteString("ok")
}

func (c *MainController) TestDefer() {
	//a := 10
	
	for i,n:=0,9; i<n; i++ {
		//a := i
	} 
	
	defer c.TestGo()
	
	now := time.Now()
	c.Ctx.WriteString("ok"+string(now.Format("2006/1/2 15:04:05")))
}

func (c *MainController) TestGo() {
	go func () {
		f,_ := os.OpenFile("a.txt",os.O_RDWR |os.O_APPEND,0777)
		for i,n:=0,9; i<n; i++ {
			
			now := time.Now()

			content := "hello world!"+string(now.Format("2006/1/2 15:04:05")+"_____")
    		f.Write([]byte(content))
    		time.Sleep(time.Duration(1)*time.Second)			
		} 
		f.Close()
	}()
}
