package controllers 

type UntilsController struct {
    BaseController
}
func (this * UntilsController) handleErr(err error){
    this.Data["result"] = err
    this.TplName = "error.html"
}