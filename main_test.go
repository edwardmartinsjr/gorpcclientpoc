package main

import (
	"testing"

	micro "github.com/micro/go-micro"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewGreeterClient(t *testing.T) {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
	)

	Convey("Using a default ServiceName ", t, func() {
		So(NewGreeterClient("", service.Client()), ShouldNotBeNil)
	})
}
