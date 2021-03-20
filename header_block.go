package slack_notice

type HeaderBlock struct {
	Type    string      `json:"type"`
	Text    *TextObject `json:"text"`
	BlockID string      `json:"block_id,omitempty"`
}

func SetHeaderBlock(blocks []interface{}, blockID string, text *TextObject) []interface{} {
	hb := NewHeaderBlock(blockID, text)
	return append(blocks, hb)
}

func NewHeaderBlock(blockID string, text *TextObject) *HeaderBlock {
	if text == nil || text.Type == "mrkdwn" {
		return nil
	}
	hb := new(HeaderBlock)
	hb.Type = "header"
	hb.Text = text
	hb.BlockID = blockID
	return hb
}
