package middleware

import (
	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestID(t *testing.T) {
	Convey("Given some Request", t, func() {
		req := httptest.NewRequest("GET", "http://dummy", nil)

		Convey("When Request ID is not in headers", func() {
			So(req.Header.Get("X-Request-ID"), ShouldBeBlank)

			Convey("RequestID MW creates and sets RequestID header", func() {
				testRID := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					So(r.Header.Get("X-Request-ID"), ShouldNotBeBlank)
					_, err := uuid.Parse(r.Header.Get("X-Request-ID"))
					So(err, ShouldBeNil)
				})
				handlerToTest := RequestID(testRID)
				handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
			})
		})
	})

}
