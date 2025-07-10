//
//  transformers.go
//  aych
//
//  Created by karim-w on 10/07/2025.
//

package aych

import "net/http"

func (c *client) Del() Response {
	c.method = http.MethodDelete
	return c.run()
}

func (c *client) Get() Response {
	c.method = http.MethodGet
	return c.run()
}

func (c *client) Patch() Response {
	c.method = http.MethodPatch
	return c.run()
}

func (c *client) Post() Response {
	c.method = http.MethodPost
	return c.run()
}

func (c *client) Put() Response {
	c.method = http.MethodPut
	return c.run()
}
