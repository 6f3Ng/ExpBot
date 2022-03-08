package main

import (
	"ExpBot/core"
	"ExpBot/sender"
	"flag"
	"fmt"
	"os"
	"strconv"

	"gopkg.in/ini.v1"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "config.ini", "default config.ini")
	flag.Parse()
	cfg, err := ini.Load(configFile)
	if err != nil {
		fmt.Println("文件读取错误", err)
		os.Exit(1)
	}
	targetUrl := cfg.Section("sploitus").Key("targetUrl").String()
	startId := cfg.Section("sploitus").Key("startId").String()
	proxyUrl := cfg.Section("sploitus").Key("proxyUrl").String()
	for {
		statusCode := core.GetEXP(targetUrl, startId, proxyUrl)
		if statusCode != 200 {
			cfg.Section("sploitus").Key("startId").SetValue(startId)
			break
		} else {
			newId, _ := strconv.Atoi(startId)
			startId = strconv.Itoa(newId + 1)
		}
	}

	dingtalkUrl := cfg.Section("dingtalk").Key("targetUrl").String()
	dingtalkAccessTokens := cfg.Section("dingtalk").Key("accessTokens").String()
	dingtalkKeyword := cfg.Section("dingtalk").Key("keyword").String()
	dingtalkProxyUrl := cfg.Section("dingtalk").Key("proxyUrl").String()
	notifyIfNotFound := cfg.Section("dingtalk").Key("notifyIfNotFound").MustBool(true)

	sender.SendToDingtalk(dingtalkUrl, dingtalkAccessTokens, dingtalkKeyword, dingtalkProxyUrl, notifyIfNotFound)
	// for _, v := range core.GetInfo() {
	// 	fmt.Println(v.Title, v.TargetUrl)
	// }
	cfg.SaveTo(configFile)

	// fmt.Println(cfg.Section("sploitus").Key("targetUrl"))
	// fmt.Println(cfg.Section("sploitus").Key("proxyUrl"))
	// fmt.Println(cfg.Section("sploitus").Key("startId"))

	// fmt.Println(cfg.Section("dingtalk").Key("targetUrl"))
	// fmt.Println(cfg.Section("dingtalk").Key("proxyUrl"))
	// fmt.Println(cfg.Section("dingtalk").Key("accessToken"))
	// fmt.Println(cfg.Section("dingtalk").Key("keyword"))
}
