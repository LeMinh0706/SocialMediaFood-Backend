package mails

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}
