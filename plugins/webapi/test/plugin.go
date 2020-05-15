package test

import (
	"github.com/iotaledger/goshimmer/plugins/webapi"
	"github.com/iotaledger/goshimmer/plugins/webapi/test/drng"
	"github.com/iotaledger/goshimmer/plugins/webapi/test/value"
	"github.com/iotaledger/hive.go/logger"
	"github.com/iotaledger/hive.go/node"
)

// PluginName is the name of the web API DRNG endpoint plugin.
const PluginName = "WebAPI Test Endpoint"

var (
	// Plugin is the plugin instance of the web API DRNG endpoint plugin.
	Plugin = node.NewPlugin(PluginName, node.Enabled, configure)
	log    *logger.Logger
)

func configure(plugin *node.Plugin) {
	log = logger.NewLogger(PluginName)
	webapi.Server.POST("test/drng", drng.SendDrngMsg)
	webapi.Server.GET("test/value", value.SendValueMsg)
}
