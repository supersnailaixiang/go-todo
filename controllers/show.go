package controllers

import . "ToDo/models"

type ShowController struct {
	BaseController
}

func (this *ShowController) Get() {
	status := this.GetString("status")
	end_date := this.GetString("end_date")
	begin_date := this.GetString("begin_date")
	/*
	       if err != nil {
	           uc := new(UntilsController)
	           uc.handleErr(err)
	   	}
	*/
	condArr := make(map[string]string)
	condArr["status"] = status
	condArr["end_date"] = end_date
	condArr["begin_date"] = begin_date

	result := ShowList(condArr)
	this.Data["result"] = result

	this.TplName = "show.html"
	//   this.TplName = "index.tpl"
}
func (this *ShowController) AddThings() {
	to_do_things := this.GetString("to_do_things")
	expect_end_date := this.GetString("expect_end_date")
	user_id := this.GetSession("user_id")
	if "" == to_do_things {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "事情内容不能为空"}
		this.ServeJSON()
		return
	}
	if "" == expect_end_date {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "事情没有截止时间不是个好习惯奥"}
		this.ServeJSON()
		return
	}
	data := []interface{}{expect_end_date, to_do_things, user_id}
	result := MAddThings(data)
	this.Data["json"] = result
	this.ServeJSON()
	return
}
func (this *ShowController) GetThings() {
	rec_id, _ := this.GetInt("rec_id")

	result := MGetThings(rec_id)

	this.Data["json"] = result
	this.ServeJSON()
	return
}
func (this *ShowController) ChgThings() {
	rec_id, _ := this.GetInt("rec_id")
	expect_end_date := this.GetString("expect_end_date")
	to_do_things := this.GetString("to_do_things")
	if "" == expect_end_date {
		this.Data["json"] = map[string]interface{}{"code": "1", "message": "期望截止时间不能为空"}
		this.ServeJSON()
		return
	}
	if "" == to_do_things {
		this.Data["json"] = map[string]interface{}{"code": "1", "message": "事情内容不能为空"}
		this.ServeJSON()
		return
	}
	if 0 == rec_id {
		this.Data["json"] = map[string]interface{}{"code": "1", "message": "选中修改的事情"}
	}
	data := []interface{}{expect_end_date, to_do_things, rec_id}
	result := MChgThings(data)
	this.Data["json"] = result
	this.ServeJSON()

}
func (this *ShowController) EndThings() {
	end_date := this.GetString("end_date")
	rec_id, _ := this.GetInt("rec_id")
	if "" == end_date {
		this.Data["json"] = map[string]interface{}{"code": "1", "message": "完成时间不能为空"}
		this.ServeJSON()
		return
	}
	if 0 == rec_id {
		this.Data["json"] = map[string]interface{}{"code": "1", "message": "选中完成的事情"}
	}
	data := []interface{}{end_date, rec_id}

	result := MEndThings(data)
	this.Data["json"] = result
	this.ServeJSON()

}

func (this *ShowController) DropThings() {
	end_date := this.GetString("end_date")
	rec_id, _ := this.GetInt("rec_id")
	if "" == end_date {
		this.Data["json"] = map[string]interface{}{"code": "1", "message": "放弃时间不能为空"}
		this.ServeJSON()
		return
	}
	if 0 == rec_id {
		this.Data["json"] = map[string]interface{}{"code": "1", "message": "选中放弃的事情"}
	}
	data := []interface{}{end_date, rec_id}

	result := MDropThings(data)
	this.Data["json"] = result
	this.ServeJSON()

}
