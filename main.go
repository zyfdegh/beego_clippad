package main

import(
  "github.com/astaxie/beego"
  // "github.com/zyfdegh/beego_clippad/controllers"
  "github.com/zyfdegh/beego_clippad/router"
)

func main(){
  // var s=Server.NewServer()
  // s.Config=Config.Load(f)
  // s.Serve()
  router.RouteAll()
  beego.Run()
}
