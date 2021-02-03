package express

// Middleware q
type Middleware func(req *Request, res *Response, next Next)

// Next q
type Next func()
