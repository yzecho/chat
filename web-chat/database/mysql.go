package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func init()  {
	dsn := "root:lz1015515@(localhost)/jwt-demo?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql init failed, err:%v\n", err)
	}
	DB.SetMaxOpenConns(20) // 设置与数据库连接的最大数目
	DB.SetMaxIdleConns(10) // 设置连接池中最大闲置连接数
}
