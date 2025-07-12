package grpcplugin

import (
	"context"
	"errors"
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"github.com/jerin-hc/integration-sdk/schema"
	"google.golang.org/grpc"
)

type HandlerPlugin struct {
	IntegrationServer IntegrationServer
}

type integrationClient struct {
	cc *grpc.ClientConn
}

func (c *integrationClient) HandleFunc(ctx context.Context, req *schema.HandleFuncRequest) (*schema.ResourceResponse, error) {
	resp := new(schema.ResourceResponse)
	err := c.cc.Invoke(ctx, "/IntegrationService/HandleFunc", req, resp, grpc.CallContentSubtype("json"))
	return resp, err
}

func (p *HandlerPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	s.RegisterService(&Provider_ServiceDesc, p.IntegrationServer)
	return nil
}

func (p *HandlerPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, cc *grpc.ClientConn) (any, error) {
	return &integrationClient{cc: cc}, nil
}

func (p *HandlerPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return nil, errors.New("terraform-plugin-go only implements gRPC servers")
}

func (p *HandlerPlugin) Client(*plugin.MuxBroker, *rpc.Client) (interface{}, error) {
	return nil, errors.New("terraform-plugin-go only implements gRPC servers")
}

type IntegrationServer interface {
	HandleFunc(context.Context, *schema.HandleFuncRequest) (*schema.ResourceResponse, error)
}

var Provider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "IntegrationService",
	HandlerType: (*IntegrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleFunc",
			Handler: func(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {

				hr := &schema.HandleFuncRequest{
					Event: schema.PostApply,
				}

				// interceptor wrapper (optional)
				if interceptor == nil {
					return srv.(IntegrationServer).HandleFunc(ctx, hr)
				}

				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/IntegrationService/HandleFunc",
				}

				handler := func(ctx context.Context, req any) (any, error) {
					return srv.(IntegrationServer).HandleFunc(ctx, hr)
				}

				return interceptor(ctx, hr, info, handler)
			},
		},
	},
	Metadata: "tfplugin5.proto",
}
