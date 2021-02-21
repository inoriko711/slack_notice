package slack_notice

type DividerBlock struct {
	Type    string `json:"type"`
	BlockID string `json:"block_id,omitempty"`
}

func SetDividerBlock(blocks []interface{}, blockID string) []interface{} {
	var dividerBlock DividerBlock

	dividerBlock.Type = "divider"
	if blockID != "" {
		dividerBlock.BlockID = blockID
	}

	blocks = append(blocks, dividerBlock)
	return blocks
}
