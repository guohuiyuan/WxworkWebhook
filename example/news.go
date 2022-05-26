package main

import (
	"fmt"

	wxworkwebhook "github.com/guohuiyuan/WxworkWebhook"
)

func main() {
	w := wxworkwebhook.NewWebhook("693a91f6-7xxx-4bc4-97a0-0ec2sifa5aaa")
	a := wxworkwebhook.Article{
		Title:       "中秋节礼品领取",
		Description: "有一次下班坐公车，鼻子忽然有点痒，但微臣是文明之人，抠鼻也太不雅观了吧，就忍了忍，车过了一站，实在是没忍住就用手抠了起来，还拿出手机想挡一下，怎知边上一大妈见了大声说：“哎呦，小姑娘，在车上抠个鼻屎还要自拍呀”",
		URL:         "www.qq.com",
		PicURL:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
	}
	fmt.Println(w.Send(wxworkwebhook.News{Articles: []wxworkwebhook.Article{a}}))
}
