package controllers

import (
	//beego "github.com/beego/beego/v2/server/web"
	"beego/models"
	"fmt"
	"time"
	//"io/ioutil"
	"os"
	//"github.com/astaxie/beego/context"
	"beego/utils"
)

type UserController struct {
	//beego.Controller
	BaseController

}

func (c *UserController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *UserController) Abc() {
	// c.Data["Website"] = "beego.me.abc"
	// c.Data["Email"] = "astaxieabc@gmail.com"
	// c.TplName = "index.tpl"

	//this.Ctx.WriteString("hello")

	//返回json
	users := models.UserList(0, 10)
	c.Data["json"] = utils.Response{Code:0, Data:users, Msg:""}

	c.ServeJSON()
}

func (c *UserController) Home() {
	pnum,err := c.GetInt("p")
	if(err != nil){
		pnum = 1
	}

	//列表分页
	psize := 1
	offset := (pnum-1)*psize

	count := models.UserCount()
	c.Data["paginator"] = utils.NewPaginator(c.Ctx.Request, psize, count)
	c.Data["Users"] = models.UserList(offset, psize)
	c.TplName = "home.html"
}

func (c *UserController) Add() {

	c.TplName = "add.html"
}

func (c *UserController) Save() {
	username := c.GetString("username")
	password := c.GetString("password")
	email := c.GetString("email")

	user := models.User{Username:username, Password:password, Email:email}
	models.UserSave(&user)

	c.Ctx.WriteString("hello")
}


func (c *UserController) Test() {
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

func (c *UserController) TestDefer() {
	//a := 10
	
	for i,n:=0,9; i<n; i++ {
		//a := i
	} 
	
	defer c.TestGo()
	
	now := time.Now()
	c.Ctx.WriteString("ok"+string(now.Format("2006/1/2 15:04:05")))
}

func (c *UserController) TestGo() {
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

