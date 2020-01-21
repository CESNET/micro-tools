package middleware

import (
	"crypto/rand"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func tokenGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func TestTokenAccess(t *testing.T) {
	Convey("Given an access token", t, func() {
		req := httptest.NewRequest("GET", "http://dummy", nil)
		testToken := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("handler", "called")
		})
		token := tokenGenerator()
		Convey("When handler is unprotected with token", func() {
			handlerToTest := Token("")(testToken)
			Convey("Token MW accepts the request", func() {
				resp := httptest.NewRecorder()
				handlerToTest.ServeHTTP(resp, req)
				So(resp.Header().Get("handler"), ShouldEqual, "called")
			})
		})

		Convey("When handler is protected with token", func() {
			handlerToTest := Token(token)(testToken)

			Convey("and correct access token is provided", func() {
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
				Convey("Token MW accepts the request", func() {
					resp := httptest.NewRecorder()
					handlerToTest.ServeHTTP(resp, req)
					So(resp.Header().Get("handler"), ShouldEqual, "called")
				})
			})
			Convey("and bad access token is provided", func() {
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", "just guessing..."))
				Convey("Token MW rejects the request with 401", func() {
					resp := httptest.NewRecorder()
					handlerToTest.ServeHTTP(resp, req)
					So(resp.Header().Get("handler"), ShouldBeBlank)
					So(resp.Code, ShouldEqual, 401)
				})
			})
		})
	})
}
