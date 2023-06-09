package service

import (
	"dbOper/model"
	"embed"
	"fmt"
	"strings"
)

// 创建表名
func CreateTable(f embed.FS) {
	if isCreate() {
		return
	}

	data, _ := f.ReadFile("resources/dbops.sql")
	sqlArr := strings.Split(string(data), ";")
	for _, sql := range sqlArr {
		if sql == "" {
			continue
		}
		result := GetMysqlCon().Exec(sql)
		if result.Error != nil {
			fmt.Println(result.Error)
		}
	}
}

//查询是否初始化表结构,true代表已经创建，false未创建
func isCreate() bool {
	result := GetMysqlCon().Find(&model.Version{})
	if result.Error != nil {
		return false
	}
	if result.RowsAffected != 0 {
		return true
	}
	return false
}
