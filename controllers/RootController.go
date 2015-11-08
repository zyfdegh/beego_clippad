package controllers

import(
  "github.com/astaxie/beego"
)

type RootController struct{
  beego.Controller
}

func (this *RootController)Get(){
  this.TplNames="index.html"
  this.Render()
}
