package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

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

func main() {
	// ファイルから読み込む
	dataJSON, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// ファイルから読み込んだデータをSlackApp型で持つ
	var data SlackApp
	err = json.Unmarshal(dataJSON, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// BlockKitBuilderで作ったblocks配列の中身を準備する
	blocks := buildBlocks(*data.SlackNoticeDataType)

	// Incoming Webhookに渡すBodyを作成
	bodyJSON, err := json.Marshal(map[string]interface{}{
		"username": data.Username,
		"icon_url": data.IconURL,
		"text":     "癒し画像のお届け",
		"blocks":   blocks,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	if bodyJSON != nil {
		fmt.Print(string(bodyJSON))
	}

	resp, err := http.Post(data.SlackURL, "application/json", bytes.NewReader(bodyJSON))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if _, err := io.Copy(ioutil.Discard, resp.Body); err != nil {
		fmt.Println(err)
	}
}

func buildBlocks(data SlackNoticeData) []interface{} {

	return []interface{}{
		&HeaderBlock{
			Type: "header",
			Text: &TextObject{
				Type: "plain_text",
				Text: ":star2:癒しタイム:star2:",
			},
		},
		&SectionBlocks{
			Type: "section",
			Text: &TextObject{
				Type: "mrkdwn",
				Text: data.Text,
			},
		},
		&ImageBlock{
			Type:     "image",
			ImageURL: data.HealingImageURL,
			AltText:  data.HealingImageText,
		},
		&ActionBlock{
			Type: "actions",
			Elements: []interface{}{
				&ButtonElement{
					Type: "button",
					Text: &TextObject{
						Type:  "plain_text",
						Text:  "inoriko",
						Emoji: true,
					},
					URL:   "https://twitter.com/inoriko711",
					Value: "inoriko Twitter",
				},
			},
		},
		&ContextBlock{
			Type: "context",
			Elements: []interface{}{
				&ImageBlock{
					Type:     "image",
					ImageURL: data.AuthorImageURL,
					AltText:  data.AuthorImageText,
				},
				&TextObject{
					Type:  "plain_text",
					Text:  fmt.Sprintf("Author: %s", data.AuthorName),
					Emoji: true,
				},
			},
		},
	}
}
