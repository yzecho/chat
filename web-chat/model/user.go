package model

import (
	"web-chat/database"
)

type UserInfo struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var DB = database.DB

func CheckPassword(user UserInfo) error {
	sqlStr := "select id from user where username = ? and password = ?"
	err := DB.Get(&user, sqlStr, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func CheckUserIfExist(user UserInfo) (bool, error) {
	sqlStr := "select * from user where username = ?"
	err := DB.Get(&user, sqlStr, user.Username)
	if err == nil {
		return false, nil
	}
	return true, nil
}

func AddUser(user UserInfo) (int, error) {
	sqlSql := "insert into user(username, password)values(?,?)"
	res, err := DB.Exec(sqlSql, user.Username, user.Password)
	if err != nil {
		return 0, err
	}
	theID, err := res.LastInsertId() // 新插入数据的id
	if err != nil {
		return 0, err
	}
	return int(theID), nil
}
