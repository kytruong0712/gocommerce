package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kytruong0712/gocommerce/api/cmd/banner"
	"github.com/kytruong0712/gocommerce/api/internal/handler/rest/check"
)

func main() {
	banner.Print()

	checkHandler := check.NewHandler()
	prefix := "/api/public"

	http.HandleFunc(prefix+"/ok", checkHandler.OK())
	http.HandleFunc(prefix+"/err", checkHandler.Err())

	fmt.Println("Starting server at port 3001")
	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatal(err)
	}
}
