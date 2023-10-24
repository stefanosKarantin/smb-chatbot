package messageplatform

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type MessageClient interface {
	SendPromotionMessage(msg PromotionMessage) error
}

type messageClient struct {
	client *http.Client
	host   string
}

func NewMessageClient(host string, client *http.Client) MessageClient {
	return messageClient{
		client: client,
		host:   host,
	}
}

func (c messageClient) SendPromotionMessage(msg PromotionMessage) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(body)
	resp, err := c.client.Post(c.host+"/start-promotion", "application/json", reader)
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
