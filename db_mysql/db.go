package db_mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"xiaoyuzi/models"
)
var Db *sql.DB
func init(){
	config :=beego.AppConfig
	dbdriver :=config.String("db_driverName")
	dbuser :=config.String("db_user")
	dbpassword:=config.String("db_password")
	dbip:=config.String("db_ip")
	dbname :=config.String("db_name")
	connurl:= dbuser+":"+dbpassword+"@tcp("+dbip+")/"+dbname+"?charset=utf8"
	db,err:=sql.Open(dbdriver,connurl)
	if err !=nil{
		panic("数据库错误")
	}
	Db=db
}
func InsertUser(user models.User)(int64,error) {
	//hash计算
	hashmd:=md5.New()
	hashmd.Write([]byte(user.Password))
	hashmd.Write([]byte(user.Name))
	hashmd.Write([]byte(user.Nick))
	hashmd.Write([]byte(user.Address))
	hashmd.Write([]byte(user.Birthday))
	bytes:=hashmd.Sum(nil)
	user.Password =hex.EncodeToString(bytes)
	user.Name=hex.EncodeToString(bytes)
	user.Nick=hex.EncodeToString(bytes)
	user.Birthday=hex.EncodeToString(bytes)
	user.Address=hex.EncodeToString(bytes)
	fmt.Println("用户名：",user.Nick,"密码：",user.Password,"小名：",user.Nick,"地址：",user.Address,"生日：",user.Birthday)
	res,err:=Db.Exec("insert into user (birthday,address,nick,password) value (?,?,?,?)",user.Birthday,user.Address,user.Nick,user.Password)
	if err !=nil {
		return -1,err
	}
	id,err:=res.RowsAffected()
	if err !=nil {
		return -1,err
	}
	return id,nil
}
