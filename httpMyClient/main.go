// httpMyClient project main.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var url = "http://localhost:8086/User"

//需要发送结构体
type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Status   int    `json:"Status"`
}

func main() {

	fmt.Println("Hello World!")
	user1 := User{Name: "lcf", Password: "12345"}
	fmt.Println(user1)
	var err error

	//GET
	err = QueryUser()
	checkErr(err)

	//POST
	err = AddUser(&user1)
	if err != nil {
		//_, file, line, ok := runtime.Caller(1)
		panic("fail in line")
	}

	//UPDATE
	user2 := user1
	user2.Password = "11111"
	err = UpdateUser(&user2)
	checkErr(err)
	//ok!

	//DELETE
	err = DeleteUser("lcf")
	checkErr(err)
}

//Post
func AddUser(user *User) (err error) {
	//格式化为json格式
	//https://studygolang.com/articles/4830 如果不通过查看改网页
	fmt.Println("POST method")
	formatUser, err := json.Marshal(user)
	body := bytes.NewBuffer([]byte(formatUser))

	resp, err := http.Post(url, "application/json", body)
	//client := &http.Client{}
	if err != nil {
		panic(err)
	}
	//得到resp后的处理
	defer resp.Body.Close()

	retBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("ReadAll")
	}

	fmt.Println(string(retBody))

	return nil
}

//Delete 因为没有http.delete方法可以调用 所以自己构造
//http.NewRequest http://www.01happy.com/golang-http-client-get-and-post/
func DeleteUser(username string) (err error) {
	fmt.Println("DELETE method")
	formatUsername, err := json.Marshal(username)
	body := bytes.NewBuffer([]byte(formatUsername))

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, body) // application/x-www-form-urlencoded
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json") // 设置Header名字
	resp, err := client.Do(req)
	defer resp.Body.Close()

	retbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(retbody))

	return nil
}

//Put
func UpdateUser(user *User) (err error) {
	fmt.Println("UPDATE method")
	formatUser, err := json.Marshal(user)
	body := bytes.NewBuffer([]byte(formatUser))

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, body) // application/x-www-form-urlencoded
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json") // 设置Header名字
	resp, err := client.Do(req)
	defer resp.Body.Close()

	retbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(retbody))
	return nil
}

//Get
//查询返回结构体切片以供显示
func QueryUser() (err error) {
	fmt.Println("GET method")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() //last thing will close

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(body))
	//json格式解析
	var users []User
	err = json.Unmarshal(body, &users)
	if err != nil {
		panic(err)
	}
	//打印输出
	ShowUsers(users)
	return nil
}

func (user *User) String() string {
	return fmt.Sprintf("id:%v name:%v password:%v status:%v \n", user.Id, user.Name, user.Password, user.Status)

}
func ShowUsers(users []User) {
	for _, user := range users {
		fmt.Println(user)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
