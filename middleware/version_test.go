package middleware

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVersion(t *testing.T) {
	Convey("Given a version string", t, func() {
		name := "app"
		version := "0.1.1"

		Convey("Version MW sets X-%name-VERSION response header", func() {
			testVersion := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				So(w.Header().Get(fmt.Sprintf("X-%s-VERSION", name)), ShouldEqual, version)
			})
			handlerToTest := Version(name, version)(testVersion)
			req := httptest.NewRequest("GET", "http://dummy", nil)
			handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
		})
	})
}