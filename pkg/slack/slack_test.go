package slack_test

import (
	"testing"

	"github.com/ONSdigital/who-goes-there/pkg/slack"
	"github.com/matryer/is"
)

func TestMarshal(t *testing.T) {

	is := is.New(t)

	sm := slack.Message{
		Text: &slack.MessageBlockText{
			Type: slack.FormatPlainText,
			Text: "Who Goes There report",
		},
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

	expected := []byte(`{"text":{"type":"plain_text","text":"Who Goes There report"},"blocks":[{"type":"section","text":{"type":"mrkdwn","text":"Here's your report from the recent run of _Who Goes There?_"}},{"type":"divider"},{"type":"context","elements":[{"type":"mrkdwn","text":"This is your report detail"}]}]}`)

	marshaled, err := sm.Marshal()
	is.NoErr(err)
	is.Equal(marshaled, expected)
}
