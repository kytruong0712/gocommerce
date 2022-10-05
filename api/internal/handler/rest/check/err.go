package check

import (
	"net/http"

	"github.com/kytruong0712/gocommerce/api/pkg/httpserver"
)

func (h handler) Err() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		return &httpserver.Error{
			Status: http.StatusBadRequest,
			Desc:   "Bad client request",
			Code:   "bad_request",
		}
	})
}
