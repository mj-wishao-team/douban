package dao

import (
	"database/sql"
	"douban/tool"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	cfg := tool.GetCfg().Database
	db, err := sql.Open(cfg.Driver, cfg.User+":"+cfg.Password+"@tcp("+cfg.Host+":"+cfg.Port+")/"+cfg.DbName+"?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
	fmt.Println("数据库链接成功")
}
