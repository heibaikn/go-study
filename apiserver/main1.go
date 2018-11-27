package main

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

//var db *sql.DB
func main() {
	//abc()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
	//db, err := sql.Open("mysql","root:rootroot@/sky001?charset=utf8")
	//checkErr(err)
	//db.Query("use sky001")
 	//res,_ := db.Query("select * from USER")
	//db.Query("drop database if exists tmpdb")
	//db.Query("create database tmpdb")

	//queryByUserName := fmt.Sprintf("'%v'", "shenzhen001")
	//res := db.Query('select * from user')
	//v := reflect.ValueOf(res)
	//fmt.Println(v)
	//fmt.Println("--增加数据测试--")
	//ids := []int{1, 2, 3}
	//p.Raw("SELECT name FROM user WHERE id IN (?, ?, ?)", ids)
	//printResult(res)
	//fmt.Print(res)
}
func printResult(query *sql.Rows) {
	column, _ := query.Columns()              //读出查询出的列字段名
	values := make([][]byte, len(column))     //values是每个列的值，这里获取到byte里
	scans := make([]interface{}, len(column)) //因为每次查询出来的列是不定长的，用len(column)定住当次查询的长度
	for i := range values {                   //让每一行数据都填充到[][]byte里面
		scans[i] = &values[i]
	}
	results := make(map[int]map[string]string) //最后得到的map
	i := 0
	for query.Next() { //循环，让游标往下移动
		if err := query.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			fmt.Println(err)
			return
		}
		row := make(map[string]string) //每行数据
		for k, v := range values {     //每行数据是放在values里面，现在把它挪到row里
			key := column[k]
			row[key] = string(v)
		}
		results[i] = row //装入结果集中
		i++
	}
	for k, v := range results { //查询出来的数组
		fmt.Println(k, v)
	}
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}



//func abc(){
//	num1 , num2 , num3 := 1 , 1 , 1
//	for  ; num1 <= 9 ;  {
//		if num2 <= num1 {
//			num3 = num1 * num2
//			fmt.Print("  ", num1, "*", num2, "=", num3)
//			num2++
//		}else {
//			num2 = 1
//			num1++
//			fmt.Print("\n")
//		}
//	}
//}