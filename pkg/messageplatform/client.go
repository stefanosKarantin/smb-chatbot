package messageplatform

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type MessageClient interface {
	SendPromotionMessage(msg PromotionMessage) error
}

type messageClient struct {
	host   string
}

func NewMessageClient(host string) MessageClient {
	return messageClient{
		host:   host,
	}
}

// Mocking the http.Client Post request
func (c *messageClient) Post(url string, contentType string, body io.Reader) (*http.Response, error) {
	response := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(body),
	}
    return response, nil
}

func (c messageClient) SendPromotionMessage(msg PromotionMessage) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(body)
	resp, err := c.Post(c.host+"/send-message", "application/json", reader)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("text message was not sent")
	}
	return nil
}

type PromotionMessage struct {
	ID           int        `json:"id"`
	CustomerID   int        `json:"customer_id"`
	CustomerName string     `json:"customer_name"`
	Telephone    string     `json:"telephone"`
	Message      Message    `json:"message"`
	Responses    []Response `json:"responses"`
}

type Message struct {
	Text  string `json:"text"`
	Image string `json:"image"`
}

type Response struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	Url      string `json:"url"`
	NextStep int    `json:"next_step"`
}
