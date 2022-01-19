package sender

import (
	"ExpBot/core"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func SendToDingtalk(targetUrl string, accessTokens string, keyword string, proxyUrl string) {
	sendJson := `
{
	"msgtype": "markdown",
	"markdown": {
		"title":"%s",
		"text": "#### %s \n %s \n"
	},
		"at": {
			"atMobiles": [],
			"atUserIds": [],
			"isAtAll": false
		}
}`
	expInfo := ""
	var sendData string
	if len(core.GetInfo()) == 0 {
		sendData = fmt.Sprintf(sendJson, keyword, keyword, "> 今日无新exp呀   \n---   \n")
	} else {
		for _, v := range core.GetInfo() {
			expInfo += fmt.Sprintf("> [%s](%s)   \n---   \n", v.Title, v.TargetUrl)
		}
		sendData = fmt.Sprintf(sendJson, keyword, keyword, expInfo)
	}
	for _, accessToken := range strings.Split(accessTokens, ",") {
		if accessToken == "" {
			continue
		}
		body := bytes.NewBuffer([]byte(sendData))

		dingtalkUrl := targetUrl + accessToken
		httpClient := &http.Client{}
		if proxyUrl != "" {
			proxy := func(_ *http.Request) (*url.URL, error) {
				return url.Parse(proxyUrl)
			}
			httpTransport := &http.Transport{
				Proxy: proxy,
			}
			httpClient = &http.Client{
				Transport: httpTransport,
			}
		}
		req, err := http.NewRequest("POST", dingtalkUrl, body)
		if err != nil {
			log.Fatal(err)
			// handle error
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := httpClient.Do(req)
		if err != nil {
			log.Fatal(err)
			// handle error
		}

		defer resp.Body.Close()
	}
}
