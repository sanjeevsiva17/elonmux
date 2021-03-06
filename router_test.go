package elonmux

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func testHandlerGET(r *http.Request) (int, map[string]interface{}) {
	return http.StatusOK, map[string]interface{}{"data": "test"}

}

func testHandlerPOST(r *http.Request) (int, map[string]interface{}) {
	return http.StatusCreated, map[string]interface{}{"data": "test"}

}

func testHandlerPUT(r *http.Request) (int, map[string]interface{}) {
	return http.StatusOK, map[string]interface{}{"data": "test"}

}

func testHandlerDELETE(r *http.Request) (int, map[string]interface{}) {
	return http.StatusNoContent, map[string]interface{}{"data": "test"}

}

func TestRouter(t *testing.T) {
	r := NewRouter()
	r.GET("/test$", testHandlerGET)
	r.POST("/test$", testHandlerPOST)
	r.PUT("/test$", testHandlerPUT)
	r.DELETE("/test$", testHandlerDELETE)

	tc := []struct {
		method string
		path   string
		code   int
	}{
		{"GET", "/test", http.StatusOK},
		{"POST", "/test", http.StatusCreated},
		{"PUT", "/test", http.StatusOK},
		{"DELETE", "/test", http.StatusNoContent},
		{"DELETE", "/test1", http.StatusNotFound},
	}

	for _, c := range tc {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(c.method, c.path, nil)

		r.ServeHTTP(rec, req)

		// Check if HTTP and JSON Response code matches
		if status := rec.Code; status != c.code {
			t.Errorf("Wrong status code for %v: got %v want %v. Response: %v",
				c, status, c.code, rec.Body.String())
		} else {
			t.Logf("Case passed for %v", c)
		}
	}
}
