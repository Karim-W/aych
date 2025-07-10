//
//  contracts.go
//  aych
//
//  Created by karim-w on 10/07/2025.
//

package aych

import (
	"context"
	"net/http"
	"net/url"
)

type _TTPRequestModifiers interface {
	AddHeader(key string, value string) Request
	AddHeaders(headers map[string]string) Request
	AddQuery(key string, value string) Request
	AddQueryArray(key string, value []string) Request
	JSONBody(body any) Request
	AddBodyRaw(body []byte) Request
	AddBasicAuth(username string, password string) Request
	AddBearerAuth(token string) Request
	AddCookie(cookie *http.Cookie) Request
}

type _TTPResponse interface {
	StatusCode() int
	Body() []byte
	Header() http.Header
	Error() error
}

type TTPContext interface {
	_TTPRequestModifiers
	_TTPResponse
	Url() string
	Query() url.Values

	Next()
}

type Request interface {
	_TTPRequestModifiers
	Use(middleware ...func(tx TTPContext)) Request
	Get() Response
	Put() Response
	Del() Response
	Post() Response
	Patch() Response
}

type Response interface {
	_TTPResponse
	JSON(v any) error
	Error() error
	Success() bool
	CURL() string
	Close() error
}

type client struct {
	url                   string
	method                string
	headers               http.Header
	response_content_leng int
	idx                   uint
	middlewares           []func(tx TTPContext)
	reqbody               []byte
	respbody              []byte
	resp_header           http.Header
	resp_status_code      int
	err                   error
	ctx                   context.Context
	cookies               []*http.Cookie
	query                 url.Values
}

func TTP(ctx context.Context, req_url string) Request {
	return &client{
		url:                   req_url,
		method:                "",
		headers:               http.Header{},
		response_content_leng: 0,
		idx:                   0,
		middlewares:           []func(tx TTPContext){},
		reqbody:               []byte{},
		respbody:              []byte{},
		resp_header:           http.Header{},
		resp_status_code:      0,
		err:                   nil,
		ctx:                   ctx,
		cookies:               []*http.Cookie{},
		query:                 url.Values{},
	}
}

// ASSERTIONS
var (
	_ _TTPRequestModifiers = (*client)(nil)
	_ _TTPResponse         = (*client)(nil)
	_ TTPContext           = (*client)(nil)
	_ Request              = (*client)(nil)
	_ Response             = (*client)(nil)
)
