package comms

func main() {

	senderInfo := GetSenderConfig()
	senderInfo.FromEmail = ""
	senderInfo.ServiceType = "aws|api"
	senderInfo.AwsRegion = ""
	senderInfo.AwsAccessKey = ""
	senderInfo.AwsSecretAccessKey = ""

	recipientInfo := GetRecipientInfo()
	recipientInfo.SendTo = []string{""}
	recipientInfo.CC = []string{""}
	recipientInfo.ContentType = "plain/text"

	client := NewEmailClient(senderInfo, recipientInfo)

	emailCommunicator := GetEmailCommunicator(client)

	emailCommunicator.Comms.SendMail()
}
