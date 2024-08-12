package MissEvan

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func (m MissEvan) ExtractStream(urls []string) error {
	//TODO implement me
	panic("implement me")
}

// 拉音频流,传入保存的文件夹路径和目标url
func (m MissEvan) ExtractAudio(path, targetUrl string) error {

	currentTime := time.Now()
	outputPattern := currentTime.Format("200601021504") + "_%02d.flv"
	//以1分钟为界限，暂时测试用
	segmentTime := 60
	//segmentTime:=60*60

	err := sliceByTime(targetUrl, outputPattern, path, segmentTime)
	if err != nil {
		log.Fatalln("Error")
	}
	return nil
}

func sliceByTime(streamURL, outputPattern, storagePath string, segmentTime int) error {

	fullOutputPattern := fmt.Sprintf("%s/%s", storagePath, outputPattern)
	cmd := exec.Command("ffmpeg", "-i", streamURL, "-c", "copy", "-f", "segment", "-segment_time",
		fmt.Sprintf("%d", segmentTime), "-reset_timestamps", "1", fullOutputPattern)

	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("FFmpeg 进程启动失败: %w", err)
	}

	log.Println("在新的命令行窗口中启动了 FFmpeg 进程")
	return nil
}
