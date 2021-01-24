package decoHandler

import "net/http"

type DecoratorFunc func(http.ResponseWriter, *http.Request, http.Handler)

type DecoHandler struct {
	fn_DecoratorFunc
	h http.Handler
}

func (self * ServeHTTP(w http.Handler, r *http.Request) {
	self.fn(w, r, self.h)
}

func NewDecoHandler(h_http.Handler, fn DecoratorFunc) http.Handler {
	return &DecoHandler {
		fn: fn,
		h: h
	}

}
