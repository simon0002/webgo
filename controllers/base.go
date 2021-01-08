package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"fmt"
)

// baseRouter implemented global settings for all other routers.
type BaseController struct {
        beego.Controller
       // i18n.Locale
       // user    models.User
       // isLogin bool
}

// Prepare implemented Prepare method for baseRouter.
func (this *BaseController) Prepare() {

        // page start time
        // this.Data["PageStartTime"] = time.Now()

        // // Setting properties.
        // this.Data["AppDescription"] = utils.AppDescription
        // this.Data["AppKeywords"] = utils.AppKeywords
        // this.Data["AppName"] = utils.AppName
        // this.Data["AppVer"] = utils.AppVer
        // this.Data["AppUrl"] = utils.AppUrl
        // this.Data["AppLogo"] = utils.AppLogo
        // this.Data["AvatarURL"] = utils.AvatarURL
        // this.Data["IsProMode"] = utils.IsProMode

		//now := time.Now()

		fmt.Println("now")
        
}