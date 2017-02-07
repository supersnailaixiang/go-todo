package models

import (
	"database/sql"

)
 
func LoginAccount(data []interface{}) (bool,string,int) {
    var user_id int
    var user_name string
  //  result := make(map[string]string,3)
    check_sql := "select user_id,user_name from cfg_user where user_name = ? and pwd = md5(?) ";
    db  := db()
    row := db.prepareQueryRow(check_sql,data[:]...)
 
    err := row.Scan(&user_id,&user_name)
    if err == sql.ErrNoRows {
         
        return false,"",0
    }
     
 
    return true,user_name,user_id;     

}