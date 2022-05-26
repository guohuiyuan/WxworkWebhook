package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"os"

	wxworkwebhook "github.com/guohuiyuan/WxworkWebhook"
)

func main() {
	data, err := os.ReadFile("example/黒金.jpg")
	if err != nil {
		fmt.Println(err)
	}
	picBase64 := base64.StdEncoding.EncodeToString(data)
	sign := fmt.Sprintf("%x", md5.Sum(data))
	w := wxworkwebhook.NewWebhook("693a91f6-7xxx-4bc4-97a0-0ec2sifa5aaa")
	fmt.Println(w.Send(wxworkwebhook.Image{Base64: picBase64, Md5: sign}))
}
