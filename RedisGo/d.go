/**  
* Date: 2017-10-09 
* Time: 11:19 
* Description:
*/
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/test?charset=utf8")
	defer db.Close()
	checkErr(err)
	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username = ?,departname = ?, created = ? ")
	checkErr(err)
	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("UPDATE userinfo SET username = ? WHERE uid = ?")
	checkErr(err)
	res, err = stmt.Exec("lvtanxi", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	var uid int
	var username, department, created string
	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department, created)
	}

	//删除数据
	stmt,err=db.Prepare("DELETE FROM userinfo WHERE uid =? ")
	checkErr(err)
	res,err=stmt.Exec(id)
	checkErr(err)
	affect,err =res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
