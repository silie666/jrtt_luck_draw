## 今日头条抽奖

## 看自己是什么环境，修改function/func.go下的chromedp.ExecPath("/usr/bin/google-chrome"),需要安装google-chrome

## 修改config/jrtt.go配置,都可以在cookie里面找到
#### urlConfig["JR_UID"] = ""
#### urlConfig["SESSIONID"] = ""
#### urlConfig["TTWID"] = ""
#### urlConfig["CSRF_SESSION_ID"] = ""   需要点赞或者转发才会出现在cookie里面，头条用js加密了
#### urlConfig["X-SECSDK-CSRF-TOKEN"] = "" 在点赞或者转到请求连接的请求头里




## 修改config/db.go数据库配置
#### dbConfig["DB_HOST"] = "" 数据库地址
#### dbConfig["DB_PORT"] = "" 端口
#### dbConfig["DB_NAME"] = "" 数据库名称
#### dbConfig["DB_USER"] = "" 账号
#### dbConfig["DB_PWD"] = "" 密码


## JR_UID可以这样找
![1642406340](https://user-images.githubusercontent.com/38691833/149730223-372f8567-cc0f-4d1b-9fb9-858e7f4f33fb.jpg)

## X-SECSDK-CSRF-TOKEN可以这样找
![image](https://user-images.githubusercontent.com/38691833/159616997-c2d770e8-81b1-45cb-8260-d6a3ded18cb5.png)

