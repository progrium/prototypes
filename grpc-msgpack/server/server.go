package main

import (
	"context"
	"fmt"
	"net"
	"strings"

	grpcmsgpack "github.com/progrium/prototypes/grpc-msgpack"
	"google.golang.org/grpc"
)

type Server interface {
	Upper(string) string
}

type server struct{}

func (s *server) Upper(str string) string {
	return strings.ToUpper(str)
}

var serviceDesc = grpc.ServiceDesc{
	ServiceName: "Demo",
	HandlerType: (*Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upper",
			Handler: func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				var str string
				if err := dec(&str); err != nil {
					return nil, err
				}
				if interceptor == nil {
					return srv.(Server).Upper(str), nil
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/Demo/Upper",
				}
				handler := func(ctx context.Context, req interface{}) (interface{}, error) {
					return srv.(Server).Upper(*(req.(*string))), nil
				}
				return interceptor(ctx, &str, info, handler)
			},
		},
	},
}

func main() {
	gs := grpc.NewServer(grpc.CustomCodec(&grpcmsgpack.MsgPackCodec{}))
	gs.RegisterService(&serviceDesc, &server{})
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("listening on 8080")
	gs.Serve(ln)
}
