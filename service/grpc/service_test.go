package grpc

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGRPCService(t *testing.T) {
	Convey("Given some GRPC Service Options", t, func() {
		ctx := context.Background()
		ropts := &[]Option{
			Name("testing"),
			Namespace("testing"),
			Version("0.0.0"),
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
			So(svc.Name(), ShouldEqual, "testing.testing")
			So(svc.Options().Context, ShouldEqual, ctx)
		})
	})
}
