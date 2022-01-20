# ExpBot
## 功能
---
监控 https://sploitus.com/ 上新发布的exp，并进行钉钉告警

## 配置方法
---
配置`config.ini`文件：
```ini
[sploitus] # 监控项
targetUrl = https://sploitus.com/exploit?id=EDB-ID: # 监控地址
proxyUrl  = http://127.0.0.1:8080 # 是否通过代理访问，支持http和socks，若留空则不通过代理
startId   = 50679 # 起始EDB-ID

[dingtalk] # 钉钉告警
targetUrl    = https://oapi.dingtalk.com/robot/send?access_token= # 钉钉告警地址
accessTokens = aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa,bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb # 钉钉告警的access_token，使用逗号分隔多个token
keyword      = 新EXP推送 # 钉钉告警的关键字配置
proxyUrl     = http://127.0.0.1:8080 # 是否通过代理告警，支持http和socks，若留空则不通过代理
```
usage
---
直接运行
``` shell
./ExpBot -config config.ini
```
配置crontab定时任务
``` shell
0 14 * * * /root/GetEXP/ExpBot -config /root/GetEXP/config.ini
```
todo list
---
1. 增加其他告警方式
2. 增加其他监控项，如cve监控，cnvd监控等
