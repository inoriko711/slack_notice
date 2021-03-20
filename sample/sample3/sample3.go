package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/inoriko711/slack_notice"
)

func main() {
	var blocks []interface{}
	blocks = slack_notice.SetHeaderBlock(blocks, "", slack_notice.NewTextObject("Budget Performance", false, false, false))
	blocks = slack_notice.SetDividerBlock(blocks, "")

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
