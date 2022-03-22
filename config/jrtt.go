package config


func GetJrttUrl() map[string]interface{} {
	// 初始化数据库配置map
	urlConfig := make(map[string]interface{})

	urlConfig["LOGS_PATH"] = "/xxx/xxx/jrtt/logger/logs.log"	//日志路径  默认是记录到数据库,不需要更改

	//这些参数都可以在今日头条官网的cookie里面查找到
	urlConfig["JR_UID"] = "*****"
	urlConfig["SESSIONID"] = "*****"
	urlConfig["TTWID"] = "*****"
	//urlConfig["CSRFTOKEN"] = "*****"
	//urlConfig["PASSPORT_CSRF_TOKEN"] = "*****"
	//urlConfig["PASSPORT_AUTH_STATUS"] = "*****"
	urlConfig["CSRF_SESSION_ID"] = "*****"
	urlConfig["X-SECSDK-CSRF-TOKEN"] = "*****"
	urlConfig["page_size"] = 5   //这个是页码，一开始不建议爬太多，因为关注接口头条有限制


	//链接
	urlConfig["URL_LIST1"] = "https://so.toutiao.com/search?keyword=抽奖详情&pd=weitoutiao&source=pagination&dvpf=pc&page_num="  //keyword关键词 page_num页码
	urlConfig["URL_LIST2"] = "https://so.toutiao.com/search?keyword=转发抽奖&pd=weitoutiao&source=pagination&dvpf=pc&page_num="  //keyword关键词 page_num页码
	urlConfig["CHECK_URL"] = "https://is.snssdk.com/ugc/lottery/v2/result/?type=2&id="  // 检测是否有效抽奖
	urlConfig["USER_INFO"] = "https://www.toutiao.com/c/user/"  // 用户信息
	urlConfig["DETAIL"] = "https://www.toutiao.com/w/a"  //文章详情
	urlConfig["REPOST"] = "https://www.toutiao.com/c/ugc/content/repost/"  //转发
	urlConfig["LIKE"] = "https://www.toutiao.com/api/pc/user/fans_digg?op=1&gid="  //转发
	urlConfig["FOLLOW"] = "https://www.toutiao.com/c/user/follow/"  //关注
	urlConfig["REFERER_FOLLOW"] = "https://www.toutiao.com/c/user/"  //关注




	return urlConfig
}
