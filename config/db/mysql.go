package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn string
var MySql *gorm.DB

func MysqlDbSourceConfigString() string {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	username := "root"       //账号
	password := "root123"    //密码
	host := "120.26.242.221" //数据库地址，可以是Ip或者域名
	port := 3306             //数据库端口
	Dbname := "sample-store" //数据库名
	timeout := "10s"         //连接超时，10秒

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	return dsn
}

func init() {
	configString := MysqlDbSourceConfigString()
	db, err := gorm.Open(mysql.Open(configString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	MySql = db
}
