package Stream

type Stream interface {
	ExtractStream(args []string) error         //提取流
	ExtractAudio(path, targetUrl string) error //提取音频
}
