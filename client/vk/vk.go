package vk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	methodConnect       = "messages.getLongPollServer"
	methodGet           = "messages.getLongPollHistory"
	methodSend          = "messages.send"
	vkBotHost           = "api.vk.com"
	groupId         int = 178512611
	vkAPIVersion        = "5.199"
	countQueryToLPS     = 1
)

type Client struct {
	host          string
	token         string
	ts            int
	countQueryLPS int
	client        http.Client
}

func New(token string) *Client {
	return &Client{
		host:          vkBotHost,
		token:         token,
		countQueryLPS: 20,
		client:        http.Client{},
	}
}

func (c *Client) Updates() ([]Message, error) {
	err := c.lpConnect()
	if err != nil {
		return nil, fmt.Errorf("[ERR] Cant connect LongPollyServer: %v", err)
	}

	res, err := c.lpRequest()
	if err != nil {
		return nil, fmt.Errorf("[ERR] Cant get MessageArray from LongPollyServer: %v", err)
	}

	return res, nil
}

func (c *Client) SendMessage(user_id int, message string) error {
	q := url.Values{}
	q.Add("user_id", strconv.Itoa(user_id))
	q.Add("random_id", "0")
	q.Add("message", message)
	q.Add("access_token", c.token)

	_, err := c.request(methodSend, q)
	if err != nil {
		return err
	}

	return nil
}

// Метод для подключения к Long Polly серверу
func (c *Client) lpConnect() error {
	if c.countQueryLPS < countQueryToLPS {
		c.countQueryLPS++
		return nil
	} else {
		c.countQueryLPS = 0
	}

	q := url.Values{}
	q.Add("access_token", c.token)
	q.Add("group_id", strconv.Itoa(groupId))

	data, err := c.request(methodConnect, q)
	if err != nil {
		return err
	}

	var res LongPollyConnect

	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	c.ts = res.Response.Ts
	return nil
}

// Метод получения сообщений из long Polly сервера
func (c *Client) lpRequest() ([]Message, error) {
	q := url.Values{}
	q.Add("access_token", c.token)
	q.Add("ts", strconv.Itoa(c.ts))
	q.Add("new_pts", "0")

	data, err := c.request(methodGet, q)
	if err != nil {
		return nil, err
	}

	var res LongPollyUpdate

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Response.MessageArray.Messages, nil
}

// Метод запросов к Vk API
func (c *Client) request(method string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join("method", method),
	}
	req, err := http.NewRequest(http.MethodPost, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("[ERR] Cant do request: %w", err)
	}

	query.Add("v", vkAPIVersion)

	req.URL.RawQuery = query.Encode()
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[ERR] Cant get request-response: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[ERR] Cant get response-body: %w", err)
	}
	return body, nil
}
