package express

import (
	"encoding/json"
	"net/http"
	"time"
)

// Response q
type Response struct {
	http.ResponseWriter
	req      *http.Request
	App      *App
	status   int
	messages []Message
}

// AddMessage func
func (res *Response) AddMessage(typ, msg string) {
	res.messages = append(res.messages, Message{typ, msg})
}

// GetMessages func
func (res *Response) GetMessages() []Message {
	return res.messages
}

// Prepare func
func (res *Response) Prepare() {
	res.WriteHeader(res.status)
}

// Message type
type Message struct {
	Typ     string
	Message string
}

func newResponse(r *http.Request, w http.ResponseWriter, app *App) *Response {
	return &Response{
		req:            r,
		ResponseWriter: w,
		App:            app,
		status:         http.StatusOK,
	}
}

type responseFunc func(interface{}) error

func (res *Response) respond(fn responseFunc, data interface{}) error {
	res.WriteHeader(res.status)

	return fn(data)
}

// JSON write json repsonse
func (res *Response) JSON(data interface{}) error {
	res.Header().Add("Content-Type", "application/json")
	fn := json.NewEncoder(res).Encode

	return res.respond(fn, data)
}

// Send write string rensponse
func (res *Response) Send(data string) error {
	res.Header().Add("Content-Type", "text/plain")

	fn := func(v interface{}) error {
		bytes := v.([]byte)
		_, err := res.Write(bytes)
		return err
	}

	return res.respond(fn, data)
}

// Status func
func (res *Response) Status(code int) *Response {
	res.status = code
	return res
}

// Error func
func (res *Response) Error(err string, code int) {
	w := res

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	res.Status(code).JSON(err)
}

// SetCookie func
func (res *Response) SetCookie(c *http.Cookie) {
	http.SetCookie(res.ResponseWriter, c)
}

// ClearCookie func
func (res *Response) ClearCookie(name string) {
	c := &http.Cookie{
		Name:    name,
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),

		HttpOnly: true,
	}

	http.SetCookie(res.ResponseWriter, c)

}

// Redirect func
func (res *Response) Redirect(url string) {
	status := res.status
	status = http.StatusFound
	http.Redirect(res.ResponseWriter, res.req, url, status)
}
