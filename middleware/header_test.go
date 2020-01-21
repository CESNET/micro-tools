package middleware

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHeader(t *testing.T) {
	Convey("Given some HTTP handler", t, func() {
		Convey("Cache MW sets proper cache response headers", func() {
			testCache := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				cc := w.Header().Get("Cache-Control")
				So(cc, ShouldEqual, "no-cache, no-store, max-age=0, must-revalidate, value")
				exp := w.Header().Get("Expires")
				So(exp, ShouldEqual, "Thu, 01 Jan 1970 00:00:00 GMT")
				mod := w.Header().Get("Last-Modified")
				So(mod, ShouldEqual, time.Now().UTC().Format(http.TimeFormat))
			})
			handlerToTest := Cache(testCache)
			req := httptest.NewRequest("OPTIONS", "http://dummy", nil)
			handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
		})
		Convey("CORS MW sets proper cors response headers", func() {
			testCORS := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "OPTIONS" {
					So(w.Header().Get("Access-Control-Allow-Origin"), ShouldBeEmpty)
					So(w.Header().Get("Access-Control-Allow-Methods"), ShouldBeEmpty)
					So(w.Header().Get("Access-Control-Allow-Headers"), ShouldBeEmpty)
					So(w.Header().Get("Allow"), ShouldBeEmpty)
				}
			})
			handlerToTest := Cors(testCORS)

			Convey("When OPTIONS request is given", func() {
				req := httptest.NewRequest("OPTIONS", "http://dummy", nil)
				resp := httptest.NewRecorder()
				handlerToTest.ServeHTTP(resp, req)
				So(resp.Code, ShouldEqual, 200)
				So(resp.Header().Get("Access-Control-Allow-Origin"), ShouldEqual, "*")
				So(resp.Header().Get("Access-Control-Allow-Methods"), ShouldEqual,
					"GET, POST, PUT, PATCH, DELETE, OPTIONS")
				So(resp.Header().Get("Access-Control-Allow-Headers"), ShouldEqual,
					"authorization, origin, content-type, accept")
				So(resp.Header().Get("Allow"), ShouldEqual,
					"HEAD, GET, POST, PUT, PATCH, DELETE, OPTIONS")
			})

			Convey("When GET request is given", func() {
				req := httptest.NewRequest("GET", "http://dummy", nil)
				handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
			})
		})
		Convey("Secure MW sets proper security response headers", func() {
			testSecure := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				So(w.Header().Get("Access-Control-Allow-Origin"), ShouldEqual,"*")
				So(w.Header().Get("X-Frame-Options"), ShouldEqual,"DENY")
				So(w.Header().Get("X-Frame-Options"), ShouldEqual,"DENY")
				So(w.Header().Get("X-Content-Type-Options"), ShouldEqual,"nosniff")
				So(w.Header().Get("X-XSS-Protection"), ShouldEqual,"1; mode=block")
				if r.TLS != nil {
					So(w.Header().Get("Strict-Transport-Security"), ShouldStartWith,"max-age=")
				}
			})
			handlerToTest := Secure(testSecure)

			Convey("When HTTP request is given", func() {
				req := httptest.NewRequest("GET", "http://dummy", nil)
				handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
			})
			Convey("When HTTPS request is given", func() {
				req := httptest.NewRequest("GET", "https://dummy", nil)
				handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
			})
		})
	})
}
