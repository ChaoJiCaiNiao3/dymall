package main

import (
	"context"
	"encoding/json"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	user "github.com/ChaoJiCaiNiao3/dymall/app/user/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{
	db *sql.DB
}

// Register implements the UserServiceImpl interface.
func initDB() (*sql.DB, error) {
	dsn:="gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=sql.Open("mysql",dsn)
	if err!=nil {
		return nil,err
	}
	err=db.Ping()
	if err!=nil {
		return nil,err
	}
	return db,nil
}
func handleRegister(w http.ResponseWriter,r *http.Request) {
	//判断post信息是否为注册信息(注册一般用post)
	if r.Method!=http.MethodPost {
		http.Error(w,"请求方式错误",http.StatusMethodNotAllowed)
		return
	}
	//处理前端请求
	var req user.RegisterReq
	err:=json.NewDecoder(r.Body).Decode(&req)
	if err!=nil {
		http.Error(w,"解析请求失败",http.StatusBadRequest)
		return
	}
	service:=&UserServiceImpl{}
	resp,err:=service.Register(context.Background(),&req)
	if err!=nil {
		http.Error(w,"注册失败",http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(resp)
}
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// TODO: Your code here...
	resp = &user.RegisterResp{}//结构体初始化,用于响应

	if req.Username==""||req.Password=="" {
		resp.Code=400
		resp.Message="用户名或密码不能为空"
		return resp,nil
	}
	if(s.db==nil) {
.db,err=initDB()
		if err!=nil {
			resp.Code=500
			resp.Message="数据库连接失败"
			return resp,nil
		}
	}

	//插入数据
	insert:="insert into user(name,password) values(?,?)"
	result,err:=s.db.ExecContext(ctx,insert,req.Username,req.Password)
	
	if err!=nil {
		resp.Code=500
		resp.Message="注册失败"
		return resp,nil
	}
	userID,err:=result.LastInsertId()
	if err!=nil {
		resp.Code=500
		resp.Message="获取用户ID失败"
		return resp,nil
	}
	defer.db.Close()
	resp.Code=200
	resp.Message=fmt.Sprintf("注册成功,用户ID为:%d",userID)
	return resp,nil
}

func handleLogin(w http.ResponseWriter,r *http.Request) {
	if r.Method!=http.MethodPost {//确保前端页面使用post发送json格式到正确接口即可
		http.Error(w,"请求方式错误",http.StatusMethodNotAllowed)
		return
	}	
	var req user.LoginReq
	err:=json.NewDecoder(r.Body).Decode(&req)
	if err!=nil {
		http.Error(w,"解析请求失败",http.StatusBadRequest)
		return
	}
	service:=&UserServiceImpl{}
	resp,err:=service.Login(context.Background(),&req)
	if err!=nil {
	http.Error(w,"登录出错",http.StatusInternalServerError)
	return	
	}
	if resp.Code==200 {
		http.Redirect(w,r,"/首页",http.StatusSeeOther)
		return
	}
	w.Header().Set("Content_Type","application/json")
	json.NewEncoder(w).Encode(resp)
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp=&user.LoginResp{}
	if req.Username==""||req.Password=="" {
		resp.Code=400
		resp.Message="用户名或密码不能为空"
		return resp,nil
	}
	if.db==nil {
.db,err=initDB()
		if err!=nil {
			resp.Code=500
			resp.Message="数据库连接失败"
			return resp,nil
		}
	}
	var storedPassword string
	query:="select password from user where name=?"
	err=s.db.QueryRowContext(ctx,query,req.Username).Scan(&storedPassword)
	if err!=nil {
		if err==sql.ErrNoRows {
			resp.Code=404
			resp.Message="用户不存在"
		}
		else {
			resp.Code=500
			resp.Message="登录失败"
		}
		return resp,nil
	}
	if req.Password!=storedPassword {
		resp.Code=401
		resp.Message="密码错误"
		return resp,nil	
	}
	resp.Code=200
	resp.Message="登录成功"

	return resp,nil
}

func (s *UserServiceImpl) GetAddress(ctx context.Context, req *user.GetAddressReq) (resp *user.GetAddressResp, err error) {
	// TODO: Your code here...
	resp=&user.GetAddressResp{}
	if req.UserID==0 {
		resp.Code=400
		resp.Message="用户ID不能为空"
		return resp,nil
	}
	if.db==nil {
.db,err=initDB()
		if err!=nil {
			resp.Code=500
			resp.Message="数据库连接失败"
			return resp,nil
		}	
	}
	var address string
	query:="select address from address where id=?"
	err=s.db.QueryRowContext(ctx,query,req.UserID).Scan(&address)
	if err!=nil {
		if err==sql.ErrNoRows {
			resp.Code=404
			resp.Message="用户不存在"
		}
		else {
			resp.Code=500
			resp.Message="获取地址失败"
		}
		return resp,nil
	}
	resp.Code=200
	resp.Message="获取地址成功"
	resp.Address=address
	return resp,nil
}
