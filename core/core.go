package core

import (
	"log"
	"net/http"
	"net/url"

	"github.com/opesun/goquery"
)

type Info struct {
	Title     string
	TargetUrl string
}

func (info *Info) Set(title string, targetUrl string) {
	info.Title = title
	info.TargetUrl = targetUrl
}

func GetInfo() []Info {
	return respInfo
}

var respInfo []Info

func GetEXP(targetUrl string, startId string, proxyUrl string) int {
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
	expUrl := targetUrl + startId
	req, err := http.NewRequest("GET", expUrl, nil)
	if err != nil {
		log.Fatal(err)
		// handle error
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:96.0) Gecko/20100101 Firefox/96.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
		// handle error
	}

	defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return resp.StatusCode
	}

	body, err := goquery.Parse(resp.Body)

	if err != nil {
		log.Fatal(err)
		// handle error
	}
	title := body.Find("title").Text()
	var i Info
	i.Set(title, expUrl)
	respInfo = append(respInfo, i)
	// fmt.Println(title)
	return resp.StatusCode
}
