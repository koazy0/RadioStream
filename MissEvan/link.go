package MissEvan

import (
	"RadioStream/lib"
	"context"
	"encoding/json"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func (m MissEvan) GetUrl2(url string) (urls []string) {

	// 创建一个新的无头浏览器上下文
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Headless,              // 设置为无头模式
		chromedp.DisableGPU,            // 禁用GPU加速
		chromedp.NoDefaultBrowserCheck, // 禁用默认的dev-shm-usage，防止部分系统上的共享内存不足
	)

	// 创建一个带有选项的上下文和取消函数
	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer allocCancel()

	// 创建一个新的浏览器上下文
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// 设置自定义请求头,防止被过滤
	headers := map[string]interface{}{
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.5735.110 Safari/537.36",
		"Accept-Language": "en-US,en;q=0.9",
		"Accept-Encoding": "gzip, deflate, br",
		"Connection":      "keep-alive",
		"Referer":         "https://missevan.com/",
	}
	// 要查找的特定URL，指向目标URL
	targetURL := "https://d1-"

	// 启用网络请求监听,找寻需要的URL
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if ev, ok := ev.(*network.EventRequestWillBeSent); ok {
			url := ev.Request.URL
			// 找寻前缀
			if strings.HasPrefix(url, targetURL) {
				fmt.Println("Found target URL:", url)
				urls = append(urls, url)
			}
		}
	})
	// 创建一个任务来启动网络监听和导航
	var res string
	err := chromedp.Run(ctx,
		// 启用网络监听
		network.Enable(),
		// 设置自定义请求头
		network.SetExtraHTTPHeaders(network.Headers(headers)),
		// 导航到目标页面
		chromedp.Navigate(url),
		// 等待页面的一部分加载完成（比如等待body元素出现）
		chromedp.WaitVisible(`body`),
		// 获取页面标题（作为示例操作）
		chromedp.Title(&res),
	)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second) //sleep 1s，使页面能加载完
	return
}

func (m MissEvan) GetUrl(url string) (args []string, err error) {

	prefix := "https://fm.missevan.com/live/"

	url = "https://fm.missevan.com/api/v2/live/" + strings.TrimPrefix(url, prefix)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}

	// 解析 JSON 字符串
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		log.Fatalf("JSON 解析失败: %v", err)
	}

	performerName, err := lib.JsonGetString(result, "info.room.creator_username")
	targetUrl, err := lib.JsonGetString(result, "info.room.channel.flv_pull_url")
	broadcastName, err := lib.JsonGetString(result, "info.room.name")

	args = append(args, targetUrl, broadcastName, performerName)
	return
}
