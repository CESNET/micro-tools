package debug

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

func TestDebugService(t *testing.T) {
	Convey("Given some Debug Service Options", t, func() {
		token := tokenGenerator()
		ropts := &[]Option{
			Name("testing"),
			Token(token),
			Version("0.0.0"),
			Address("0.0.0.0:4321"),
		}
		opts := newOptions(*ropts...)
		Convey("Options are initialized properly", func() {
			So(opts.Version, ShouldEqual, "0.0.0")
			So(opts.Address, ShouldEqual, "0.0.0.0:4321")
			So(opts.Name, ShouldEqual, "testing")
			So(opts.Token, ShouldEqual, token)
		})

		Convey("With all handlers configured", func() {
			// TODO: test more debug handlers...
			healthz := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("health", "testing")
				w.WriteHeader(200)
			})
			readyz := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("ready", "testing")
				w.WriteHeader(200)
			})
			*ropts = append(*ropts, Health(healthz))
			*ropts = append(*ropts, Ready(readyz))

			Convey("and properly initialized service", func() {
				svc := NewService(*ropts...)
				So(svc.Addr, ShouldEqual, "0.0.0.0:4321")

				Convey("Health endpoint should return 200", func() {
					req := httptest.NewRequest("GET", "http://0.0.0.0:4321/healthz", nil)
					res := httptest.NewRecorder()
					svc.Handler.ServeHTTP(res, req)
					So(res.Code, ShouldEqual, 200)
					So(res.Header().Get("health"), ShouldEqual, "testing")
				})

				Convey("Readiness endpoint should return 200", func() {
					req := httptest.NewRequest("GET", "http://0.0.0.0:4321/readyz", nil)
					res := httptest.NewRecorder()
					svc.Handler.ServeHTTP(res, req)
					So(res.Code, ShouldEqual, 200)
					So(res.Header().Get("ready"), ShouldEqual, "testing")
				})
			})
		})
	})
}
