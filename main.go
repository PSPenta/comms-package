package main

import comms "github.com/PSPenta/comms-package/cmd"

func main() {

	senderInfo := comms.GetSenderConfig()
	senderInfo.FromEmail = ""
	senderInfo.ServiceType = "aws|api"
	senderInfo.AwsRegion = ""
	senderInfo.AwsAccessKey = ""
	senderInfo.AwsSecretAccessKey = ""

	recipientInfo := comms.GetRecipientInfo()
	recipientInfo.SendTo = []string{""}
	recipientInfo.CC = []string{""}
	recipientInfo.ContentType = "plain/text"

	client := comms.NewEmailClient(senderInfo, recipientInfo)

	emailCommunicator := comms.GetEmailCommunicator(client)

	emailCommunicator.Comms.SendMail()
}
