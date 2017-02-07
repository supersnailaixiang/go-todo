package init
import (
    "ToDo/common"
    "github.com/astaxie/beego"
)
func init(){
    beego.AddFuncMap("to_do_status",common.ToDoStatus)
}