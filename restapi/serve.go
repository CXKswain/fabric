package restapi

import (
	"github.com/danielmiessler/fabric/core"
	"github.com/gin-gonic/gin"
)

func Serve(registry *core.PluginRegistry, address string) (err error) {
	r := gin.New()
	r.SetTrustedProxies(nil) // or set to a specific list of trusted proxies

	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Register routes
	fabricDb := registry.Db
	NewPatternsHandler(r, fabricDb.Patterns)
	NewContextsHandler(r, fabricDb.Contexts)
	NewSessionsHandler(r, fabricDb.Sessions)
	NewChatHandler(r, registry, fabricDb)
	NewConfigHandler(r, fabricDb)
	NewModelsHandler(r, registry.VendorManager)

	// Start server
	err = r.Run(address)
	if err != nil {
		return err
	}

	return
}
