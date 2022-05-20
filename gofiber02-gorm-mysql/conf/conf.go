package conf

import (
	"github.com/go-ini/ini"
	"log"
)

/*
读取配置文件，并将配置信息存入全局变量

*/

//数据库配置
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var (
	DatabaseSetting = &Database{} //初始化一个数据库对象
	cfg             *ini.File
)

// 初始化配置实例
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("database", DatabaseSetting)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
