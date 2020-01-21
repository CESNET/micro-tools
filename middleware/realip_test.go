package middleware

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRealIP(t *testing.T) {
	Convey("Given some IP address", t, func() {
		ip := "1.2.3.4"
		req := httptest.NewRequest("GET", "http://dummy", nil)
		req.Header.Set("X-Forwarded-For", ip)
		So(req.RemoteAddr, ShouldNotEqual, ip)

		Convey("RealIP MW updates Remote Address with Real IP", func() {
			testRealIP := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				So(r.RemoteAddr, ShouldEqual, ip)
			})
			handlerToTest := RealIP(testRealIP)
			handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
		})
	})
}