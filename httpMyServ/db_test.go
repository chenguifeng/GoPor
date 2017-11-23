// db_test.go
package main

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	//直接查看结果是否为空？
	fmt.Println("start test")
	users, err := dbQueryUser(Db)

	if err != nil {
		t.Error("Query func failed.")
	}
	//看出是否有数据 判断切片的大小
	if len(users) == 0 {
		t.Error("database got no data")
	} else {
		ShowUsers(users)
	}

}

func TestInsert(t *testing.T) {
	//dbInsertUser(db *sql.DB, user User) (err error)
	//构造插入的User
	testuser := User{Name: "chenguifeng", Password: "chenguifeng"}
	err := dbInsertUser(Db, testuser)
	if err != nil {
		panic(err)
	}
	//查询看是否成功
	rows, err := Db.Query("select * from user where name='chenguifeng'")
	if err != nil {
		t.Error(err)
	}
	// rows means result set
	defer rows.Close()
	var user User

	//遍历 遍历的时候添加
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Password, &user.Status)

		if err != nil {
			t.Error(err)
		}
		fmt.Println(user.Id, user.Name, user.Password, user.Status)
		//添加到切片中
		//users = append(users, user)
	}
	if user.Name == "chenguifeng" {
		fmt.Println("InsertFunc ok!")
	} else {
		t.Error("InsertFunc failed.")
	}

}

func TestUpdate(t *testing.T) {
	fmt.Println("test Update")
	user := User{Name: "chenguifeng", Password: "aaaaaa"}
	err := dbUpdateUser(Db, user)
	if err != nil {
		t.Error("update func fail")
	}
	resUsers, err := dbQueryByname(Db, user.Name)
	if err != nil {
		t.Error("update func fail")
	}
	if len(resUsers) == 0 {
		t.Error("update func fail")
	}
	//遍历对比是否全部修改完成
	for _, v := range resUsers {
		if v.Password != user.Password {
			t.Error("update func fail")
		}
	}

}

func TestDelete(t *testing.T) {
	fmt.Println("TestDelete")
	username := "chenguifeng"
	err := dbDeleteUser(Db, username)
	if err != nil {
		t.Error("Delete failed.")
	}
	users, err := dbQueryByname(Db, username)
	if err != nil {
		t.Error("dbDelete failed.")
	}

	if len(users) != 0 {
		t.Error("still got data, Delete fail")
	}

}
