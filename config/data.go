package config

type RecipientInfo struct {
	SendTo       []string `json:"send_to"`
	CC           []string `json:"cc"`
	BCC          []string `json:"bcc"`
	Body         string   `json:"body"`
	ContentType  string   `json:"content_type"`
	TemplateName string   `json:"template_name"`
	EmailSubject string   `json:"email_subject"`
}

type SenderInfo struct {
	Address            string `json:"address"`
	Port               int    `json:"port"`
	UserName           string `json:"user_name"`
	Password           string `json:"password"`
	FromEmail          string `json:"from_email"`
	ServiceType        string `json:"service_type"`
	Method             string `json:"method"`
	URL                string `json:"url"`
	Payload            interface{}
	Header             interface{}
	AwsRegion          string `json:"region"`
	AwsAccessKey       string `json:"aws_access_key"`
	AwsSecretAccessKey string `json:"aws_secret_access_key"`
	AwsToken           string `json:"aws_token"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
