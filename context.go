//
//  context.go
//  aych
//
//  Created by karim-w on 10/07/2025.
//

package aych

import "net/url"

// Query returns the query parameters for the request.
func (c *client) Query() url.Values {
	return c.query
}

// Url returns the URL for the request.
func (c *client) Url() string {
	return c.url
}
