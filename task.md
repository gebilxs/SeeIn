第一周：基础学习，并完成命令行工具 完成命令行工具与学习任务 命令行工具


完成upnp协议location的api请求，将响应信息结构化展示

命令行输入：支持 -url 参数。例子： ./upnp -url http://175.205.17.185:5531/rss/Starter_desc.xml

输出：跟shodan一样UPnP Device:

Device Type: urn:d-link-com:D-Link-Starter:1

Friendly Name: Wireless N300 Router

Model Name: D-Link DIR-604M

Model Number: DIR-604M

Model Description: ADSL

Model URL: http://www.d-link.co.kr/

Manufacturer: D-Link Corporation

Manufacturer URL: http://www.d-link.co.kr/

Serial Number: v1.00KR

UDN: uuid:00000000-0008-2fe4-0400-4000009b4001

Presentation URL: http://192.168.1.1/

Service #1:
Service Type: urn:d-link-com:service:D-Link-Starter:1
案例:
https://fofa.info/result?qbase64=cHJvdG9jb2w9InVwbnAi
https://www.shodan.io/search?query=%22D-Link+DIR-806A%22
学习任务：
golang基础学习，如：基础语法，编译，调试等
熟悉 Fofa 的搜索语法和功能，理解 Fofa 规则的概念和作用，并录入一条新的规则，并完成规则审核员审核
参考资料：
Fofa 官方文档：可以访问 Fofa 官方网站，查找官方文档和教程，了解 Fofa 的功能和使用方法。
规则手册：   https://ones.baimaohui.net:24688/wiki/#/team/2hTgeDe2/space/Gt4DCMwE/page/NQtUkKPu
网络教程： https://nosec.org/home/detail/5060.html
https://tour.go-zh.org/welcome/1