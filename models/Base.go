package models

import (
    beego "github.com/beego/beego/v2/server/web"
	"github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"

	//"fmt"
)

// Model Struct
type BaseModel struct {
    
}

func init() {
    url,_ := beego.AppConfig.String("db.url")
    port,_ := beego.AppConfig.String("db.port")
    username,_ := beego.AppConfig.String("db.username")
    password,_ := beego.AppConfig.String("db.password")

    //fmt.Println(url1, port1, username1, password1)

    orm.RegisterDataBase("default", "mysql", username+":"+password+"@tcp("+url+":"+port+")/laravel56?charset=utf8", 30)
}


