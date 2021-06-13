package main

import (
	"net/http"

	"github.com/sanjeevsiva17/route69"
)

func main() {
	r := route69.NewRouter()

	r.GET("/name", func(r *http.Request) (statusCode int, data map[string]interface{}) {
		return 200, map[string]interface{}{
			"name": "route69",
		}
	})

	http.ListenAndServe(":8000", r)
}
