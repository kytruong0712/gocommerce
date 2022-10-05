package check

import (
	"net/http"

	"github.com/kytruong0712/gocommerce/api/pkg/httpserver"
)

func (h handler) OK() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		httpserver.RespondJSON(w, httpserver.Success{Message: "OK"})

		return nil
	})
}
