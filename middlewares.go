//
//  middlewares.go
//  aych
//
//  Created by karim-w on 10/07/2025.
//

package aych

// Use adds middleware functions to the client's middleware chain.
// Middleware functions are executed in the order they are added.
// Each middleware function receives the TTPContext and can modify it or call the next middleware.
func (c *client) Use(middleware ...func(tx TTPContext)) Request {
	c.middlewares = append(c.middlewares, middleware...)
	return c
}

// Next executes the next middleware in the chain.
func (c *client) Next() {
	if int(c.idx) < len(c.middlewares) {
		fn := c.middlewares[c.idx]
		c.idx++
		fn(c)
	}
}

// run executes the middleware chain and performs the HTTP request.
func (c *client) run() Response {
	c.idx = 0

	// Define a final function that performs the actual HTTP request
	c.middlewares = append(c.middlewares, func(tx TTPContext) {
		c.doRequest()
	})

	// Kick off the middleware chain
	c.Next()

	return c
}
