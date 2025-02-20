package main

import (
	"bytes"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gookit/goutil/fsutil"
	"log"
	"os"
	"os/exec"
	"sync_azur_lane/conf"
)

func main() {
	conf.LoadConfig()
	spew.Dump(conf.AdbSetting)
	err := connectAdb()
	if err != nil {
		fmt.Println("connectAdb: ", err)
		return
	}
	// 拉取仓库舰娘图片
	i := conf.AdbSetting.Total // 截图总数
	for ; i > 0; i-- {
		err = nextShip()
		if err != nil {
			fmt.Println("nextShip: ", err)
			return
		}
		err = screencap(i)
		if err != nil {
			fmt.Println("screencap: ", err)
			return
		}
	}
}

func connectAdb() (err error) {
	cmd := exec.Command("adb", "connect", conf.AdbSetting.Path)
	var errBuff bytes.Buffer
	cmd.Stderr = &errBuff
	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf(errBuff.String())
		return
	}
	return
}

// 屏幕分辨率： 720X1280
func nextShip() (err error) {
	cmd := exec.Command("adb", "shell", "input", "swipe", "400", "360", "300", "360")
	var errBuff bytes.Buffer
	cmd.Stderr = &errBuff
	//var out, errOut bytes.Buffer
	//cmd.Stdout = &out
	//cmd.Stderr = &errOut
	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf(errBuff.String())
		return
	}
	return
}

// 截图：
func screencap(i int) (err error) {
	originPath := "/sdcard/screen.png"
	// 截图
	cmd := exec.Command("adb", "screencap", originPath)
	var errBuff bytes.Buffer
	cmd.Stderr = &errBuff
	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf(errBuff.String())
		return
	}

	// 拉到本地
	picDir := GetPath("screens")
	if !fsutil.DirExist(picDir) {
		err = fsutil.Mkdir(picDir, 0777)
		if err != nil {
			return
		}
	}
	picPath := fsutil.JoinPaths(picDir, fmt.Sprintf("%d.png", i))
	cmd = exec.Command("adb", "pull", originPath, picPath)
	cmd.Stderr = &errBuff
	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf(errBuff.String())
		return
	}

	// 删除远程文件
	cmd = exec.Command("adb", "shell", "rm", originPath)
	cmd.Stderr = &errBuff
	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf(errBuff.String())
		return
	}
	return
}

func GetPath(subPath ...string) (pwd string) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Panic("GetPath err: ", err)
	}
	subPath = append([]string{pwd}, subPath...)
	return fsutil.JoinPaths(subPath...)
}
