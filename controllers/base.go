package controllers

import (
	"github.com/astaxie/beego"
    "time"
 
 
)
 type BaseController struct {
     beego.Controller
     is_login bool 
     user_id int 
     user_name string

 }
 // 给自己的看的加上session
 func (this *BaseController) Prepare() {
     user_id := this.GetSession("user_id")
     user_name := this.GetSession("user_name")
     method :=   this.Ctx.Request.Method
 
     if user_id == nil {
         this.is_login = false
         if method == "POST" {
             this.Data["json"] = map[string]interface{} {"code":1,"message":"请先登录"}
         }else {
             this.Redirect("/login",302) 
         }
       // this.TplName = "login.html"
         return 
     }else {
         this.is_login = true 
         this.user_id = user_id.(int)
         this.user_name = user_name.(string)
        
     }
     year := time.Now().Year() 
     this.Data["is_login"] = this.is_login
     this.Data["user_name"] = user_name 
     this.Data["year"] = year


 
	this.Data["Email"] = "astaxie@gmail.com"
 

 }