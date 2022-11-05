package webutil

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/zhouweigithub/goutil/errutil"
	"github.com/zhouweigithub/goutil/logutil"
)

// Post 发送Post请求到指定URL
func Post(url string, postData string, contentType string) (string, error) {
	para := bytes.NewBuffer([]byte(postData))
	startTime := time.Now()

	resp, err := http.Post(url, contentType, para)

	if err != nil {
		timeDiff := time.Since(startTime).Seconds()
		msg := fmt.Sprintf("POST URL ERROR:COST %.2fs\nURL:%s\nPOST DATA:%s\nERROR:%s", timeDiff, url, postData, err.Error())
		logutil.Error(msg)
		return "", err
	}
	//用完后立即关闭连接
	resp.Close = true
	defer resp.Body.Close()

	timeDiff := time.Since(startTime).Seconds()
	resString, _ := ioutil.ReadAll(resp.Body)

	msg := fmt.Sprintf("POST URL:COST %.2fs\nurl:%s\nPOST DATA:%s\nRESULT:%s", timeDiff, url, postData, resString)
	logutil.Debug(msg)

	return string(resString), nil
}

// Get 发送Get请求到指定URL
func Get(url string) (string, error) {
	defer errutil.CatchError()
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	//用完后立即关闭连接
	resp.Close = true
	defer resp.Body.Close()
	str, _ := ioutil.ReadAll(resp.Body)
	return string(str), nil
}

// 发送Get请求到指定URL
func GetWithTimeOut(url string, timeout time.Duration) (string, float64, error) {
	defer errutil.CatchError()

	client := &http.Client{
		Timeout: timeout,
	}
	//开始时间
	tStart := time.Now()
	resp, err := client.Get(url)
	//请求用时（秒）
	timeCost := time.Since(tStart).Seconds()
	if err != nil {
		logutil.Error(fmt.Sprintf("client.Get 请求出错。用时 %.2f 秒\n%s\n%v", timeCost, url, err.Error()))
		return "", timeCost, err
	}

	defer resp.Body.Close()

	//返回内容
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), timeCost, err
}

// PostWithProxy 使用代理发送请求
func PostWithProxy(postUrl string, json string, headers map[string]string, proxy string, timeout int) (string, string, float64) {
	defer errutil.CatchError()

	var client = &http.Client{}
	//转换成postBody
	bytesData := bytes.NewBuffer([]byte(json))
	//请求的参数
	requestMsg := fmt.Sprintf("postJson:%s", json)
	//设置超时时间
	client.Timeout = time.Second * time.Duration(timeout)

	//是否使用代理
	if proxy != "" {
		var p = func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxy)
		}
		var transport = &http.Transport{DisableKeepAlives: true, Proxy: p}
		client.Transport = transport
	}

	//post请求
	req, err := http.NewRequest("POST", postUrl, bytesData)
	if err != nil {
		logutil.Error(err.Error())
		return "", "", 0
	}
	//用完后立即关闭连接
	req.Close = true

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	//开始时间
	tStart := time.Now()

	resp, err := client.Do(req)
	//请求用时（秒）
	timeCost := time.Since(tStart).Seconds()
	if err != nil {
		logutil.Error(fmt.Sprintf("PostWithProxy.Do 请求出错。用时 %.2f 秒\n%s\n%v", timeCost, requestMsg, err.Error()))
		return "", "", timeCost
	}

	defer resp.Body.Close()

	//返回内容
	body, _ := ioutil.ReadAll(resp.Body)

	//解析返回的cookie
	var cookieStr string
	cookies := resp.Cookies()
	for _, c := range cookies {
		cookieStr += c.Name + "=" + c.Value + ";"
	}
	//WriteDebug(fmt.Sprintf("Post Result, Cost %.2fs\n%s\n%s\n%s", timeCost, requestMsg, body, cookieStr))
	return string(body), cookieStr, timeCost
}
