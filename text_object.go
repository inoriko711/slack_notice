package slack_notice

type TextObject struct {
	Type     string `json:"type"`
	Text     string `json:"text"`
	Emoji    bool   `json:"emoji,omitempty"`
	Verbatim bool   `json:"verbatim,omitempty"`
}

func SetTextObject(blocks []interface{}, text string, isMarkdown, isEmoji, isVerbatim bool) []interface{} {
	to := NewTextObject(text, isMarkdown, isEmoji, isVerbatim)
	blocks = append(blocks, to)
	return blocks
}

func NewTextObject(text string, isMarkdown, isEmoji, isVerbatim bool) *TextObject {
	if text == "" {
		// return nil, errors.New("text is empty")
		return nil
	}

	to := new(TextObject)
	if isMarkdown {
		if isEmoji {
			// return nil, errors.New("isEmoji field is only usable when isMarkdown is false")
			return nil
		}
		to.Type = "mrkdwn"
	} else {
		if isVerbatim {
			// return nil, errors.New("isVerbatim field is only usable when isMarkdown is true")
			return nil
		}
		to.Type = "plain_text"
	}
	to.Text = text
	to.Emoji = isEmoji
	to.Verbatim = isVerbatim
	return to
}
