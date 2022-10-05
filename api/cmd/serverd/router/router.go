package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/kytruong0712/gocommerce/api/internal/handler/rest/check"
)

// Router defines the routes & handlers of the app
type Router struct{}

func (rtr Router) routes(r chi.Router) {
	r.Group(rtr.public)
}

func (rtr Router) public(r chi.Router) {
	prefix := "/api/public"
	checkHandler := check.NewHandler()
	r.Group(func(r chi.Router) {
		r.Get(prefix+"/ok", checkHandler.OK())
		r.Get(prefix+"/err", checkHandler.Err())
	})
}
