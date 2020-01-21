package http

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHTTPService(t *testing.T) {
	Convey("Given some HTTP Service Options", t, func() {
		ctx := context.Background()
		ropts := &[]Option{
			Name("testing"),
			Version("0.0.0"),
			Namespace("testing"),
			Context(ctx),
			Address("0.0.0.0:4321"),
		}
		opts := newOptions(*ropts...)
		Convey("Options are initialized properly", func() {
			So(opts.Context, ShouldEqual, ctx)
			So(opts.Version, ShouldEqual, "0.0.0")
			So(opts.Address, ShouldEqual, "0.0.0.0:4321")
			So(opts.Namespace, ShouldEqual, "testing")
			So(opts.Name, ShouldEqual, "testing")
		})

		Convey("New Service is initialized properly", func() {
			svc := NewService(*ropts...).Service
			So(svc.Options().Name, ShouldEqual, "testing.testing")
			So(svc.Options().Version, ShouldEqual, "0.0.0")
			So(svc.Options().Address, ShouldEqual, "0.0.0.0:4321")
		})
	})
}
