//
//  response.go
//  aych
//
//  Created by karim-w on 10/07/2025.
//

package aych

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Body Returns the response body as a byte slice.
func (c *client) Body() []byte { return c.respbody }

// Error Returns the error encountered during the request, if any.
func (c *client) Error() error { return c.err }

// Hearder Returns the response headers as an http.Header.
func (c *client) Header() http.Header { return c.resp_header }

// StatusCode Returns the HTTP status code of the response.
func (c *client) StatusCode() int { return c.resp_status_code }

// JSON Unmarshals the response body into the provided json.Unmarshaler.
func (c *client) JSON(v any) error { return json.Unmarshal(c.respbody, v) }

// Success Checks if the response was successful based on the status code and error.
func (c *client) Success() bool {
	return (c.resp_status_code >= 200 && c.resp_status_code < 300) && c.err == nil
}

// CURL Generates a cURL command representation of the request.
func (c *client) CURL() string {
	builder := strings.Builder{}
	builder.WriteString("curl -X ")
	builder.WriteString(c.method)
	builder.WriteString(" '")
	builder.WriteString(c.url)
	builder.WriteString("'")
	for k, v := range c.headers {
		builder.WriteString(" -H '")
		builder.WriteString(k)
		builder.WriteString(": ")
		builder.WriteString(v[0])
		builder.WriteString("'")
	}
	if c.reqbody != nil {
		builder.WriteString(" -d '")
		b := string(c.reqbody)
		builder.WriteString(b)
		builder.WriteString("'")
	}
	return builder.String()
}

// Close Cleans up the client by resetting its fields.
func (c *client) Close() error {
	c.respbody = nil
	c.resp_header = nil
	c.cookies = nil
	c.query = nil
	c.headers = nil
	c.reqbody = nil
	c.url = ""
	c.method = ""
	c.idx = 0
	c.middlewares = nil

	return nil
}
