// httpMyServ project main.go
package main
//练习
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var Url = "http://localhost:8086/User"
var Db *sql.DB //作为全局变量 只需要连接一次

//需要发送结构体
//结构体转json需要 只有字段名是大写的，才会被编码到json当中。
//http://www.jianshu.com/p/f3c2105bd06b 例子
type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Status   int    `json:"Status"`
}

// 在外面声明了全局变量Db 在init函数中使用了 := 相当于重复 导致全局的Db依然为空
func init() {
	//fmt.Println("first run here")
	var err error
	Db, err = sql.Open("mysql", "root:root@/mydb")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}

}

//收到了改网页之后 判断相应的http方法来执行不同的动作 Restful风格 资源都是user
//动作不体现在url上
func main() {
	fmt.Println("Hello World!")
	//连接数据库
	/*
		Db, err := sql.Open("mysql", "root:sees7&chanting@/mydb")
		if err != nil {
			log.Fatalf("Open database error: %s\n", err)
		}
	*/
	defer Db.Close()

	addHttpHandle()

	err := http.ListenAndServe("localhost:8086", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func addHttpHandle() {
	http.HandleFunc("/User", Dispatch) //设置访问的路由

}

//判断到不同的请求类型分发到各自函数 C中可使用函数指针实现 暂用switch
func Dispatch(w http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	switch req.Method {
	case "GET": //查询
		fmt.Println("http GET method")
		hanQueryUser(w, req)
	case "POST": //增加
		fmt.Println("http POST method")
		handInserUser(w, req)
	case "DELETE": //删除
		fmt.Println("http DELETE method")
		handDeleteUser(w, req)
	case "PUT": //改变
		fmt.Println("http PUT method")
		handModifyUser(w, req)
	default:
		fmt.Println("非法值！！")
		http.Error(w, "The method is not allowed.", http.StatusMethodNotAllowed)
	}

}

/*
//数据库处理
	result,_ := ioutil.ReadAll(req.Body)
	req.Body.Close()
	fmt.Println("%s\n", result)
*/
func hanQueryUser(w http.ResponseWriter, req *http.Request) {
	//从sql中获取数据 传回客户端
	//users := make([]User, 0)
	users, err := dbQueryUser(Db)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("get from database")
	ShowUsers(users)
	//改为json格式
	resp, _ := json.Marshal(users)

	//传回给客户端
	fmt.Println("after string ", string(resp))
	fmt.Fprintf(w, string(resp))

}

func handInserUser(w http.ResponseWriter, req *http.Request) {
	//从客户端读取数据 json解码
	result, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()

	var user User
	json.Unmarshal(result, &user)
	fmt.Println(user)
	//插入数据库
	err := dbInsertUser(Db, user)
	if err != nil {
		panic(err)
	}
	//返回
	resp := "Insert ok!"
	fmt.Fprintf(w, string(resp))
}

func handDeleteUser(w http.ResponseWriter, req *http.Request) {
	//从客户端读取数据 json解码
	result, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()

	var username string
	json.Unmarshal(result, &username)
	fmt.Println(username)
	//从数据库中删除
	err := dbDeleteUser(Db, username)
	if err != nil {
		panic(err)
	}
	//返回
	resp := "Delete ok!"
	fmt.Fprintf(w, string(resp))
}

func handModifyUser(w http.ResponseWriter, req *http.Request) {
	//从客户端读取数据 json解码
	result, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()

	var user User
	json.Unmarshal(result, &user)
	fmt.Println(user)
	//从数据库中修改
	err := dbUpdateUser(Db, user)
	if err != nil {
		panic(err)
	}
	//返回
	resp := "Modify ok!"
	fmt.Fprintf(w, string(resp))

}

func (user User) String() string {

	return fmt.Sprintf("id:%-3v name:%-10v password:%-10v status:%-2v ", user.Id, user.Name, user.Password, user.Status)

}
func ShowUsers(users []User) {
	fmt.Println("ShowUsers")
	for _, user := range users {
		fmt.Println(user)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
