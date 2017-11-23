// dataBaseFunc
package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//http://www.01happy.com/golang-mysql-demo/
func dbInsertUser(db *sql.DB, user User) (err error) {

	stmt, err := db.Prepare("INSERT INTO user(name, password, status) VALUES(?, ?, ?)")
	defer stmt.Close()

	if err != nil {
		log.Println(err)
		return err
	}
	//插入动作
	stmt.Exec(user.Name, user.Password, user.Status)

	return nil
}

func dbDeleteUser(db *sql.DB, username string) (err error) {
	stmt, err := db.Prepare("DELETE FROM user WHERE name=?")
	defer stmt.Close()

	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(username)
	checkErr(err)
	return nil
}

//通过名字判断改哪里 把所有数据传过来
func dbUpdateUser(db *sql.DB, user User) (err error) {
	stmt, err := db.Prepare("UPDATE user SET password=?,status=? WHERE name=?")
	defer stmt.Close()

	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(user.Password, user.Status, user.Name)
	checkErr(err)
	return nil
}

//暂时实现为全表查询
func dbQueryUser(db *sql.DB) (users []User, err error) {

	rows, err := db.Query("select * from user")
	//rows, err := db.Query("select * from user where id = ?", 1)
	if err != nil {
		log.Println(err)
	}
	// rows means result set
	defer rows.Close()
	var user User

	//遍历 遍历的时候添加
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Password, &user.Status)

		if err != nil {
			log.Fatal(err)
		}
		//log.Println(user.id, user.name, user.password, user.status)
		//添加到切片中
		users = append(users, user)
	}
	//ShowUsers(users)
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		//return err
	}
	return users, nil
}

func dbQueryByname(db *sql.DB, username string) (users []User, err error) {
	rows, err := db.Query("select * from user where name='" + username + "'")
	if err != nil {
		log.Println(err)
	}
	// rows means result set
	defer rows.Close()
	var user User

	//遍历 遍历的时候添加
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Password, &user.Status)

		if err != nil {
			log.Fatal(err)
		}
		log.Println(user.Id, user.Name, user.Password, user.Status)
		//添加到切片中
		users = append(users, user)
	}
	//ShowUsers(users)
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return users, nil
}
