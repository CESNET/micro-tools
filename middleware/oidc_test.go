package middleware

import (
	"github.com/CESNET/micro-tools/oidc"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestOIDC(t *testing.T) {
	Convey("Given some OIDC Options", t, func() {
		ropts := &[]oidc.Option{
			oidc.Realm("testing"),
			oidc.Endpoint("http://test/oidc"),
			oidc.Insecure(true),
		}
		opts := newOIDCOptions(*ropts...)
		Convey("Options are initialized properly", func() {
			So(opts.SigningAlgs, ShouldBeEmpty)
			So(opts.Insecure, ShouldBeTrue)
			So(opts.Endpoint, ShouldEqual, "http://test/oidc")
			So(opts.Realm, ShouldEqual, "testing")
		})

		Convey("OpenIDConnect handler is created", func() {
			oidcHandler := OpenIDConnect(*ropts...)
			So(oidcHandler, ShouldNotBeNil)
			//	TODO: add more tests
		})
	})
}
