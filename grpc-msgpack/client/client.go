package main

import (
	"context"
	"fmt"

	grpcmsgpack "github.com/progrium/prototypes/grpc-msgpack"
	"google.golang.org/grpc"
)

type Client interface {
	Upper(ctx context.Context, str string, opts ...grpc.CallOption) (string, error)
}

type client struct {
	cc *grpc.ClientConn
}

func (c *client) Upper(ctx context.Context, str string, opts ...grpc.CallOption) (string, error) {
	var ret string
	err := c.cc.Invoke(ctx, "/Demo/Upper", str, &ret)
	if err != nil {
		return "", err
	}
	return ret, nil
}

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithCodec(&grpcmsgpack.MsgPackCodec{}), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	c := &client{conn}
	str, err := c.Upper(context.Background(), "hello world")
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}
