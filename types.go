package wxworkwebhook

type UploadMediaResponse struct {
	CommonResponse
	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

type CommonResponse struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

const (
	MsgFile     = "file"
	MsgImage    = "image"
	MsgMarkdown = "markdown"
	MsgNews     = "news"
	MsgText     = "text"
)

type MsgReq struct {
	Msgtype  string   `json:"msgtype"`
	File     File     `json:"file,omitempty"`
	Image    Image    `json:"image,omitempty"`
	Markdown Markdown `json:"markdown,omitempty"`
	News     News     `json:"news,omitempty"`
	Text     Text     `json:"text,omitempty"`
}

type News struct {
	Articles []Article `json:"articles"`
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"picurl"`
}

type Image struct {
	Base64 string `json:"base64"`
	Md5    string `json:"md5"`
}

type Text struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

type Markdown struct {
	Content string `json:"content"`
}

type File struct {
	MediaID string `json:"media_id"`
}
