package routers

import (
	"beego/controllers"
    //"beego/filters"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    //beego.InsertFilter("/*",beego.BeforeRouter,filters.FilterUser)

    beego.Router("/", &controllers.UserController{})

    beego.Router("/login", &controllers.LoginController{}, "get:Login")
    beego.Router("/dologin", &controllers.LoginController{}, "post:DoLogin")

    beego.Router("/abc", &controllers.UserController{}, "get:Abc")

    beego.Router("/home", &controllers.UserController{}, "get:Home")
    beego.Router("/add", &controllers.UserController{}, "get:Add")
    beego.Router("/save", &controllers.UserController{}, "post:Save")
    beego.Router("/avatar", &controllers.UserController{}, "post:UpdateAvatar")


    beego.Router("/test", &controllers.UserController{}, "get:Test")
    beego.Router("/testDefer", &controllers.UserController{}, "get:TestDefer")
}
