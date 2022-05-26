package main

import (
	"fmt"

	wxworkwebhook "github.com/guohuiyuan/WxworkWebhook"
)

func main() {
	w := wxworkwebhook.NewWebhook("693a91f6-7xxx-4bc4-97a0-0ec2sifa5aaa")
	fmt.Println(w.Send(wxworkwebhook.Markdown{Content: `实时新增用户反馈<font color=\"warning\">132例</font>，请相关同事注意。\n
	>类型:<font color=\"comment\">用户反馈</font>
	>普通用户反馈:<font color=\"comment\">117例</font>
	>VIP用户反馈:<font color=\"comment\">15例</font>`}))
}
