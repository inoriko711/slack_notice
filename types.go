package slack_notice

// Element の構造体
type ButtonElement struct {
	Type     string                    `json:"type"`
	Text     *TextObject               `json:"text"`
	ActionId string                    `json:"action_id,omitempty"`
	URL      string                    `json:"url,omitempty"`
	Value    string                    `json:"value,omitempty"`
	Style    string                    `json:"style,omitempty"`
	Confirm  *ConfirmationDialogObject `json:"confirm,omitempty"`
}

func NewButtonElement() *ButtonElement {
	return &ButtonElement{
		Type: "button",
		Text: &TextObject{
			Type: "plain_text",
		},
	}
}

type CheckboxGroups struct {
	Type           string                    `json:"type"`
	ActionId       string                    `json:"action_id"`
	Options        []*OptionObject           `json:"options"`
	InitialOptions []*OptionObject           `json:"initial_options,omitempty"`
	Confirm        *ConfirmationDialogObject `json:"confirm,omitempty"`
}

func NewCheckboxGroups() *CheckboxGroups {
	return &CheckboxGroups{
		Type: "checkboxes",
	}
}

type DatePickerElement struct {
	Type        string                    `json:"type"`
	ActionId    string                    `json:"action_id"`
	Placeholder *TextObject               `json:"placeholder,omitempty"`
	InitialDate string                    `json:"initial_date,omitempty"`
	Confirm     *ConfirmationDialogObject `json:"confirm,omitempty"`
}

func NewDatePickerElement() *DatePickerElement {
	return &DatePickerElement{
		Type: "datepicker",
	}
}

type ImageElement struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url"`
	AltText  string `json:"alt_text"`
}

func NewImageElement() *ImageElement {
	return &ImageElement{
		Type: "image",
	}
}

// ---MultiSelectMenuElement, Select menu element
type StaticOptions struct {
	Type             string                    `json:"type"`
	Placeholder      *TextObject               `json:"placeholder"`
	ActionID         string                    `json:"action_id"`
	Options          []*OptionObject           `json:"options"`
	OptionGroups     []*OptionObject           `json:"option_groups,omitempty"`
	InitialOption    interface{}               `json:"initial_option,omitempty"`
	InitialOptions   []*OptionObject           `json:"initial_options,omitempty"`
	Confirm          *ConfirmationDialogObject `json:"confirm,omitempty"`
	MaxSelectedItems uint                      `json:"max_selected_items,omitempty"`
}

type ExternalDataSource struct {
	Type             string                    `json:"type"`
	Placeholder      TextObject                `json:"placeholder"`
	ActionID         string                    `json:"action_id"`
	MinQueryLength   uint                      `json:"min_query_length,omitempty"`
	InitialOption    *OptionObject             `json:"initial_option,omitempty"`
	Confirm          *ConfirmationDialogObject `json:"confirm,omitempty"`
	MaxSelectedItems uint                      `json:"max_selected_items,omitempty"`
}

type UserList struct {
	Type             string      `json:"type"`
	Placeholder      TextObject  `json:"placeholder"`
	ActionID         string      `json:"action_id"`
	InitialUsers     []string    `json:"initial_users,omitempty"`
	InitialUser      string      `json:"initial_user,omitempty"`
	Confirm          interface{} `json:"confirm,omitempty"`
	MaxSelectedItems uint        `json:"max_selected_items,omitempty"`
}

type ConversationsList struct {
	Type                         string      `json:"type"`
	Placeholder                  TextObject  `json:"placeholder"`
	ActionID                     string      `json:"action_id"`
	InitialConversations         []string    `json:"initial_conversations,omitempty"`
	InitialConversation          string      `json:"initial_conversation,omitempty"`
	DefaultToCurrentConversation bool        `json:"default_to_current_conversation,omitempty"`
	Confirm                      interface{} `json:"confirm,omitempty"`
	MaxSelectedItems             uint        `json:"max_selected_items,omitempty"`
	ResponseURLEnabled           bool        `json:"response_url_enabled,omitempty"`
	Filter                       interface{} `json:"filter,omitempty"`
}

type PublicChannelsList struct {
	Type               string      `json:"type"`
	Placeholder        TextObject  `json:"placeholder"`
	ActionID           string      `json:"action_id"`
	InitialChannels    []string    `json:"initial_channels,omitempty"`
	Confirm            interface{} `json:"confirm,omitempty"`
	MaxSelectedItems   uint        `json:"max_selected_items,omitempty"`
	ResponseUrlEnabled bool        `json:"response_url_enabled,omitempty"`
}

// ---

type OverflowMenuElement struct {
	Type     string                    `json:"type"`
	ActionID string                    `json:"action_id"`
	Options  []*OptionObject           `json:"options"`
	Confirm  *ConfirmationDialogObject `json:"confirm,omitempty"`
}

func NewOverflowMenuElement() *OverflowMenuElement {
	return &OverflowMenuElement{
		Type: "overflow",
	}
}

type PlainTextInputElement struct {
	Type                 string                       `json:"type"`
	ActionID             string                       `json:"action_id"`
	Placeholder          *TextObject                  `json:"placeholder,omitempty"`
	InitialValue         string                       `json:"initial_value,omitempty"`
	Multiline            bool                         `json:"multiline,omitempty"`
	MinLength            uint                         `json:"min_length,omitempty"`
	MaxLength            uint                         `json:"max_length,omitempty"`
	DispatchActionConfig *DispatchActionConfiguration `json:"dispatch_action_config,omitempty"`
}

func NewPlainTextInputElement() *PlainTextInputElement {
	return &PlainTextInputElement{
		Type: "plain_text_input",
	}
}

type RadioButtonGroupElement struct {
	Type          string                    `json:"type"`
	ActionID      string                    `json:"action_id"`
	Options       []*OptionObject           `json:"options"`
	InitialOption *OptionObject             `json:"initial_option,omitempty"`
	Confirm       *ConfirmationDialogObject `json:"confirm,omitempty"`
}

func NewRadioButtonGroupElement() *RadioButtonGroupElement {
	return &RadioButtonGroupElement{
		Type: "radio_buttons",
	}
}

type TimePickerElement struct {
	Type        string                    `json:"type"`
	ActionID    string                    `json:"action_id"`
	Placeholder *TextObject               `json:"placeholder,omitempty"`
	InitialTime string                    `json:"initial_time,omitempty"`
	Confirm     *ConfirmationDialogObject `json:"confirm,omitempty"`
}

func NewTimePickerElement() *TimePickerElement {
	return &TimePickerElement{
		Type: "timepicker",
	}
}

// Object
type TextObject struct {
	Type     string `json:"type"`
	Text     string `json:"text"`
	Emoji    bool   `json:"emoji,omitempty"`
	Verbatim bool   `json:"verbatim,omitempty"`
}

type ConfirmationDialogObject struct {
	Title   *TextObject `json:"title"`
	Text    *TextObject `json:"text"`
	Confirm *TextObject `json:"confirm"`
	Deny    *TextObject `json:"deny"`
	Style   string      `json:"style,omitempty"`
}

type OptionObject struct {
	Text        *TextObject `json:"text"`
	Value       string      `json:"value"`
	Description *TextObject `json:"description,omitempty"`
	URL         string      `json:"url,omitempty"`
}

type OptionGroupObject struct {
	Label   *TextObject     `json:"label"`
	Options []*OptionObject `json:"options"`
}

type DispatchActionConfiguration struct {
	TriggerActionsOn []string `json:"trigger_actions_on,omitempty"`
}

type FilterObjectForConversationLists struct {
	Include                       []string `json:"include,omitempty"`
	ExcludeExternalSharedChannels bool     `json:"exclude_external_shared_channels,omitempty"`
	ExcludeBotUsers               bool     `json:"exclude_bot_users,omitempty"`
}

// Blocks
type ActionBlock struct {
	Type     string        `json:"type"`
	Elements []interface{} `json:"elements"`
	BlockID  string        `json:"block_id,omitempty"`
}
type ContextBlock struct {
	Type     string        `json:"type"`
	Elements []interface{} `json:"elements"`
	BlockID  string        `json:"block_id,omitempty"`
}

type DividerBlock struct {
	Type    string `json:"type"`
	BlockID string `json:"block_id,omitempty"`
}

type FileBlock struct {
	Type       string `json:"type"`
	ExternalId string `json:"external_id"`
	Source     string `json:"source"`
	BlockID    string `json:"block_id,omitempty"`
}

type HeaderBlock struct {
	Type    string      `json:"type"`
	Text    *TextObject `json:"text"`
	BlockID string      `json:"block_id,omitempty"`
}

type ImageBlock struct {
	Type     string      `json:"type"`
	ImageURL string      `json:"image_url"`
	AltText  string      `json:"alt_text"`
	Title    *TextObject `json:"title,omitempty"`
	BlockID  string      `json:"block_id,omitempty"`
}

type InputBlock struct {
	Type           string      `json:"type"`
	Label          TextObject  `json:"label"`
	Element        interface{} `json:"element"`
	DispatchAction bool        `json:"dispatch_action,omitempty"`
	BlockID        string      `json:"block_id,omitempty"`
	Hint           *TextObject `json:"hint,omitempty"`
	Optional       bool        `json:"optional,omitempty"`
}

type SectionBlock struct {
	Type      string        `json:"type"`
	Text      *TextObject   `json:"text"`
	BlockID   string        `json:"block_id,omitempty"`
	Fields    []*TextObject `json:"fields,omitempty"`
	Accessory interface{}   `json:"accessory,omitempty"`
}
