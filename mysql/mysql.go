package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync_azur_lane/conf"
	"time"
)

var MysqlDb *sql.DB

func Linksql() {
	conf.LoadConfig()
	//dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", conf.DatabaseSetting.User, conf.DatabaseSetting.Password, conf.DatabaseSetting.Host, conf.DatabaseSetting.Port, conf.DatabaseSetting.Database, "utf8")
	var MysqlDbErr error
	// 打开连接失败
	MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)
	//defer MysqlDb.Close();
	if MysqlDbErr != nil {
		log.Println("dbDSN: " + dbDSN)
		panic("数据源配置不正确: " + MysqlDbErr.Error())
	}

	// 最大连接数
	MysqlDb.SetMaxOpenConns(10)
	// 闲置连接数
	MysqlDb.SetMaxIdleConns(5)
	// 最大连接周期
	MysqlDb.SetConnMaxLifetime(3600 * time.Second)

	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
		panic("数据库链接失败: " + MysqlDbErr.Error())
	}
}
