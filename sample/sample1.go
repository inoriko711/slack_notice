package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	. "github.com/inoriko711/slack_notice"
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

func main() {
	// ファイルから読み込む
	dataJSON, err := ioutil.ReadFile("../data/data.json")
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
