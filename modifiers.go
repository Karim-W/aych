//
//  modifiers.go
//  aych
//
//  Created by karim-w on 10/07/2025.
//

package aych

import (
	"encoding/json"
	"net/http"

	"github.com/karim-w/aych/httpheaders"
)

// Adds a Basic Authentication header to the request.
// Does not check if the username and password are empty.
// if you want to check for empty values, do it before calling this method like
// a good boy
// example usage:
//
// client.AddBasicAuth("your_username", "your_password")
func (c *client) AddBasicAuth(username string, password string) Request {
	c.headers.Add(httpheaders.AUTHORIZATION, string_builder("Basic ", base64_encode(username+":"+password)))
	return c
}

// Adds a Bearer Authentication header to the request.
// example usage:
//
//	client.AddBearerAuth("your_token_here")
func (c *client) AddBearerAuth(token string) Request {
	c.headers.Add(httpheaders.AUTHORIZATION, string_builder("Bearer ", token))
	return c
}

// Adds a body to the request.
// The body can be any type that can be marshaled to JSON.
// If the body cannot be marshaled to JSON, the error will be set in the client with
// Status Code
func (c *client) JSONBody(body any) Request {
	data, err := json.Marshal(body)
	if err != nil {
		c.err = err
	}

	c.reqbody = data

	return c
}

// Adds a raw body to the request.
func (c *client) AddBodyRaw(body []byte) Request {
	c.reqbody = body
	return c
}

// AddCookie adds a single cookie to the request.
func (c *client) AddCookie(cookie *http.Cookie) Request {
	c.cookies = append(c.cookies, cookie)
	return c
}

func (c *client) AddHeader(key string, value string) Request {
	c.headers.Add(key, value)
	return c
}

// AddHeaders adds multiple headers to the request.
func (c *client) AddHeaders(headers map[string]string) Request {
	for key, value := range headers {
		c.headers.Add(key, value)
	}
	return c
}

// AddQuery adds a single query parameter to the request.
// If the key already exists, the value will be appended to the existing values.
func (c *client) AddQuery(key string, value string) Request {
	values, exists := c.query[key]
	if !exists {
		values = []string{}
	}

	values = append(values, value)
	c.query[key] = values

	return c
}

// AddQueryArray adds a slice of strings as a query parameter to the request.
func (c *client) AddQueryArray(key string, value []string) Request {
	if len(value) == 0 {
		return c
	}

	values, exists := c.query[key]
	if !exists {
		values = []string{}
	}
	values = append(values, value...)

	return c
}
