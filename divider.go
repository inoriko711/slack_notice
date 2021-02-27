package slack_notice

type DividerBlock struct {
	Type    string `json:"type"`
	BlockID string `json:"block_id,omitempty"`
}

func SetDividerBlock(blocks []interface{}, blockID string) []interface{} {
	dividerBlock := NewDividerBlock(blockID)

	blocks = append(blocks, dividerBlock)
	return blocks
}

func NewDividerBlock(blockID string) *DividerBlock {
	dividerBlock := new(DividerBlock)
	dividerBlock.Type = "divider"
	if blockID != "" {
		dividerBlock.BlockID = blockID
	}
	return dividerBlock
}
