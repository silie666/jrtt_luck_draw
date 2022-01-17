package main

import (
	config2 "jrtt/config"
	"jrtt/function"
	"strconv"
	"time"
)

func main() {
	var config = config2.GetJrttUrl()
	page_size := config["page_size"].(int)
	for i:=0 ; i <= page_size ;i++{
		function.GetList(config["URL_LIST1"].(string)+strconv.Itoa(i))
		time.Sleep(1*time.Minute)
		function.GetList(config["URL_LIST2"].(string)+strconv.Itoa(i))
		time.Sleep(1*time.Minute)
	}
	function.Repost()
	function.Like()
	function.FollowUser()
}


