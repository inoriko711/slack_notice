package slack_notice

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
