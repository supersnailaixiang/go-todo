package routers

import (
	"ToDo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ShowController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/show/add", &controllers.ShowController{}, "*:AddThings")
	beego.Router("/show/get", &controllers.ShowController{}, "*:GetThings")
	beego.Router("/show/chg", &controllers.ShowController{}, "*:ChgThings")
	beego.Router("/show/end", &controllers.ShowController{}, "*:EndThings")
	beego.Router("/show/drop", &controllers.ShowController{}, "*:DropThings")
}
