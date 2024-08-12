package MissEvan

import (
	"fmt"
	"os"
	"time"
)

func (m MissEvan) TargetPath() {
	//TODO implement me
	panic("implement me")
}

// InitDir 传入url，房间名，主播名,返回应当存储的路径
func (m MissEvan) InitDir(args []string) (string, bool) {
	pathstr := "./Record/"
	performerName := "猫耳" + args[2] //平台+主播名

	currentTime := time.Now()
	formattedDate := currentTime.Format("20060102") // 格式化为 "YYYYMMDD" 字符串
	broadcastName := formattedDate + "_" + args[1]  //日期+房间名

	pathstr = pathstr + performerName
	err := os.Mkdir(pathstr, os.ModePerm)
	if err != nil {
		// 如果文件夹已存在，也忽略错误
		if !os.IsExist(err) {
			fmt.Println("[*]Error creating directory:", err)
			return "", false
		}
	}

	pathstr = pathstr + "/" + broadcastName
	err = os.Mkdir(pathstr, os.ModePerm)
	if err != nil {
		// 如果文件夹已存在，也忽略错误
		if !os.IsExist(err) {
			fmt.Println("[*]Error creating directory:", err)
			return "", false
		}
	}
	return pathstr, true
}
