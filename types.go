package slack_notice

// Slack通知可変箇所のデータ構造体
type SlackApp struct {
	SlackURL            string           `json:"slack_url"`
	Username            string           `json:"username"`
	IconURL             string           `json:"icon_url"`
	SlackNoticeDataType *SlackNoticeData `json:"slack_notice_data"`
}

// Slack投稿内容可変箇所のデータ構造体
type SlackNoticeData struct {
	Text             string `json:"text"`
	HealingImageURL  string `json:"healing_image_url"`
	HealingImageText string `json:"healing_image_text"`
	AuthorImageURL   string `json:"author_image_url"`
	AuthorImageText  string `json:"author_image_text"`
	AuthorName       string `json:"author_name"`
}

// Block Kit の構造体
type TextObject struct {
	Type     string `json:"type"`
	Text     string `json:"text"`
	Emoji    bool   `json:"emoji,omitempty"`
	Verbatim bool   `json:"verbatim,omitempty"`
}
type ButtonElement struct {
	Type     string      `json:"type"`
	Text     interface{} `json:"text"`
	ActionId string      `json:"action_id,omitempty"`
	URL      string      `json:"url,omitempty"`
	Value    string      `json:"value,omitempty"`
	Style    string      `json:"style,omitempty"`
	Confirm  interface{} `json:"confirm,omitempty"`
}
type ContextBlock struct {
	Type     string        `json:"type"`
	Elements []interface{} `json:"elements"`
	BlockID  string        `json:"block_id,omitempty"`
}
type ActionBlock struct {
	Type     string        `json:"type"`
	Elements []interface{} `json:"elements"`
	BlockID  string        `json:"block_id,omitempty"`
}
type ImageBlock struct {
	Type     string      `json:"type"`
	ImageURL string      `json:"image_url"`
	AltText  string      `json:"alt_text"`
	Title    *TextObject `json:"title,omitempty"`
	BlockID  string      `json:"block_id,omitempty"`
}
type SectionBlocks struct {
	Type      string      `json:"type"`
	Text      *TextObject `json:"text"`
	BlockID   string      `json:"block_id,omitempty"`
	Fields    string      `json:"fields,omitempty"`
	Accessory interface{} `json:"accessory,omitempty"`
}
type HeaderBlock struct {
	Type    string      `json:"type"`
	Text    *TextObject `json:"text"`
	BlockID string      `json:"block_id,omitempty"`
}
