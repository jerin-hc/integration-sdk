package integratonsdk

import (
	"context"

	"github.com/hashicorp/go-plugin"
	grpcplugin "github.com/jerin-hc/integration-sdk/grpc-plugin"
	"github.com/jerin-hc/integration-sdk/schema"
	"google.golang.org/grpc"
)

const (
	grpcMaxMessageSize        = 256 << 20
	protocolVersionMajor uint = 5
)

func New(id string) *Serve {
	return &Serve{}
}

type Serve struct {
	Event      schema.Event
	Resources  []schema.Resource
	Ctx        schema.Ctx
	handleFunc func(event schema.Event, resources []schema.Resource, ctx schema.Ctx) *schema.ResourceResponse
}

func (s *Serve) HandleFunc(ctx context.Context, req *schema.HandleFuncRequest) (*schema.ResourceResponse, error) {
	return s.handleFunc(req.Event, req.Resources, ctx), nil
}

func (s *Serve) Handle(event []schema.Event, handleFunc func(event schema.Event, resources []schema.Resource, ctx schema.Ctx) *schema.ResourceResponse) {
	s.handleFunc = handleFunc
}

func (s *Serve) Run() {
	serveConfig := &plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  protocolVersionMajor,
			MagicCookieKey:   "TF_PLUGIN_MAGIC_COOKIE",
			MagicCookieValue: "d602bf8f470bc67ca7faa0386276bbdd4330efaf76d1a219cb4d6991ca9872b2",
		},
		Plugins: plugin.PluginSet{
			"provider": &grpcplugin.HandlerPlugin{
				IntegrationServer: s,
			},
		},
		GRPCServer: func(opts []grpc.ServerOption) *grpc.Server {
			opts = append(opts, grpc.MaxRecvMsgSize(grpcMaxMessageSize))
			opts = append(opts, grpc.MaxSendMsgSize(grpcMaxMessageSize))

			return grpc.NewServer(opts...)
		},
	}

	plugin.Serve(serveConfig)
}
