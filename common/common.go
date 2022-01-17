package common

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/chromedp/chromedp"
	"io"
	"io/ioutil"
	"jrtt/config"
	"jrtt/logger"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)



/**
获取今日头条sign
 */
func GetSign(url string)string  {
	var api = "http://127.0.0.1:8080/sign?url="+url
	var sign string
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("enable-automation", false),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, _ := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	err := chromedp.Run(ctx,chromedp.Tasks{
		chromedp.Navigate(api),
		chromedp.Value(`#sign`, &sign, chromedp.ByID),
	})
	if err != nil {
		logger.LogToMysql(err.Error(),true)
	}
	return sign
}


func JsonEncode(data interface{})string {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.Encode(data)
	return string(buffer.Bytes())
}

func JsonDecode(json []byte)*simplejson.Json {
	js,_ := simplejson.NewJson(json)
	return js
}


func TimeStamp(timeString string)int64 {
	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local")    //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, timeString, loc)
	timestamp := tmp.Unix()    //转化为时间戳 类型是int64
	return timestamp
}

func JsonAtoi(str interface{}) int {
	str = str.(json.Number).String()
	str,_ = strconv.Atoi(str.(string))
	return str.(int)
}



//发送GET请求
//url:请求地址
//response:请求返回的内容
func Get(url string,cookie ...http.Cookie) (response string) {
	config := config.GetJrttUrl()
	client := &http.Client{Timeout: 5 * time.Second}
	var req *http.Request

	req, _ = http.NewRequest("GET", url, nil)
	cookie1 := &http.Cookie{Name: "ttwid",Value: config["TTWID"].(string)}
	cookie2 := &http.Cookie{Name: "sessionid",Value: config["SESSIONID"].(string)}
	req.AddCookie(cookie1)
	req.AddCookie(cookie2)

	for _,v := range cookie {
		req.AddCookie(&v)
	}

	req.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36")
	resp, error := client.Do(req)

	defer resp.Body.Close()
	if error != nil {
		panic(error)
	}
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	response = result.String()

	fileName := "./test.html"
	dstFile,err := os.Create(fileName)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	s:=response
	dstFile.WriteString(s + "\n")

	return
}

//发送POST请求
//url:请求地址，data:POST请求提交的数据,contentType:请求体格式，如：application/json
//content:请求放回的内容
func Post(url string,referer string, data url.Values, contentType string,cookie ...http.Cookie) (content string) {
	config := config.GetJrttUrl()
	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	req.Header.Set("content-type", contentType)
	req.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36")
	req.Header.Set("referer", referer)
	req.Header.Set("x-csrftoken", config["CSRFTOKEN"].(string))
	cookie1 := &http.Cookie{Name: "ttwid",Value: config["TTWID"].(string)}
	cookie2 := &http.Cookie{Name: "sessionid",Value: config["SESSIONID"].(string)}
	cookie3 := &http.Cookie{Name: "csrftoken",Value: config["CSRFTOKEN"].(string)}
	//cookie4 := &http.Cookie{Name: "passport_csrf_token",Value: config["PASSPORT_CSRF_TOKEN"].(string)}
	cookie4 := &http.Cookie{Name: "passport_auth_status",Value: config["PASSPORT_AUTH_STATUS"].(string)}
	req.AddCookie(cookie1)
	req.AddCookie(cookie2)
	req.AddCookie(cookie3)
	req.AddCookie(cookie4)
	for _,v := range cookie {
		req.AddCookie(&v)
	}
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)
	return
}
