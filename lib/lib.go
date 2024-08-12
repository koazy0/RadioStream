package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadUrlFromTxt(path string) (urls []string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	return
}

func JsonGetString(json map[string]interface{}, target string) (str string, err error) {
	strs := strings.Split(target, ".")
	err_str := "[*]Parse Json Error: "

	//存储中间值
	tmp_map := json
	for i := 0; i < len(strs)-1; i++ {
		m, ok := tmp_map[strs[i]].(map[string]interface{})
		err_str = err_str + strs[i] + "."

		// 类型断言失败，即解析出错，直接返回空字符串和err
		if !ok {
			err_str = err_str + "\n"
			err = fmt.Errorf(err_str)
			return "", err
		}

		//继续解析
		tmp_map = m
	}

	s, ok := tmp_map[strs[len(strs)-1]].(string)
	//最后一步解析失败
	if !ok {
		err_str = err_str + strs[len(strs)-1] + "\n"
		err = fmt.Errorf(err_str)
		return "", err
	}

	str = s
	return str, nil
}
