package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		log.Printf("The message %s for the source %s = %s\n", message.MessageId, message.EventSource, message.Body)
	}

	// ------

	svc := ses.New(session.New())

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String("john@johngregory.me.uk"),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String("Hallo welt!"),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String("Hallo welt! SUBJECT"),
			},
		},
		Source: aws.String("john.gregory@ons.gov.uk"),
	}

	_, err := svc.SendEmail(input)
	if err != nil {
		return err
	}

	// -------

	return nil
}

func main() {
	lambda.Start(handler)
}

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
// 				return errors.Wrap(aerr, ses.ErrCodeMessageRejected)
// 			case ses.ErrCodeMailFromDomainNotVerifiedException:
// 				return errors.Wrap(aerr, ses.ErrCodeMailFromDomainNotVerifiedException)
// 			case ses.ErrCodeConfigurationSetDoesNotExistException:
// 				return errors.Wrap(aerr, ses.ErrCodeConfigurationSetDoesNotExistException)
// 			default:
// 				return aerr
// 			}
// 		} else {
// 			// Print the error, cast err to awserr.Error to get the Code and
// 			// Message from an error.
// 			log.Println(err.Error())
// 		}

// 		return err
// 	}
// 	return nil
// }
