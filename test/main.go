package main

import (
	"context"
	"log"
	"os/exec"

	"github.com/hashicorp/go-plugin"
	grpcplugin "github.com/jerin-hc/integration-sdk/grpc-plugin"
	"github.com/jerin-hc/integration-sdk/jsoncodec"
	"github.com/jerin-hc/integration-sdk/schema"
)

func main() {

	jsoncodec.Init()

	var myHandshake = plugin.HandshakeConfig{
		ProtocolVersion:  5,
		MagicCookieKey:   "TF_PLUGIN_MAGIC_COOKIE",
		MagicCookieValue: "d602bf8f470bc67ca7faa0386276bbdd4330efaf76d1a219cb4d6991ca9872b2",
	}

	pluginMap := map[string]plugin.Plugin{
		"integration": &grpcplugin.HandlerPlugin{},
	}

	pluginPath := "/Users/jerin/code/integrations-sdk/myplugin"

	pluginClient := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  myHandshake,
		Plugins:          pluginMap,
		Cmd:              exec.Command(pluginPath),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	})
	rpcClient, err := pluginClient.Client()
	if err != nil {
		log.Println("ERROR :: ")
		log.Fatal(err)
	}

	raw, err := rpcClient.Dispense("integration")
	if err != nil {
		log.Fatal("failed to dispense plugin:", err)
	}

	integration, ok := raw.(grpcplugin.IntegrationServer)

	if !ok {
		log.Fatal("NOT FOUND")
	}

	resp, err := integration.HandleFunc(context.Background(), &schema.HandleFuncRequest{
		Event:     schema.PostPlan,
		Resources: nil,
	})

	log.Println(*resp)
}
