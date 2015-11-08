package router

import(
  "github.com/astaxie/beego"
  "github.com/zyfdegh/beego_clippad/controllers"
)

func RouteAll(){
  beego.Router("/",&controllers.RootController{})
  // beego.Router("/*",&controllers.ParseController{})
}
