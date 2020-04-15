package slack

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type (
	blockType  string
	formatType string
)

// Types of blocks
const (
	SectionBlock blockType = "section"
	DividerBlock blockType = "divider"
	ContextBlock blockType = "context"
)

// Types of text formatting
const (
	FormatMarkdown  formatType = "mrkdwn"
	FormatPlainText formatType = "plain_text"
)

type (
	// Message is a slack message composed of one or more Blocks
	Message struct {
		Text     string          `json:"text,omitempty"`
		Markdown bool            `json:"mrkdwn,omitempty"`
		Blocks   []*MessageBlock `json:"blocks"`
	}

	// MessageBlock is a block that may be included in a Message
	MessageBlock struct {
		Type blockType `json:"type"`

		// For Blocks of type "section", the following are defined:
		Text *MessageBlockText `json:"text,omitempty"`

		// For Blocks of type "context", the following are defined:
		Elements []*MessageBlockText `json:"elements,omitempty"`
	}

	// MessageBlockText defines a piece of text to be displayed in a block
	MessageBlockText struct {
		Type formatType `json:"type"`
		Text string     `json:"text"`

		// When Type is "plain_text", the following are defined:
		Emoji bool `json:"emoji,omitempty"`
	}
)

// Marshal performs a JSON marshal operation on the message to format it ready
// to send to slack
func (m Message) Marshal() ([]byte, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal slack message")
	}
	return data, nil
}
