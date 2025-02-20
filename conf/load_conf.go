package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
)

type databaseSetting struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
}

type adbSetting struct {
	Path     string
	PicTotal int
}

var (
	DatabaseSetting = &databaseSetting{}
	AdbSetting      = &adbSetting{}
	cfg             *ini.File
)

func LoadConfig() {
	confFile := os.Getenv("CONG_FILE")
	// 没传，用默认的
	if confFile == "" {
		confFile = "local.ini"
	}
	var err error
	cfg, err = ini.Load(filepath.Join("conf", confFile))
	if err != nil {
		panic(fmt.Sprintf("init cfg failed, when load cfg file, err:%s", err.Error()))
	}
	mapTo("db", DatabaseSetting)
	mapTo("adb", AdbSetting)
}

func mapTo(section string, target any) {
	err := cfg.Section(section).MapTo(target)
	if err != nil {
		panic(fmt.Sprintf("init cfg failed, when mapTo section:[%s], err:%s", section, err.Error()))
	}
}
