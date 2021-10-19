package slack_test

import (
	"testing"

	"github.com/ONSdigital/who-goes-there/pkg/slack"
	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {

	sm := slack.Message{
		Text: "Who Goes There report",
		Blocks: []*slack.MessageBlock{
			{
				Type: slack.SectionBlock,
				Text: &slack.MessageBlockText{
					Type: slack.FormatMarkdown,
					Text: "Here's your report from the recent run of _Who Goes There?_",
				},
			},
			{
				Type: slack.DividerBlock,
			},
			{
				Type: slack.ContextBlock,
				Elements: []*slack.MessageBlockText{
					{
						Type: slack.FormatMarkdown,
						Text: "This is your report detail",
					},
				},
			},
		},
	}

	expected := `{"text":"Who Goes There report","blocks":[{"type":"section","text":{"type":"mrkdwn","text":"Here's your report from the recent run of _Who Goes There?_"}},{"type":"divider"},{"type":"context","elements":[{"type":"mrkdwn","text":"This is your report detail"}]}]}`

	marshaled, err := sm.Marshal()
	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(marshaled))
}
