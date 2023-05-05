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

// 发送Post请求
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

// 发送Get请求
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

// 使用代理发送请求
func PostWithProxy(postUrl string, json string, headers map[string]string, cookies map[string]string, proxy string, timeout int) (string, string, float64) {
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

	for k, v := range cookies {
		req.AddCookie(&http.Cookie{Name: k, Value: v})
	}

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
	for _, c := range resp.Cookies() {
		cookieStr += c.Name + "=" + c.Value + ";"
	}
	//WriteDebug(fmt.Sprintf("Post Result, Cost %.2fs\n%s\n%s\n%s", timeCost, requestMsg, body, cookieStr))
	return string(body), cookieStr, timeCost
}

// 发送Post请求
//
//	requestUrl: 请求地址
//	postData: 请求的内容
//	headers: 请求头
//	cookies: 请求cookies
//	proxy: 代理
//	timeout: 超时时间（秒）
//	返回值：返回的内容，返回的cookies，请求用时，错误信息
func PostWeb(requestUrl, postData string, headers, cookies map[string]string, proxy string, timeout int) (string, []*http.Cookie, float64, error) {
	return sendRequest(requestUrl, http.MethodPost, postData, headers, cookies, proxy, timeout)
}

// 发送Get请求
//
//	requestUrl: 请求地址
//	headers: 请求头
//	cookies: 请求cookies
//	proxy: 代理
//	timeout: 超时时间（秒）
//	返回值：返回的内容，返回的cookies，请求用时，错误信息
func GetWeb(requestUrl string, headers, cookies map[string]string, proxy string, timeout int) (string, []*http.Cookie, float64, error) {
	return sendRequest(requestUrl, http.MethodGet, "", headers, cookies, proxy, timeout)
}

// 使用代理发送请求
func sendRequest(requestUrl, method, postData string, headers, cookies map[string]string, proxy string, timeout int) (string, []*http.Cookie, float64, error) {
	defer errutil.CatchError()

	var client = &http.Client{}

	//设置超时时间
	if timeout > 0 {
		client.Timeout = time.Second * time.Duration(timeout)
	}

	//是否使用代理
	if proxy != "" {
		var p = func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxy)
		}
		var transport = &http.Transport{DisableKeepAlives: true, Proxy: p}
		client.Transport = transport
	}

	//转换成postBody
	bytesData := bytes.NewBuffer([]byte(postData))
	req, err := http.NewRequest(method, requestUrl, bytesData)
	if err != nil {
		logutil.Error(err.Error())
		return "", nil, 0, err
	}
	//用完后立即关闭连接
	req.Close = true

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	for k, v := range cookies {
		req.AddCookie(&http.Cookie{Name: k, Value: v})
	}

	tStart := time.Now()
	resp, err := client.Do(req)

	//请求用时（秒）
	timeCost := time.Since(tStart).Seconds()
	if err != nil {
		return "", nil, timeCost, err
	}

	defer resp.Body.Close()

	//返回内容
	content, _ := ioutil.ReadAll(resp.Body)

	return string(content), resp.Cookies(), timeCost, nil
}

// 获取请求head的基本参数（User-Agent/Accept/Accept-Language）
func GetBaseHeaders() map[string]string {
	var headers = make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.58"
	headers["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"
	headers["Accept-Language"] = "zh-CN,zh;q=0.9"
	return headers
}
