package service

import (
	"dbOper/model"
	"fmt"
)

/**
用户登录
*/
func Login(username string, pwd string) (user *model.User, err error) {
	fmt.Println("login")
	if err = GetMysqlCon().Where("user_name=? and pwd=", username, pwd).First(user).Error; err != nil {
		return nil, err
	}
	return
}

/**
新建User信息
*/
func CreateUser(user *model.User) (err error) {
	if err = GetMysqlCon().Create(user).Error; err != nil {
		return err
	}
	return
}

/**
获取user集合
*/
func GetAllUser() (userList []*model.User, err error) {
	if err := GetMysqlCon().Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

/**
根据id删除user
*/
func DeleteUserById(id string) (err error) {
	err = GetMysqlCon().Where("id=?", id).Delete(&model.User{}).Error
	return
}

/**
根据id查询用户User
*/
func GetUserById(id string) (user *model.User, err error) {
	if err = GetMysqlCon().Where("id=?", id).First(user).Error; err != nil {
		return nil, err
	}
	return
}

/**
更新用户信息
*/
func UpdateUser(user *model.User) (err error) {
	err = GetMysqlCon().Save(user).Error
	return
}
