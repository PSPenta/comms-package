package internal

import (
	"bytes"
	"comms-package/config"
	"comms-package/pkg"
	"encoding/json"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/fatih/structs"
)

type EmailClient struct {
	SenderInfo    *config.SenderInfo
	RecipientInfo *config.RecipientInfo
	Comms         pkg.EmailCommunicator
}

/**
 * This will send mail based on email service type
 */
func (e *EmailClient) SendMail() error {

	emailService := strings.ToLower(e.SenderInfo.ServiceType)

	if emailService == "smtp" {
		return e.smtpService()

	} else if emailService == "api" {
		return e.apisService()

	} else if emailService == "aws" {
		return e.awsSesService()

	}

	return nil
}

/**
 * This will parse template
 * e.recipientInfo.TemplateName with the interface
 * which is passed as parameter in the function call
 */
func (e *EmailClient) ParseTemplate(data interface{}) (*EmailClient, error) {

	templateName := e.RecipientInfo.TemplateName

	if len(strings.TrimSpace(templateName)) > 0 && templateName != " " {

		t, err := template.ParseFiles(e.RecipientInfo.TemplateName)
		if err != nil {
			return e, err
		}
		buf := new(bytes.Buffer)
		if err = t.Execute(buf, data); err != nil {
			return e, err
		}

		e.RecipientInfo.Body = buf.String()

	}

	return e, nil
}

/**
 * This will convert message into bytes
 */
func (e *EmailClient) message() []byte {
	var msg []byte

	if e.RecipientInfo.TemplateName != "" && len(e.RecipientInfo.TemplateName) > 0 {
		mime := "MIME-version: 1.0;\nContent-Type:" + e.RecipientInfo.ContentType + "; charset=\"UTF-8\";\n\n"

		msg = []byte(mime + "\n" + e.RecipientInfo.Body)
	} else {
		msg = []byte(e.RecipientInfo.Body)
	}

	return msg
}

func (e *EmailClient) smtpService() error {

	SMTP := fmt.Sprintf("%s:%d", e.SenderInfo.Address, e.SenderInfo.Port)

	auth := smtp.PlainAuth("", e.SenderInfo.UserName, e.SenderInfo.Password, e.SenderInfo.Address)

	smtp.SendMail(SMTP, auth, e.SenderInfo.FromEmail, e.RecipientInfo.SendTo, e.message())

	return nil
}

func (e *EmailClient) awsSesService() error {

	session, err := session.NewSession(
		&aws.Config{
			Region:      aws.String(e.SenderInfo.AwsRegion),
			Credentials: credentials.NewStaticCredentials(e.SenderInfo.AwsAccessKey, e.SenderInfo.AwsSecretAccessKey, e.SenderInfo.AwsToken),
		},
	)
	if err != nil {
		fmt.Println("Error occurred while creating aws session", err)
		return nil
	}

	// set to section
	var recipients []*string
	for _, r := range e.RecipientInfo.SendTo {
		recipient := r
		recipients = append(recipients, &recipient)
	}

	// set cc section
	var ccRecipients []*string
	if len(e.RecipientInfo.CC) > 0 {
		for _, r := range e.RecipientInfo.CC {
			ccRecipient := r
			ccRecipients = append(ccRecipients, &ccRecipient)
		}
	}

	// set bcc section
	var bccRecipients []*string
	if len(e.RecipientInfo.BCC) > 0 {
		for _, r := range e.RecipientInfo.BCC {
			bccRecipient := r
			bccRecipients = append(bccRecipients, &bccRecipient)
		}
	}

	// create an SES session.
	sesSession := ses.New(session)

	// Assemble the email.
	input := &ses.SendEmailInput{

		// Set destination emails
		Destination: &ses.Destination{
			ToAddresses:  recipients,
			CcAddresses:  ccRecipients,
			BccAddresses: bccRecipients,
		},

		// Set email message and subject
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(e.RecipientInfo.Body),
				},
			},

			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(e.RecipientInfo.EmailSubject),
			},
		},

		// send from email
		Source: aws.String(e.SenderInfo.FromEmail),
	}

	// Call AWS send email function which internally calls to SES API
	_, err = sesSession.SendEmail(input)
	if err != nil {
		fmt.Println("Error sending mail - ", err)
	}

	return nil
}

func (e *EmailClient) apisService() error {

	request, err := pkg.CurlRequest(e.SenderInfo)

	if err != nil {
		return err
	}

	s := structs.New(e.SenderInfo.Header)

	request.SetBasicAuth(e.SenderInfo.UserName, e.SenderInfo.Password)

	for _, v := range s.Fields() {

		fieldName := v.Name()

		if v.IsExported() {
			request.Header.Add(fieldName, v.Value().(string))
		}
	}

	curlResponse := pkg.CurlResponse(request)

	result := &config.Response{}

	json.Unmarshal(curlResponse, result)

	if result.Status != 200 {

		err := fmt.Errorf("failed to send email: %v", result)

		return err
	}

	return nil
}
