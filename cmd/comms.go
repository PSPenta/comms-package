package comms

import (
	"github.com/PSPenta/comms-package/config"
	"github.com/PSPenta/comms-package/internal"
	"github.com/PSPenta/comms-package/pkg"
)

func NewEmailClient(emailConfig *config.SenderInfo, recipientInfo *config.RecipientInfo) *internal.EmailClient {

	client := &internal.EmailClient{
		SenderInfo:    emailConfig,
		RecipientInfo: recipientInfo,
	}

	return client
}

func GetEmailCommunicator(client *internal.EmailClient) *pkg.EmailCommunicator {

	emailComms := &pkg.EmailCommunicator{
		Comms: client,
	}

	return emailComms
}

func GetSenderConfig() *config.SenderInfo {

	return &config.SenderInfo{}
}

func GetRecipientInfo() *config.RecipientInfo {
	return &config.RecipientInfo{}
}
