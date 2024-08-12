package MissEvan

import "fmt"

func (m MissEvan) Parse(url string) {
	args, _ := m.GetUrl(url)
	path, ok := m.InitDir(args)
	if !ok {
		fmt.Printf("[*]Error: Init Directory Error!")
	}

	//提取音频流
	m.ExtractAudio(path, args[0])
}
