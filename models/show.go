package models

import "fmt"

type ToDoThings struct {
	RecID         int
	Status        int
	ToDoThings    string
	UserName      string
	EndDate       string
	ExpectEndDate string
}

func ShowList(condArr map[string]string) []ToDoThings {
	status := condArr["status"]
	end_date := condArr["end_date"]
	begin_date := condArr["begin_date"]

	data := []interface{}{}
	sql := "select rec_id,status,expect_end_date, end_date,to_do_things,user_name from to_do_list tdl left join cfg_user cu on cu.user_id = tdl.user_id where true "

	if status != "-1" && status != "" {
		sql += " And tdl.status = ? "
		data = append(data, status)
	}
	if end_date != "" {
		sql += " And tdl.end_date <= ? "
		data = append(data, end_date)
	}
	if begin_date != "" {
		sql += " And tdl.end_date >= ? "
		data = append(data, begin_date)
	}
	sql += " order by rec_id desc "
	db := db()
	var tds []ToDoThings
	fmt.Println(sql)
	rows := db.prepareQuery(sql, data[:]...)
	for rows.Next() {
		var td ToDoThings
		err := rows.Scan(&td.RecID, &td.Status, &td.ExpectEndDate, &td.EndDate, &td.ToDoThings, &td.UserName)
		checkErr(err)

		tds = append(tds, td)
	}
	return tds

}
func MAddThings(data []interface{}) map[string]interface{} {
	sql := "insert into to_do_list(status,expect_end_date,to_do_things,user_id,created)values(0,?,?,?,now())"
	db := db()

	db.prepareExecute(sql, data...)
	return map[string]interface{}{"code": "0", "message": "添加成功"}
}
func MGetThings(rec_id int) ToDoThings {
	sql := "select rec_id,expect_end_date,to_do_things from to_do_list where rec_id = ?"
	data := []interface{}{rec_id}
	db := db()
	row := db.prepareQueryRow(sql, data...)
	var td ToDoThings
	_ = row.Scan(&td.RecID, &td.ExpectEndDate, &td.ToDoThings)

	return td
}
func MChgThings(data []interface{}) map[string]interface{} {
	sql := "update to_do_list set expect_end_date=?,to_do_things = ? where rec_id = ? "
	db := db()
	result := db.prepareExecute(sql, data...)
	if true == result {
		return map[string]interface{}{"code": "0", "message": "修改成功"}
	} else {
		return map[string]interface{}{"code": "0", "message": "修改错误"}
	}
}
func MEndThings(data []interface{}) map[string]interface{} {
	sql := "update to_do_list set end_date = ?,status=10 where rec_id = ?"
	db := db()

	result := db.prepareExecute(sql, data...)
	if true == result {
		return map[string]interface{}{"code": "0", "message": "恭喜你圆满完成计划"}
	} else {
		return map[string]interface{}{"code": "0", "message": "系统错误"}
	}
}
func MDropThings(data []interface{}) map[string]interface{} {
	sql := "update to_do_list set end_date = ?,status=5 where rec_id = ?"
	db := db()

	result := db.prepareExecute(sql, data...)
	if true == result {
		return map[string]interface{}{"code": "0", "message": "你放弃了你的计划。。。"}
	} else {
		return map[string]interface{}{"code": "0", "message": "系统错误"}
	}
}
