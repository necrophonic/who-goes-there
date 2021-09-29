package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ONSdigital/who-goes-there/pkg/report"
	"github.com/ONSdigital/who-goes-there/pkg/slack"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// WebhookSpec contains the values required for connecting to the slack webhook
type WebhookSpec struct {
	URL string
}

type (
	messageBody struct {
		ResponsePayload report.Report `json:"responsePayload"`
	}
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {

	var s WebhookSpec
	err := envconfig.Process("WEBHOOK", &s)
	if err != nil {
		return err
	}

	for _, message := range sqsEvent.Records {
		var m messageBody
		err := json.Unmarshal([]byte(message.Body), &m)
		if err != nil {
			return errors.Wrap(err, "failed to unmarshal report data")
		}

		r := m.ResponsePayload

		message := slack.Message{
			Text: "New report from Who Goes There",
			Blocks: []*slack.MessageBlock{
				{
					Type: slack.HeaderBlock,
					Text: &slack.MessageBlockText{
						Type: slack.FormatMarkdown,
						Text: "Here's your report from the recent run of _Who Goes There?_",
					},
				},
				{
					Type: slack.DividerBlock,
				},
				{
					Type: slack.SectionBlock,
					Text: &slack.MessageBlockText{
						Type: slack.FormatMarkdown,
						Text: r.SummaryTableMarkdown(),
					},
				},
				{
					Type: slack.ContextBlock,
					Elements: []*slack.MessageBlockText{
						{
							Type: slack.FormatMarkdown,
							Text: fmt.Sprintf("report generated at %s", r.Generated),
						},
					},
				},
			},
		}

		// TODO refactor url call
		if err := message.Post(ctx, s.URL); err != nil {
			return errors.Wrap(err, "failed to post slack message")
		}
	}
	return nil
}

func postSlackMessage(ctx context.Context, s WebhookSpec, msg slack.Message) error {

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	data, err := msg.Marshal()
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		s.URL,
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
		return fmt.Errorf("non-200 status returned from slack: %d", resp.StatusCode)
	}

	log.Println("All ok")
	return nil
}

func main() {
	lambda.Start(handler)
}
