package main

import (
	"context"
	"os"

	"github.com/ONSdigital/who-goes-there/pkg/github"
	"github.com/ONSdigital/who-goes-there/pkg/report"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
)

// Handler recieves a CloudWatch Event and triggers the rules engine
func Handler(ctx context.Context, cwEvent events.CloudWatchEvent) (*report.Report, error) {

	queueURL := os.Getenv("QUEUE_URL")
	if queueURL == "" {
		return nil, errors.New("missing QUEUE_URL")
	}

	token := os.Getenv("API_TOKEN")
	if token == "" {
		return nil, errors.New("missing API_TOKEN")
	}

	organisation := os.Getenv("GITHUB_ORG_NAME")
	if organisation == "" {
		return nil, errors.New("missing GITHUB_ORG_NAME")
	}

	client := github.NewClient(token)
	users, err := client.FetchOrganizationMembers(organisation)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve members")
	}

	if os.Getenv("REPORT_RECIPIENT") == "" || os.Getenv("REPORT_SENDER") == "" {
		return nil, errors.New("Missing sender and recipient")
	}

	report := report.Report{}

	for _, user := range users {
		report.Summary.TotalUsers++
		if !user.HasTwoFactorEnabled {
			report.Summary.UsersMissingMFA++
		}
		if user.Role == "ADMIN" {
			report.Summary.AdminUsers++
			if !user.HasTwoFactorEnabled {
				report.Summary.AdminUsersFailingRules++
			}
		}
	}

	// sess := session.New()
	// svc := sqs.New(sess)

	// rj, err := json.Marshal(report)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to marshal reports")
	// }

	// result, err := svc.SendMessage(&sqs.SendMessageInput{
	// 	MessageAttributes: map[string]*sqs.MessageAttributeValue{
	// 		"Title": &sqs.MessageAttributeValue{
	// 			DataType:    aws.String("String"),
	// 			StringValue: aws.String("W"),
	// 		},
	// 	},
	// 	MessageBody: aws.String(string(rj)),
	// 	QueueUrl:    &queueURL,
	// })

	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to send report message")
	// }
	// log.Println("Successfully published message:", *result.MessageId)

	return &report, nil
}

// func sendReportMessage(r *report.Report) error {

// 	rj, err := json.Marshal(r)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to marshal reports")
// 	}

// 	log.Println(rj)

// 	return nil
// }

// func sendReportEmail(rs *ReportSummary) error {
// 	sess, err := session.NewSession(&aws.Config{
// 		Region: aws.String("us-west-2")},
// 	)

// 	// Create an SES session.
// 	svc := ses.New(sess)

// 	// Assemble the email.
// 	input := &ses.SendEmailInput{
// 		Destination: &ses.Destination{
// 			CcAddresses: []*string{},
// 			ToAddresses: []*string{
// 				aws.String(os.Getenv("REPORT_RECIPIENT")),
// 			},
// 		},
// 		Message: &ses.Message{
// 			Body: &ses.Body{
// 				Text: &ses.Content{
// 					Charset: aws.String("UTF-8"),
// 					Data:    aws.String(rs.ToString()),
// 				},
// 			},
// 			Subject: &ses.Content{
// 				Charset: aws.String("UTF-8"),
// 				Data:    aws.String("Github Who Goes There Report Summary"),
// 			},
// 		},
// 		Source: aws.String(os.Getenv("REPORT_SENDER")),
// 	}

// 	// Attempt to send the email.
// 	_, err = svc.SendEmail(input)

// 	// Display error messages if they occur.
// 	if err != nil {
// 		if aerr, ok := err.(awserr.Error); ok {
// 			switch aerr.Code() {
// 			case ses.ErrCodeMessageRejected:
// 				return nil, errors.Wrap(aerr, ses.ErrCodeMessageRejected)
// 			case ses.ErrCodeMailFromDomainNotVerifiedException:
// 				return nil, errors.Wrap(aerr, ses.ErrCodeMailFromDomainNotVerifiedException)
// 			case ses.ErrCodeConfigurationSetDoesNotExistException:
// 				return nil, errors.Wrap(aerr, ses.ErrCodeConfigurationSetDoesNotExistException)
// 			default:
// 				return aerr
// 			}
// 		} else {
// 			// Print the error, cast err to awserr.Error to get the Code and
// 			// Message from an error.
// 			log.Println(err.Error())
// 		}

// 		return nil, err
// 	}
// 	return nil
// }

func main() {
	lambda.Start(Handler)
}
