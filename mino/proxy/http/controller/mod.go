package controller

import (
	"go.dedis.ch/dela/cli"
	"go.dedis.ch/dela/cli/node"
	"go.dedis.ch/dela/mino/proxy/http"
)

const defaultAddr = "127.0.0.1:8080"

// NewMinimal returns a new minimal initializer
func NewMinimal() node.Initializer {
	return minimal{}
}

// minimal is an initializer with the minimum set of commands. Indeed it only
// creates and injects a new client proxy
//
// - implements node.Initializer
type minimal struct{}

// Build implements node.Initializer. In this case we don't need any command.
func (m minimal) SetCommands(builder node.Builder) {
	cmd := builder.SetCommand("proxy")
	sub := cmd.SetSubCommand("start")

	sub.SetDescription("start the proxy http server")
	sub.SetFlags(cli.StringFlag{
		Name:     "clientaddr",
		Required: false,
		Usage:    "the address of the http client",
		Value:    defaultAddr,
	})
	sub.SetAction(builder.MakeAction(startAction{}))
}

// OnStart implements node.Initializer. It creates, starts, and registers a
// client proxy.
func (m minimal) OnStart(ctx cli.Flags, inj node.Injector) error {
	return nil
}

// OnStop implements node.Initializer. It stops the http server.
func (m minimal) OnStop(inj node.Injector) error {
	var proxy *http.HTTP
	err := inj.Resolve(&proxy)
	if err == nil {
		proxy.Stop()
	}

	return nil
}
