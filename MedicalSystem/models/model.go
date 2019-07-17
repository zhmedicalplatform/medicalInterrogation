package models

import (
	_"MedicalSystem/utility"
	"github.com/astaxie/beego/orm"
	_"MedicalSystem/utility"
	_ "github.com/lib/pq"
)


//表的设计
type AccountsInfo struct {
	Id int
	Name string `orm:"unique"`	//
	Password string				//密码6-18
	LoginType int				//登陆类型(1/手机登陆 2/网页登陆)
	Code int					//验证码
}




func init(){
	//utility.GenerateRSAKey(1024)
	orm.RegisterDriver("postgres", orm.DRPostgres) // 注册驱动
	// 设置数据库基本信息
	orm.RegisterDataBase("default", "postgres", "user=postgres password=123456 dbname=accounts host=127.0.0.1 port=5432 sslmode=disable")
	// 映射model数据
	orm.RegisterModel(new(AccountsInfo))
	// 生成表
	orm.RunSyncdb("default", false, true)
}

