package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

func init() {
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
}

func insert() {
	stmt, err := db.Prepare("INSERT user_info SET username=?,departname=?,create_time=?")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec("zhangsan", "技术部", "2016-12-09")
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
}

func update() {
	stmt, err := db.Prepare("UPDATE user_info set username=? WHERE id=?")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec("王建鹏", 1)
	if err != nil {
		panic(err)
	}

	id, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println(id)

}

func query() {
	rows, err := db.Query("SELECT * FROM user_info")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var (
			id         int
			username   string
			department string
			createTime string
		)
		err = rows.Scan(&id, &username, &department, &createTime)
		fmt.Println(id, username, department, createTime)
	}

}

func delete() {
	stmt, err := db.Prepare("DELETE FROM user_info where id=?")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(3)
	if err != nil {
		panic(err)
	}

	id, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println(id)

}

func main() {
	// insert()
	// update()
	// query()
	delete()
	query()

}
