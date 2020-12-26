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

	blocks := []interface{}{
		&SectionBlocks{
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
				Options: []interface{}{
					&OptionObject{
						Text: TextObject{
							Type: "plain_text",
							Text: "*this is plain_text text*",
						},
						Value: "value-0",
					},
					&OptionObject{
						Text: TextObject{
							Type: "plain_text",
							Text: "*this is plain_text text*",
						},
						Value: "value-1",
					},
					&OptionObject{
						Text: TextObject{
							Type: "plain_text",
							Text: "*this is plain_text text*",
						},
						Value: "value-2",
					},
				},
			},
		},
	}

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
