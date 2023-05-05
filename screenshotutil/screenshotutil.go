package screenshotutil

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"math"
	"path"
	"strings"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/zhouweigithub/goutil/fileutil"
)

// 网页截图功能
type screenShot struct {
	chromePath string // chrome程序路径
	imagePath  string // 网页截图的存放目录
}

// 创建新实例
//
//	chromePath: chrome程序路径
//	imageSaveFolder: 截图保存目录
func New(chromePath, imageSaveFolder string) *screenShot {
	var ss = screenShot{
		chromePath: chromePath,
		imagePath:  imageSaveFolder,
	}
	return &ss
}

// 进行网页截屏
func (s *screenShot) screenshot(ctx context.Context, fileName string) error {
	_, _, contentSize, _, _, _, err := page.GetLayoutMetrics().Do(ctx)
	if err != nil {
		return err
	}

	width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

	// force viewport emulation
	err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
		WithScreenOrientation(&emulation.ScreenOrientation{
			Type:  emulation.OrientationTypePortraitPrimary,
			Angle: 0,
		}).
		Do(ctx)
	if err != nil {
		return err
	}

	// capture screenshot
	var buf []byte
	buf, err = page.CaptureScreenshot().
		WithQuality(90).
		WithClip(&page.Viewport{
			X:      contentSize.X,
			Y:      contentSize.Y,
			Width:  contentSize.Width,
			Height: contentSize.Height,
			Scale:  1,
		}).Do(ctx)
	if err != nil {
		return err
	}
	var filePath = path.Dir(fileName)
	fileutil.CreateFolderIfNotExists(filePath)
	if err := ioutil.WriteFile(fileName, buf, 0644); err != nil {
		return err
	}
	return nil
}

// 进行网页截屏
func (s *screenShot) ScreenShot(pageUrl string) (fileName string, err error) {
	if len(s.chromePath) == 0 {
		return "", errors.New("chrome程序路径未配置")
	}
	if len(pageUrl) == 0 {
		return "", errors.New("pageUrl不能为空")
	}
	//设置chrome安装路径，如果是windows请自行设置允许路径，比如：C:\Program Files (x86)\Google\Chrome\Application\chrome.exe
	chromedp.ExecPath(s.chromePath)
	//增加选项，不允许chrome窗口显示出来
	options := []chromedp.ExecAllocatorOption{
		//不允许chrome窗口显示出来
		chromedp.Flag("headless", true),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
		//设置浏览器UserAgent
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	//创建chrome窗口
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var title string
	fileName = s.getImgFileName(pageUrl)
	//运行爬虫代码
	err = chromedp.Run(ctx,
		//请求网址
		chromedp.Navigate(pageUrl),
		//等待网页加载完毕
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		//获取网页标题
		chromedp.Evaluate("document.title", &title),
		//开始进行截屏
		chromedp.ActionFunc(func(ctx context.Context) error {
			return s.screenshot(ctx, fileName)
		}),
	)
	if err != nil {
		fileName = ""
	}
	return fileName, err
}

// 获取截图保存的文件名
func (s *screenShot) getImgFileName(pageUrl string) string {
	pageUrl = strings.TrimRight(pageUrl, "/")
	var todayString = time.Now().Format("2006-01-02")
	var reg = md5.Sum([]byte(pageUrl))
	return s.imagePath + "/" + todayString + "/" + hex.EncodeToString(reg[:]) + ".png"
}
