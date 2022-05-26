package wxworkwebhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"reflect"

	"github.com/FloatTech/zbputils/web"
)

type Webhook struct {
	WebhookKey string
}

func NewWebhook(key string) *Webhook {
	return &Webhook{
		WebhookKey: key,
	}
}

const (
	baseURL        = "https://qyapi.weixin.qq.com"
	uploadMediaURL = baseURL + "/cgi-bin/webhook/upload_media?key=%s&type=file"
	sendURL        = baseURL + "/cgi-bin/webhook/send?key=%s"
)

// UploadMedia 上传文件
func (w *Webhook) UploadMedia(path string) (response UploadMediaResponse, err error) {
	requestURL := fmt.Sprintf(uploadMediaURL, w.WebhookKey)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	part1, err := writer.CreateFormFile("media", filepath.Base(path))
	_, err = io.Copy(part1, file)
	if err != nil {
		return
	}
	err = writer.Close()
	if err != nil {
		return
	}
	data, err := web.PostData(requestURL, writer.FormDataContentType(), payload)
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &response); err != nil {
		return
	}
	if response.Errcode != 0 {
		err = errors.New(response.Errmsg)
	}
	return
}

// Send 发送消息
func (w *Webhook) Send(msg interface{}) (response CommonResponse, err error) {
	v := reflect.TypeOf(msg)
	r := BotMsgReq{}
	switch v.Name() {
	case "File":
		r.Msgtype = BotMsgFile
		r.File = msg.(File)
	case "Image":
		r.Msgtype = BotMsgImage
		r.Image = msg.(Image)
	case "Markdown":
		r.Msgtype = BotMsgMarkdown
		r.Markdown = msg.(Markdown)
	case "News":
		r.Msgtype = BotMsgNews
		r.News = msg.(News)
	case "Text":
		r.Msgtype = BotMsgText
		r.Text = msg.(Text)
	default:
		err = errors.New("msg的类型不合法,合法的类型有File,Image,Markdown,News,Text")
		return
	}
	requestURL := fmt.Sprintf(sendURL, w.WebhookKey)
	b, _ := json.Marshal(r)
	data, err := web.PostData(requestURL, "application/json", bytes.NewReader(b))
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &response)
	return
}

// SendPathFile 通过本地路径，上传并且发送文件（也许这个不应该提供）
func (w *Webhook) SendPathFile(path string) (response CommonResponse, err error) {
	umr, err := w.UploadMedia(path)
	if err != nil {
		return
	}
	response, err = w.Send(File{
		MediaID: umr.MediaID,
	})
	return
}
