package main

import (
	"fmt"

	wxworkwebhook "github.com/guohuiyuan/WxworkWebhook"
)

func main() {
	w := wxworkwebhook.NewWebhook("693a91f6-7xxx-4bc4-97a0-0ec2sifa5aaa")
	fmt.Println(w.Send(wxworkwebhook.File{MediaID: "3V2NQKuV8aUdRjYIN-Lfr6adJvdoOPvKZ0WwQ7DPmxGhRmeY8HBWw3DKBlElWcheP"}))
}
