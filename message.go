package dingrobot

const (
	msgTypeText     = "text"
	msgTypeLink     = "link"
	msgTypeMarkdown = "markdown"
)

type textMessage struct {
	MsgType string    `json:"msgtype"`
	Text    textParam `json:"text"`
	At      atParam   `json:"at"`
}

type linkMessage struct {
	MsgType string    `json:"msgtype"`
	Link    linkParam `json:"link"`
}

type markdownMessage struct {
	MsgType  string        `json:"msgtype"`
	Markdown markdownParam `json:"markdown"`
	At       atParam       `json:"at"`
}

type textParam struct {
	Content string `json:"content"`
}

type atParam struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	IsAtAll   bool     `json:"isAtAll,omitempty"`
}

type linkParam struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	MessageURL string `json:"messageUrl"`
	PicURL     string `json:"picUrl,omitempty"`
}

type markdownParam struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
