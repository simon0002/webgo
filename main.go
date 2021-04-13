package main

import (
	_ "beego/crontab"
	_ "beego/routers"

	"github.com/beego/beego/v2/adapter/toolbox"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true

	// 定时任务
	toolbox.StartTask()
	defer toolbox.StopTask()

	beego.Run()
}

// var mystruct map[string]interface{}
// json.Unmarshal(this.Ctx.Input.RequestBody, &mystruct)
