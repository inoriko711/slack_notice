package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	. "github.com/inoriko711/slack_notice"
)

func main() {
	item1TO := NewTextObject("Item1", false, false, false)

	blocks := []interface{}{
		&SectionBlock{
			Type:    "section",
			BlockID: "section678",
			Text: &TextObject{
				Type: "mrkdwn",
				Text: "Pick items from the list",
			},
			Accessory: &StaticOptions{
				ActionID: "text1234",
				Type:     "multi_static_select",
				Placeholder: &TextObject{
					Type: "plain_text",
					Text: "Select an item",
				},
				Options: []*OptionObject{
					{
						Text:  item1TO,
						Value: "value-0",
					},
					{
						Text: &TextObject{
							Type: "plain_text",
							Text: "Item2",
						},
						Value: "value-1",
					},
					{
						Text: &TextObject{
							Type: "plain_text",
							Text: "Item3",
						},
						Value: "value-2",
					},
				},
			},
		},
	}
	blocks = SetDividerBlock(blocks, "")
	blocks = SetImageBlock(
		blocks,
		"https://3.bp.blogspot.com/-d2L3hyx3JTU/WzC92mKYt1I/AAAAAAABM8Y/s76v5v1piCMN0eKy9jUEQlLHhCQcfmHMwCLcBGAs/s800/omatsuri_hashigonori.png",
		"はしご乗りのイラスト",
		"",
		NewTextObject("いらすとや", false, false, false),
	)
	bodyJSON, err := json.Marshal(map[string]interface{}{
		"username": "test",
		"text":     "sample",
		"blocks":   blocks,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	if bodyJSON != nil {
		fmt.Println(string(bodyJSON))
	}

	// 環境変数に設定しているSlackのWebhook URLを取得する
	slackURL := os.Getenv("SLACK_WEBHOOKS")
	resp, err := http.Post(slackURL, "application/json", bytes.NewReader(bodyJSON))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if _, err := io.Copy(ioutil.Discard, resp.Body); err != nil {
		fmt.Println(err)
	}
}
