package main

import (
	"RadioStream/MissEvan"
	"fmt"
	"os"
)

func init() {
	folder := "./录播"
	err := os.Mkdir(folder, os.ModePerm)
	if err != nil {
		// 如果文件夹已存在，也忽略错误
		if !os.IsExist(err) {
			fmt.Println("[*]Error creating directory:", err)
			os.Exit(1)
		}
	}
}

func main() {

	url := "https://fm.missevan.com/live/247817392"
	p := MissEvan.MissEvan{}
	p.Parse(url)
}
