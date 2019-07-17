package main

import (
	_ "MedicalSystem/routers"
	_"MedicalSystem/models"
	_"MedicalSystem/utility"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

