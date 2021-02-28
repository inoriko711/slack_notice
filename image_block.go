package slack_notice

type ImageBlock struct {
	Type     string      `json:"type"`
	ImageURL string      `json:"image_url"`
	AltText  string      `json:"alt_text"`
	Title    *TextObject `json:"title,omitempty"`
	BlockID  string      `json:"block_id,omitempty"`
}

func SetImageBlock(blocks []interface{}, imageURL, altText, blockID string, title *TextObject) []interface{} {
	ib := NewImageBlock(imageURL, altText, blockID, title)

	return append(blocks, ib)
}

func NewImageBlock(imageURL, altText, blockID string, title *TextObject) *ImageBlock {
	if imageURL == "" || altText == "" {
		return nil
	}
	ib := new(ImageBlock)
	ib.Type = "image"
	ib.ImageURL = imageURL
	ib.AltText = altText
	ib.Title = title
	if blockID != "" {
		ib.BlockID = blockID
	}

	return ib
}
