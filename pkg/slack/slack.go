package slack

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

var client *http.Client

type (
	// Using some custom types for syntactic sugar purposes
	blockType  string
	formatType string
)

// Types of blocks
const (
	HeaderBlock  blockType = "header"
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

func (m Message) Post(ctx context.Context, webhookURL string) error {
	if client == nil {
		client = &http.Client{
			Timeout: time.Second * 5,
		}
	}

	data, err := m.Marshal()
	if err != nil {
		return err
	}

	log.Printf("Posting slack message body: %s", string(data))

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		webhookURL,
		bytes.NewBuffer(data),
	)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		e := fmt.Errorf("non-200 status returned from slack: %d", resp.StatusCode)

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.WithMessagef(e, "error parsing response body: %v", err)
		}
		return errors.WithMessagef(e, "response: %s", response)
	}
	return nil
}
