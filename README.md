<p align="center"><img src="https://managesandbox.purplle.com/static/admin/layout/img/finder-logo.png"></p>

# Communication Package

```
This is a Communication library for Go Language.
Which sends communications either using SES service, or using CRM comms APIs.
```
## Installation

Installation is done using `go get`.
```
go get -u gitlab.com/purplle-experience/comms-package
```

#### Step to follow :


- import `comms "gitlab.com/purplle-experience/comms-package/cmd"`
- Initialize the SenderConfig and RecipientInfo, 
```
import comms "gitlab.com/purplle-experience/comms-package/cmd"

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
}
```
-  Initialize the email client once the SenderConfig and RecipientInfo is configured,
```
import comms "comms-package/cmd"

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
}
```
-  Then initialize the email communicator using the client initialized previously,
```
import comms "comms-package/cmd"

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
}
```
-  Now the emails can be sent using the SendMail method as shown below,
```
import comms "comms-package/cmd"

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
```
