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
}
