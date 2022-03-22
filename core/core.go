package core

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
