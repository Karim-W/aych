//
//  runner.go
//  aych
//
//  Created by karim-w on 10/07/2025.
//

package aych

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

func (c *client) doRequest() {
	if c.err != nil {
		return
	}

	var bd io.Reader
	if c.reqbody != nil {
		bd = bytes.NewReader(c.reqbody)
	}

	url := c.url
	if len(c.query) > 0 {
		url += "?" + c.query.Encode()
	}

	req, err := http.NewRequest(c.method, url, bd)
	if err != nil {
		c.err = err
		return
	}

	req.Header = c.headers

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		c.err = err
		return
	}
	defer resp.Body.Close()

	if resp == nil {
		c.err = errors.New("received nil response")
		return
	}

	c.resp_status_code = resp.StatusCode
	c.resp_header = resp.Header
	c.respbody, c.err = io.ReadAll(resp.Body)
	c.response_content_leng = int(resp.ContentLength)
}
