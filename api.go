package wxworkwebhook

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
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

func (w *Webhook) Send(msg interface{}) (response CommonResponse, err error) {
	v := reflect.TypeOf(msg)
	r := MsgReq{}
	switch v.Name() {
	case "File":
		r.Msgtype = MsgFile
		r.File = msg.(File)
	case "Image":
		r.Msgtype = MsgImage
		r.Image = msg.(Image)
	case "Markdown":
		r.Msgtype = MsgMarkdown
		r.Markdown = msg.(Markdown)
	case "News":
		r.Msgtype = MsgNews
		r.News = msg.(News)
	case "Text":
		r.Msgtype = MsgText
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

func (w *Webhook) SendFile(path string) (response CommonResponse, err error) {
	if IsNetFile(path) {
		err = errors.New("请发送本地路径")
		return
	}
	umr, err := w.UploadMedia(path)
	if err != nil {
		return
	}
	response, err = w.Send(File{
		MediaID: umr.MediaID,
	})
	return
}

func (w *Webhook) SendImage(path string) (response CommonResponse, err error) {
	if IsNetFile(path) {
		err = errors.New("请发送本地路径")
		return
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}
	picBase64 := base64.StdEncoding.EncodeToString(data)
	sign := fmt.Sprintf("%x", md5.Sum(data))
	response, err = w.Send(Image{Base64: picBase64, Md5: sign})
	return
}

func (w *Webhook) SendText(text string) (response CommonResponse, err error) {
	response, err = w.Send(Text{Content: text})
	return
}

func (w *Webhook) SendMarkdown(text string) (response CommonResponse, err error) {
	response, err = w.Send(Markdown{Content: text})
	return
}

func (w *Webhook) SendNews(title, description, jumpURL, picURL string) (response CommonResponse, err error) {
	a := Article{
		Title:       title,
		Description: description,
		URL:         jumpURL,
		PicURL:      picURL,
	}
	response, err = w.Send(News{Articles: []Article{a}})
	return
}
