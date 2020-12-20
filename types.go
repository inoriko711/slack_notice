package slack_notice

// Block Elements の構造体
type ButtonElement struct {
	Type     string      `json:"type"`
	Text     interface{} `json:"text"`
	ActionId string      `json:"action_id,omitempty"`
	URL      string      `json:"url,omitempty"`
	Value    string      `json:"value,omitempty"`
	Style    string      `json:"style,omitempty"`
	Confirm  interface{} `json:"confirm,omitempty"`
}

type CheckboxGroups struct {
	Type           string        `json:"type"`
	ActionId       string        `json:"action_id"`
	Options        []interface{} `json:"options"`
	InitialOptions []interface{} `json:"initial_options,omitempty"`
	Confirm        interface{}   `json:"confirm,omitempty"`
}
type DatePickerElement struct {
	Type        string      `json:"type"`
	ActionId    string      `json:"action_id"`
	Placeholder interface{} `json:"placeholder,omitempty"`
	InitialDate string      `json:"initial_date,omitempty"`
	Confirm     interface{} `json:"confirm,omitempty"`
}

type ImageElement struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url"`
	AltText  string `json:"alt_text"`
}

type MultiSelectMenuElement struct {
	Type             string        `json:"type"`
	Placeholder      interface{}   `json:"placeholder"`
	ActionID         string        `json:"action_id"`
	Options          []interface{} `json:"options"`
	OptionGroups     []interface{} `json:"option_groups,omitempty"`
	InitialOptions   []interface{} `json:"initial_options,omitempty"`
	Confirm          interface{}   `json:"confirm,omitempty"`
	MaxSelectedItems uint          `json:"max_selected_items,omitempty"`
}

type ExternalDataSource struct {
	Type             string        `json:"type"`
	Placeholder      interface{}   `json:"placeholder"`
	ActionID         string        `json:"action_id"`
	MinQueryLength   int           `json:"min_query_length,omitempty"`
	InitialOptions   []interface{} `json:"initial_options,omitempty"`
	Confirm          interface{}   `json:"confirm,omitempty"`
	MaxSelectedItems uint          `json:"max_selected_items,omitempty"`
}

type UserList struct {
	Type             string      `json:"type"`
	Placeholder      interface{} `json:"placeholder"`
	ActionID         string      `json:"action_id"`
	InitialUsers     []string    `json:"initial_users,omitempty"`
	Confirm          interface{} `json:"confirm,omitempty"`
	MaxSelectedItems uint        `json:"max_selected_items,omitempty"`
}

type ConversationsList struct {
	Type                         string      `json:"type"`
	Placeholder                  interface{} `json:"placeholder"`
	ActionID                     string      `json:"action_id"`
	InitialConversations         []string    `json:"initial_conversations,omitempty"`
	DefaultToCurrentConversation bool        `json:"default_to_current_conversation,omitempty"`
	Confirm                      interface{} `json:"confirm,omitempty"`
	MaxSelectedItems             uint        `json:"max_selected_items,omitempty"`
	Filter                       interface{} `json:"filter,omitempty"`
}

type TextObject struct {
	Type     string `json:"type"`
	Text     string `json:"text"`
	Emoji    bool   `json:"emoji,omitempty"`
	Verbatim bool   `json:"verbatim,omitempty"`
}

// Blocks
type ContextBlock struct {
	Type     string        `json:"type"`
	Elements []interface{} `json:"elements"`
	BlockID  string        `json:"block_id,omitempty"`
}
type ActionBlock struct {
	Type     string        `json:"type"`
	Elements []interface{} `json:"elements"`
	BlockID  string        `json:"block_id,omitempty"`
}
type ImageBlock struct {
	Type     string      `json:"type"`
	ImageURL string      `json:"image_url"`
	AltText  string      `json:"alt_text"`
	Title    *TextObject `json:"title,omitempty"`
	BlockID  string      `json:"block_id,omitempty"`
}
type SectionBlocks struct {
	Type      string      `json:"type"`
	Text      *TextObject `json:"text"`
	BlockID   string      `json:"block_id,omitempty"`
	Fields    string      `json:"fields,omitempty"`
	Accessory interface{} `json:"accessory,omitempty"`
}
type HeaderBlock struct {
	Type    string      `json:"type"`
	Text    *TextObject `json:"text"`
	BlockID string      `json:"block_id,omitempty"`
}
