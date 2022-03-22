package function

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"jrtt/common"
	config2 "jrtt/config"
	"jrtt/logger"
	"jrtt/model"
	"jrtt/respdata"
	"log"
	"math/rand"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/**
获取列表
 */
func GetList(url string)  {
	var list []map[string]string
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),   //是否隐藏浏览器窗口，如果是win10要测试改为false
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36"),
		//使用chromium页面会出现滑块验证，用google-chrome不会，如果安装了chromium会默认使用chromium，需要安装google-chrome指定路径（用whereis google-chrome获取路径）；win10不会出现滑块验证可以删除这项配置
		chromedp.ExecPath("/usr/bin/google-chrome"),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, _ := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	err := chromedp.Run(ctx,chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(`.cs-view.margin-top-8.margin-bottom-8.pad-bottom-8.cs-view-block.cs-card`, chromedp.ByQuery),
		chromedp.AttributesAll(`div[data-log-extra]`, &list,chromedp.ByQueryAll),
	})
	if err != nil {
		logger.LogToMysql(err.Error(),true)
	}

	for _,v := range list{
		var luck_json respdata.List
		err = json.Unmarshal([]byte(v["data-log-extra"]),&luck_json)
		if err != nil {
			panic(err.Error())
		}
		CheckId(luck_json)
	}
	//body := common.Get(url)
	//logger.LogToMysql(body,false)
}

/**
检测是否可以抽奖
 */
func CheckId(luck respdata.List) {
	var config = config2.GetJrttUrl()
	var list_model model.List
	var user_model model.User
	response := common.Get(config["CHECK_URL"].(string)+luck.SearchResultId)
	var list respdata.LuckData
	err := json.Unmarshal([]byte(response),&list)
	if err != nil {
		panic(err.Error())
	}
	if list.ErrNo == 0 && list.Status == 1{

		//爬取详情
		detail:= GetDetail(luck.SearchResultId)



		timeobj := time.Unix(int64(list.LotteryTime), 0)
		date := timeobj.Format("2006-01-02 15:04:05")
		user_id := strconv.Itoa(list.Data.User.Info.UserId)
		user_url := config["USER_INFO"].(string)+user_id+"/"
		var token string
		if !user_model.IsTrue("uid = "+user_id+" and zhuanfa_uid = "+config["JR_UID"].(string)) {
			user_body := common.Get(user_url)
			compile1 := regexp.MustCompile(`<a class="avatar" href="\/c\/user\/token\/(.+?)\/\?source=author_home"`)
			submatch1 := compile1.FindStringSubmatch(user_body)
			token  = submatch1[1]
		}
		user_model.Edit(model.User{
			Uid: list.Data.User.Info.UserId,
			Name: list.Data.User.Info.Name,
			Url:user_url,
			ZhuanfaUid: config["JR_UID"].(string),
			Token: token,
		})
		list_model.Edit(model.List{
			SearchId: luck.SearchId,
			QueryId: luck.QueryId,
			SearchResultId: luck.SearchResultId,
			IsWinner: list.IsWinner,
			LotteryTime: date,
			ParticipateType: list.ParticipateType,
			Reward: list.Reward,
			Status: list.Status,
			WinnerType: list.WinnerType,
			LuckData: response,
			UserId: list.Data.User.Info.UserId,
			UserName: list.Data.User.Info.Name,
			Detail: detail,
			ZhuanfaUid: config["JR_UID"].(string),
		})
	}else{
		fmt.Println("不满足条件")
	}
}

func GetDetail(id string)string {
	var detail string
	var user_model model.User
	var config = config2.GetJrttUrl()

	body := common.Get(config["DETAIL"].(string)+id)

	dom,err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		panic(err.Error())
	}

	detail,err = dom.Find("div.weitoutiao-html").Html()
	if err != nil {
		panic(err.Error())
	}
	compile := regexp.MustCompile(`<a href="https:\/\/www\.toutiao\.com\/c\/user\/(\d+)\/\?tab=weitoutiao">@(.+?)<\/a>`)
	submatch := compile.FindAllStringSubmatch(detail, -1)

	for _,v := range submatch {
		var token string
		user_id,_ := strconv.Atoi(v[1])
		user_url := config["USER_INFO"].(string)+v[1]+"/"
		if !user_model.IsTrue("uid = "+v[1]+" and zhuanfa_uid = "+config["JR_UID"].(string)) {
			user_body := common.Get(user_url)
			compile1 := regexp.MustCompile(`<a class="avatar" href="\/c\/user\/token\/(.+?)\/\?source=author_home">`)
			submatch1 := compile1.FindStringSubmatch(user_body)
			token  = submatch1[1]
		}
		user_model.Edit(model.User{
			Uid: user_id,
			Name: v[2],
			Url:user_url,
			ZhuanfaUid: config["JR_UID"].(string),
			Token: token,
		})
	}
	return detail
}


/**
关注
 */
func FollowUser() {
	var config = config2.GetJrttUrl()
	var user_model  model.User
	user_list := user_model.List("is_modify = 0")
	for _,v := range user_list{
		var resp respdata.Api
		uid := strconv.Itoa(v.Uid)
		referer := config["REFERER_FOLLOW"].(string)+uid+"/"
		data := url.Values{}
		data.Add("user_id",uid)
		body := common.Post(config["FOLLOW"].(string),referer,data,"application/x-www-form-urlencoded")
		err := json.Unmarshal([]byte(body),&resp)
		if err != nil {
			panic("解析错误")
		}
		if resp.Message != "success" {
			panic("接口错误:"+body)
		} else {
			user_model.Edit(model.User{
				Id: v.Id,
				Uid: v.Uid,
				ZhuanfaUid: v.ZhuanfaUid,
				IsModify: 1,
			})
			fmt.Println("关注成功")
		}
		rand.Seed(time.Now().UnixNano())
		//加睡眠防止频繁请求被头条限制
		time.Sleep(time.Duration(rand.Intn(3)+2) * time.Minute)
	}
}

/**
转发
 */
func Repost() {
	var config = config2.GetJrttUrl()
	var list_model  model.List
	list := list_model.List("is_repost = 0")
	for _,v := range list {

		strArr := [5]string{"来了来了","好耶","真不错",""}
		rand.Seed(time.Now().UnixNano())
		str := strArr[rand.Intn(len(strArr)-1)]
		var strr string
		if strings.Index(v.Detail,"好友") != -1 || strings.Index(v.Detail,"1好友") != -1 || strings.Index(v.Detail,"1个好友") != -1 || strings.Index(v.Detail,"一好友") != -1 || strings.Index(v.Detail,"一个好友") != -1 || strings.Index(v.Detail,"1位好友") != -1 || strings.Index(v.Detail,"一位好友") != -1{
			strr = " @天猫 "
		}
		if strings.Index(v.Detail,"2好友") != -1 || strings.Index(v.Detail,"2个好友") != -1 || strings.Index(v.Detail,"二好友") != -1 || strings.Index(v.Detail,"二个好友") != -1 || strings.Index(v.Detail,"两好友") != -1 || strings.Index(v.Detail,"两位好友") != -1 || strings.Index(v.Detail,"二位好友") != -1 || strings.Index(v.Detail,"两个好友") != -1{
			strr = " @天猫 @京东 "
		}

		if strings.Index(v.Detail,"3好友") != -1 || strings.Index(v.Detail,"3个好友") != -1 || strings.Index(v.Detail,"三好友") != -1 || strings.Index(v.Detail,"三个好友") != -1 || strings.Index(v.Detail,"三好友") != -1 || strings.Index(v.Detail,"三位好友") != -1 || strings.Index(v.Detail,"三位好友") != -1 || strings.Index(v.Detail,"三个好友") != -1{
			strr = " @天猫 @京东 @腾讯 "
		}
		str += strr

		var resp respdata.Api
		referer := config["DETAIL"].(string)+v.SearchResultId
		data := url.Values{}
		data.Add("content",str)
		data.Add("fw_id_type","2")
		data.Add("opt_id_type","2")
		data.Add("repost_type","212")
		data.Add("repost_to_comment","0")
		data.Add("fw_user_id",strconv.Itoa(v.UserId))
		data.Add("fw_id",v.SearchResultId)
		data.Add("opt_id",v.SearchResultId)
		data.Add("group_id",v.SearchResultId)
		data.Add("item_id",v.SearchResultId)
		body := common.Post(config["REPOST"].(string),referer,data,"application/x-www-form-urlencoded")
		err := json.Unmarshal([]byte(body),&resp)
		if err != nil {
			panic("解析错误")
		}
		if resp.Message != "success" {
			panic("接口错误:"+body)
		}else{
			list_model.Edit(model.List{
				Id: v.Id,
				SearchResultId: v.SearchResultId,
				ZhuanfaUid: v.ZhuanfaUid,
				IsRepost: 1,
			})
			fmt.Println("转发成功")
		}
		time.Sleep(2*time.Minute)
	}
}

/**
转发
*/
func Like() {
	var config = config2.GetJrttUrl()
	var list_model  model.List
	list := list_model.List("is_like = 0")
	for _,v := range list {

		var resp respdata.Api
		referer := config["DETAIL"].(string)+v.SearchResultId
		data := url.Values{}
		body := common.Post(config["LIKE"].(string)+v.SearchResultId,referer,data,"application/x-www-form-urlencoded")
		err := json.Unmarshal([]byte(body),&resp)
		if err != nil {
			panic("解析错误")
		}
		if resp.Message != "success" {
			logger.LogToMysql("接口错误:"+body,false)
		}else{
			list_model.Edit(model.List{
				Id: v.Id,
				SearchResultId: v.SearchResultId,
				ZhuanfaUid: v.ZhuanfaUid,
				IsLike: 1,
			})
			fmt.Println("点赞成功")
		}
		time.Sleep(2*time.Minute)
	}
}
