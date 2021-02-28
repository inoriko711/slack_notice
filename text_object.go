package slack_notice

import "errors"

type TextObject struct {
	Type     string `json:"type"`
	Text     string `json:"text"`
	Emoji    bool   `json:"emoji,omitempty"`
	Verbatim bool   `json:"verbatim,omitempty"`
}

func SetTextObject(blocks []interface{}, text string, isMarkdown, isEmoji, isVerbatim bool) ([]interface{}, error) {
	to, err := NewTextObject(text, isMarkdown, isEmoji, isVerbatim)
	if err != nil {
		return nil, err
	}

	blocks = append(blocks, to)
	return blocks, nil
}

func NewTextObject(text string, isMarkdown, isEmoji, isVerbatim bool) (*TextObject, error) {
	if text == "" {
		return nil, errors.New("text is empty")
	}

	to := new(TextObject)
	if isMarkdown {
		if isEmoji {
			return nil, errors.New("isEmoji field is only usable when isMarkdown is false")
		}
		to.Type = "mrkdwn"
	} else {
		if isVerbatim {
			return nil, errors.New("isVerbatim field is only usable when isMarkdown is true")
		}
		to.Type = "plain_text"
	}
	to.Text = text
	to.Emoji = isEmoji
	to.Verbatim = isVerbatim
	return to, nil
}
