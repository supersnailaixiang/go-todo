package controllers 

import ."ToDo/models"
import "fmt"
 
type LoginController struct {
   BaseController
}
func(this *LoginController)Prepare(){
    
}
func (this *LoginController)Get(){
    user_id := this.GetSession("user_id")
   fmt.Println(user_id)
    if user_id != nil {
        this.Redirect("/",302)
    }else{
       this.TplName = "login.html"
    }
}
func (this *LoginController)Post(){
  
  
    user_name := this.GetString("user_name")
    pwd := this.GetString("pwd")

    data := make([]interface{},2)
    data[0] = user_name
    data[1] = pwd 
  
    code,user_name,user_id := LoginAccount(data)
    if code == true {
        this.SetSession("user_name",user_name)
        this.SetSession("user_id",user_id)
        this.Redirect("/",302)
    }else {
        this.Data["result"] = "账号或者密码错误"
        this.TplName = "error.html"
    }

}