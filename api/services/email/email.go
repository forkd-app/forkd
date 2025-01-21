package email

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"forkd/util"
	"io"
	"net/http"
)

type EmailService interface {
	SendMagicLink(ctx context.Context, token string, email string) (*SendEmailResponseBody, error)
}

type emailService struct {
	apiKey  string
	baseUrl string
}

type sendEmailReqBody struct {
	Sender        string            `json:"sender"`
	To            []string          `json:"to,omitempty"`
	Cc            []string          `json:"cc,omitempty"`
	Bcc           []string          `json:"bcc,omitempty"`
	Subject       string            `json:"subject,omitempty"`
	HtmlBody      string            `json:"html_body,omitempty"`
	TextBody      string            `json:"text_body,omitempty"`
	CustomHeaders []emailHeader     `json:"custom_headers,omitempty"`
	Attachments   []emailAttachment `json:"attachments,omitempty"`
	Inlines       []emailAttachment `json:"inlines,omitempty"`
	TemplateId    string            `json:"template_id,omitempty"`
	TemplateData  any               `json:"template_data,omitempty"`
}

type emailHeader struct {
	Header string `json:"header,omitempty"`
	Value  string `json:"value,omitempty"`
}

type emailAttachment struct {
	Filename string `json:"filename,omitempty"`
	FileBlob string `json:"fileblob,omitempty"`
	MimeType string `json:"mimetype,omitempty"`
	Url      string `json:"url,omitempty"`
}

type SendEmailResponseBody struct {
	RequestId string                    `json:"request_id,omitempty"`
	Data      SendEmailResponseBodyData `json:"data,omitempty"`
}

type SendEmailResponseBodyData struct {
	Failed    int8          `json:"failed,omitempty"`
	Failures  []interface{} `json:"-"`
	Succeeded int8          `json:"succeeded,omitempty"`
	EmailID   string        `json:"email_id,omitempty"`
}

// TODO: This really needs some love lol
func (e emailService) SendMagicLink(ctx context.Context, token string, email string) (*SendEmailResponseBody, error) {
	// Grab the base url to use for the magic link from the environment
	linkBaseUrl := util.GetEnv().GetBaseUrl()
	// Setup some info for the email like the recipient, subject, and body (We do both plain text and html so we have a fallback in case the email client doesn't support html)
	// TODO: Put this in an env var
	from := "Forkd <auth@forkd.gvasquez.dev>"
	subject := "Forkd Login"
	plainTextBody := fmt.Sprintf("Click here or copy and paste to login: %s/auth/validate?token=%s", linkBaseUrl, token)
	htmlBody := `
    <html>
      <body>
        <p>Click or copy and paste into your browser to login: <a href="%[1]s/auth/validate?token=%[2]s">%[1]s/auth/validate?token=%[2]s</a></p>
      </body>
    </html>
  `
	requestBody := sendEmailReqBody{
		Sender:   from,
		To:       []string{email},
		Subject:  subject,
		TextBody: plainTextBody,
		HtmlBody: fmt.Sprintf(htmlBody, linkBaseUrl, token),
	}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling requestBody: %w", err)
		return nil, err
	}
	fmt.Printf("%s/email/send\n", e.baseUrl)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/email/send", e.baseUrl), bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		fmt.Println("Error building request: %w", err)
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Smtp2go-Api-Key", e.apiKey)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request: %w", err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body: %w", err)
		return nil, err
	}
	fmt.Printf("Response Body: %s\n", string(body))
	var responseBody SendEmailResponseBody
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		fmt.Println("Error unmarshaling responseBody: %w", err)
		return nil, err
	}
	return &responseBody, nil
}

func New() EmailService {
	env := util.GetEnv()
	return emailService{
		apiKey:  env.GetEmailApiKey(),
		baseUrl: env.GetEmailBaseUrl(),
	}
}
