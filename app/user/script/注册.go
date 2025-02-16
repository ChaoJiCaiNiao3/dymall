package main
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

//连接数据库
func initDB()(*sql.DB,error){
dsn:="gorm:gorm@tcp( 127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
db,err:=sql.Open("mysql",dsn)
if err!=nil{
	fmt.Println("数据库连接失败:",err)
	return nil,err
}
err=db.Ping()
if err!=nil {
	fmt.Println("数据库连接失败:",err)
	return nil,err
}
return db,nil
}

//获取前端请求
func registerPage(w http.RsponseWriter,r *http.Request) {
	tmp1:=template.Must(template.ParseFiles("这里是前端页面地址"));
	err:=tmp1.Execute(w,nil)
}
//处理表单提交信息
func registerSubmit(w http.ResponseWriter,r *http.Request) {
	if r.Method!=http.MethodPost {
		http.Redirect(w,r,"/register",http.StatusSeeOther)
		return
	}
	err:=r.ParseForm()
	if err!=nil {
		http.Error(w,"解析表单失败",http.StatusBadRequest)
		return
	}
	//获取表单数据
	username:=r.FormValue("username")
	password:=r.FormValue("password")
	
	//连接数据库
	db,err:=initDB()
	if err!=nil {
		http.Error(w,"数据库连接失败",http.StatusInternalServerError)
		return
	}
	defer db.Close()
	//插入数据
	insert:="insert into client(name,password) values(?,?)"
	result,err:=db.Exec(insert,username,password)
	if err!=nil {
		http.Error(w,"插入数据失败",http.StatusInternalServerError)
		return
	}
	// 获取插入的记录ID
    lastInsertID, err := result.LastInsertId()
    if err != nil {
        http.Error(w, "获取插入记录ID失败", http.StatusInternalServerError)
        return
    }

    // 获取受影响的行数
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        http.Error(w, "获取受影响行数失败", http.StatusInternalServerError)
        return
    }
	if rowsAffected==1 {
    fmt.Fprintf(w, "注册成功，您的身份ID为: %d\n", lastInsertID)
	}
	else {
		fmt.Fprintln(w,"注册失败")
		return
	}


}
func main() {
	//注册页面的访问路径
	http.HandleFunc("/registerPage",registerPage)
	//处理表单提交的路径
	http.HandleFunc("/registerSubmit",registerSubmit)
	log.Println("服务器启动成功")
	log.Fatal(http.ListenAndServe(":8080",nil))


}