package shiori

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) get(path string, vres interface{}) error {
	return c.request("GET", path, nil, vres)
}

func (c *Client) post(path string, body, vres interface{}) error {
	return c.request("POST", path, body, vres)
}

func (c *Client) put(path string, body, vres interface{}) error {
	return c.request("PUT", path, body, vres)
}

func (c *Client) delete(path string, vres interface{}) error {
	return c.request("DELETE", path, nil, vres)
}

func (c *Client) request(method, path string, body, vres interface{}) (err error) {
	if vres == nil {
		vres = &map[string]interface{}{}
	}

	bodyReader := bytes.NewBuffer([]byte{})
	if body != nil {
		if err = json.NewEncoder(bodyReader).Encode(body); err != nil {
			return
		}
	}

	url := fmt.Sprintf("%s/%s", c.endpoint, path)
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("basic %s", c.token))

	res, err := c.client.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode >= 400 {
		var msg []byte
		msg, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return
		}
		err = fmt.Errorf("%d %s", res.StatusCode, string(msg))
		return
	}

	return json.NewDecoder(res.Body).Decode(vres)
}
