package main

import (
	"context"
	"fmt"

	proto "github.com/edwardmartinsjr/gorpcserverpoc/proto"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
	)

	// create the greeter client using the service name and client
	greeter := NewGreeterClient("greeter", service.Client())

	// request the Hello OR Goodbye method on the Greeter handler
	//rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{
	rsp, err := greeter.Goodbye(context.TODO(), &proto.GoodbyeRequest{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp)
}

type greeterClient struct {
	c           client.Client
	serviceName string
}

// NewGreeterClient -
func NewGreeterClient(serviceName string, c client.Client) GreeterClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "greeter"
	}
	return &greeterClient{
		c:           c,
		serviceName: serviceName,
	}
}

// GreeterClient -
type GreeterClient interface {
	Hello(ctx context.Context, in *proto.HelloRequest, opts ...client.CallOption) (*proto.HelloResponse, error)
	Goodbye(ctx context.Context, in *proto.GoodbyeRequest, opts ...client.CallOption) (*proto.GoodbyeResponse, error)
}

func (c *greeterClient) Hello(ctx context.Context, in *proto.HelloRequest, opts ...client.CallOption) (*proto.HelloResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Greeter.Hello", in)
	out := new(proto.HelloResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Goodbye(ctx context.Context, in *proto.GoodbyeRequest, opts ...client.CallOption) (*proto.GoodbyeResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Greeter.Goodbye", in)
	out := new(proto.GoodbyeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
