package models

import (
	"fmt"
	"gofiber02-gorm/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var (
	//数据库连接实例
	DBConn *gorm.DB
)

//gorm + mysql 初始化数据库连接实例
func Setup() {
	//mysql连接的配置
	mysqlConfig := mysql.Config{
		DSN:                       getConnetInfo(), // DSN data source name
		DefaultStringSize:         255,             // string 类型字段的默认长度
		DisableDatetimePrecision:  true,            // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,            // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,            // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,           // 根据当前 MySQL 版本自动配置
	}

	//gorm的配置
	gormConfig := &gorm.Config{
		//使用单数表名
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info), //日志
		//CreateBatchSize: 1000, //指定批量操作每次 操作的行数
	}

	var err error
	//连接mysql
	DBConn, err = gorm.Open(mysql.New(mysqlConfig), gormConfig)
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	} else {
		//根据mysql的连接，维护连接池
		sqlDB, _ := DBConn.DB()
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(10)
		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(100)
		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Hour * 1)
		fmt.Println("connect to database successs!")
	}

}

//生成连接字符串
func getConnetInfo() string {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DatabaseSetting.User, conf.DatabaseSetting.Password, conf.DatabaseSetting.Host, conf.DatabaseSetting.Name)

	return dsn
}
