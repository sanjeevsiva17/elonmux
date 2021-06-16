# elonmux
Custom golang multiplexer(router)

![alt text](https://miro.medium.com/max/239/1*EvsjfKJSWsCvoOHMfSKvBA.jpeg)

credits : reddit

## Install 
`go get -u github.com/sanjeevsiva17/elonmux`

## Examples 

```
package main

import (
	"net/http"

	"github.com/sanjeevsiva17/elonmux"
)

func main() {
	r := elonmux.NewRouter()

	r.GET("/name", func(r *http.Request) (statusCode int, data map[string]interface{}) {
		return 200, map[string]interface{}{
			"name": "elonmux",
		}
	})

	http.ListenAndServe(":8000", r)
}
```
