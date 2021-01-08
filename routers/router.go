package routers

import (
	"beego/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/abc", &controllers.MainController{}, "get:Abc")

    beego.Router("/home", &controllers.MainController{}, "get:Home")
    beego.Router("/add", &controllers.MainController{}, "get:Add")
    beego.Router("/save", &controllers.MainController{}, "post:Save")


    beego.Router("/test", &controllers.MainController{}, "get:Test")
    beego.Router("/testDefer", &controllers.MainController{}, "get:TestDefer")
}
