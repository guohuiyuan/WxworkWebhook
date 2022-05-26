package main

import (
	"fmt"

	wxworkwebhook "github.com/guohuiyuan/WxworkWebhook"
)

func main() {
	w := wxworkwebhook.NewWebhook("693a91f6-7xxx-4bc4-97a0-0ec2sifa5aaa")
	fmt.Println(w.SendFile("example/黒金.jpg"))
	fmt.Println(w.SendImage("example/黒金.jpg"))
	fmt.Println(w.SendImage("http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png"))
	fmt.Println(w.SendText("大家要好好吃饭"))
	fmt.Println(w.SendMarkdown("# 大家要好好吃饭"))
	fmt.Println(w.SendNews("晚上好", "大家要好好吃饭", "https://b23.tv/BV1bz4y1z7uu", "http://i2.hdslb.com/bfs/archive/d04b95ba3658a39f594fc60379eaf84385126a81.png"))
}
