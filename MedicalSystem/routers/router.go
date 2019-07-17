package routers

import (
	"MedicalSystem/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    //注册
	beego.Router("/register/:name/:Password", &controllers.MainController{},"get:ShowRegister;post:HandleRegister")
	beego.Router("/login/:name/:Password/:loginType", &controllers.MainController{},"get:ShowLogin;post:HandleLogin")
}
