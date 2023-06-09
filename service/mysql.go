package service

import (
	"dbOper/util"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// 指定驱动
const DRIVER = "mysql"

var mysqlCon *gorm.DB

// 初始化连接数据库，生成可操作基本增删改查结构的变量
func InitMySql() {
	//连接数据库
	fmt.Println(getConnectUrl())
	var err error
	mysqlCon, err = gorm.Open(DRIVER, getConnectUrl())
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
	}
	mysqlCon.DB().SetMaxOpenConns(util.ApplicationConfig.GetInt("mysql.MaxOpenConns"))
	mysqlCon.DB().SetMaxIdleConns(util.ApplicationConfig.GetInt("mysql.MaxIdleConns"))
	mysqlCon.DB().SetConnMaxIdleTime(time.Duration(util.ApplicationConfig.GetInt("mysql.MaxIdleTime")) * time.Second)
}

// 数据库连接字符串
func getConnectUrl() string {
	//将yaml配置参数拼接成连接数据库的url
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		util.ApplicationConfig.GetString("mysql.User"),
		util.ApplicationConfig.GetString("mysql.Password"),
		util.ApplicationConfig.GetString("mysql.IP"),
		util.ApplicationConfig.GetString("mysql.Port"),
		util.ApplicationConfig.GetString("mysql.Database"),
	)
}

//获取MySQL数据库连接
func GetMysqlCon() *gorm.DB {
	return mysqlCon
}

// 关闭数据库连接
func Close() {
	mysqlCon.Close()
}
