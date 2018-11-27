package db

import (
	"database/sql"
	"fmt"
)

var db *sql.DB
func init(){
	db, _ = sql.Open("mysql","root:rootroot@/sky001?charset=utf8")
	db.Query("use sky001")
	//db.SetMaxOpenConns(2000)
	//db.SetMaxIdleConns(1000)
	//db.Ping()
}


func Exec(sqlInsert string)sql.Result{
	stmt, err := db.Prepare(sqlInsert)
	checkErr(err)
	r, err := stmt.Exec()
	fmt.Print(err)
	return r
}


//// 更新数据
//func Update(sqlUpdate string){
//	stmt, err := db.Prepare(sqlUpdate)
//	checkErr(err)
//	res, err := stmt.Exec()
//	checkErr(err)
//	num, err := res.RowsAffected()
//	checkErr(err)
//	fmt.Println("影响的行数为：", num)
//}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
