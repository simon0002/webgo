package controllers

import (
	//beego "github.com/beego/beego/v2/server/web"
	"beego/models"
	"fmt"
	"time"
	//"io/ioutil"
	"os"
	"net/http"
	//"github.com/astaxie/beego/context"
	"beego/utils"
	 "github.com/astaxie/beego/validation"
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

	//表单验证
	valid := validation.Validation{}
    valid.Required(username, "username").Message("请输入账号")
    valid.Required(password, "password").Message("请输入密码")
    valid.MinSize(email, 5, "email").Message("email太短")

    if valid.HasErrors() {
        // 如果有错误信息，证明验证没通过
        message := ""
        for _, err := range valid.Errors {
            fmt.Println(err.Key, err.Message)
            message += err.Message + "__"
        }
		c.Data["json"] = utils.Response{Code:1002, Data:nil, Msg:message}
		c.ServeJSON()

    }

	user := models.User{Username:username, Password:password, Email:email}
	models.UserSave(&user)

	c.Ctx.WriteString("hello")
}

func (c *UserController) UpdateAvatar() {
	f, h, err := c.GetFile("avatar");
	if err == http.ErrMissingFile {
		c.Data["json"] = utils.Response{Code:1002, Data:nil, Msg:"请选择文件"}
		c.ServeJSON()
	}

	defer f.Close()

	if err != nil {
		c.Data["json"] = utils.Response{Code:1002, Data:nil, Msg:"上传失败"}
		c.ServeJSON()
	} else {
		c.SaveToFile("avatar", "upload/"+h.Filename)
 		//var user models.User = c.GetSession("user")
 		//fmt.Println(user)

		// user.Avatar = "upload/" + h.Filename
        // models.UserUpdate(&user)

    	c.Data["json"] = utils.Response{Code:1002, Data:nil, Msg:"上传成功"}
		c.ServeJSON()
	}
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

