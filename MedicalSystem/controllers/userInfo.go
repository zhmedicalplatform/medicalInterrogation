package controllers

import (
	"MedicalSystem/models"
	"MedicalSystem/utility"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index.tpl"
}

//注册
func (this *MainController) ShowRegister() {
	_nsuccess := 0
	_strInfo:= "注册成功"

	//1.拿到数据
	userName := this.GetString("userName")
	pwd := this.GetString("pwd")

	_name := utility.RSA_Decrypt([]byte(userName),"Medical_System_Private.pem")
	_password := utility.RSA_Decrypt([]byte(pwd),"Medical_System_Private.pem")
	userName = string(_name)
	pwd = string(_password)

	//2.对数据进行校验
	if userName == "" || pwd == ""{
		beego.Info("数据不能为空")
	}

	//3.查询账号是否存在
	o := orm.NewOrm()
	user := models.AccountsInfo{}

	user.Name = userName
	user.Password = pwd
	err := o.Read(&user,"Name")
	if err == nil {
		_nsuccess = 1
		_strInfo = "当前名字已经注册"
		beego.Info("当前名字已经注册")
	}

	user.Password = pwd
	_,err_Insert := o.Insert(&user)
	if err_Insert != nil{
		_nsuccess = 2
		_strInfo = "注册失败"
		beego.Info("注册失败")
	}

	this.Data["json"] = map[string]interface{}{"success": _nsuccess, "info": _strInfo}
	this.ServeJSON()
}

func (this *MainController) HandleRegister() {

	beego.Info("点击注册")
	//1.拿到数据
	userName := this.GetString("userName")
	pwd := this.GetString("pwd")
	//2.对数据进行校验
	if userName == "" || pwd == ""{
		beego.Info("数据不能为空")
		this.Redirect("/register",302)
		return
	}
	//3.插入数据库
	o := orm.NewOrm()

	user := models.AccountsInfo{}
	user.Name = userName
	user.Password = pwd
	_,err := o.Insert(&user)
	if err != nil{
		beego.Info("插入数据失败")
		this.Redirect("/register",302)
		return
	}
	//4.返回登陆界面
	this.Redirect("/login",302)
}

//登陆
func (this *MainController) ShowLogin() {
	_nsuccess := 0
	_strInfo:= "登陆成功"
	//加密
	name:= this.GetString(":name")
	Password := this.GetString(":Password")
	//beego.Info("handleLogin end name ",name," password  ",Password)

	o := orm.NewOrm()
	user := models.AccountsInfo{}

	_name := utility.RSA_Decrypt([]byte(name),"Medical_System_Private.pem")
	_password := utility.RSA_Decrypt([]byte(Password),"Medical_System_Private.pem")
	user.Name = string(_name)
	user.Password = string(_password)
	err := o.Read(&user,"Name")
	if err != nil {
		beego.Info("没有查询到当前用户")
		_nsuccess = 1
		_strInfo = "没有查询到当前用户"
	}

	this.Data["json"] = map[string]interface{}{"success": _nsuccess, "info": _strInfo}
	this.ServeJSON()

}

func (this *MainController) HandleLogin() {
	beego.Info("点击登陆")
	//1.拿到数据
	userName := this.GetString("userName")
	pwd := this.GetString("pwd")
	//2.判断数据是否合法
	if userName == ""|| pwd ==""{
		beego.Info("输入数据不合法")
		this.TplName = "login.html"
		return
	}
	//3.查询账号密码是否正确
	o := orm.NewOrm()
	user := models.AccountsInfo{}

	user.Name = userName
	err := o.Read(&user,"Name")
	if err != nil {
		beego.Info("查询失败")
		this.TplName = "login.html"
		return
	}
	//4.跳转
	this.Ctx.WriteString("登陆成功")
	this.Data["json"] = map[string]interface{}{"success": 0, "message": "111"}
	this.ServeJSON()
}

