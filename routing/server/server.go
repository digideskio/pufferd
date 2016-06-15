package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pufferpanel/pufferd/httphandlers"
	"github.com/pufferpanel/pufferd/permissions"
	"github.com/pufferpanel/pufferd/programs"
)

func RegisterRoutes(e *gin.Engine) {
	l2 := e.Group("/server", httphandlers.AdminServerAccessHandler, httphandlers.HasServerAccessHandler)
	{
		l2.PUT("/:id", CreateServer)
		l2.DELETE("/:id", DeleteServer)
	}

	l1 := e.Group("/server", httphandlers.UserServerAccessHandler)
	{
		l1.GET("/:id/start", StartServer)
		l1.GET("/:id/stop", StopServer)
	}
}

func StartServer(c *gin.Context) {
}

func StopServer(c *gin.Context) {
}

func CreateServer(c *gin.Context) {
	serverId := c.Param("id")
	privKey := c.Query("privkey")
	serverType := c.Query("type")
	data := make(map[string]interface{}, 0)
	data["memory"] = "1024M"

	if !permissions.GetGlobal().HasPermission(privKey, "server.create") {
		c.AbortWithStatus(403)
		return
	}

	existing := programs.GetFromCache(serverId)

	if existing != nil {
		c.AbortWithStatus(409)
		return
	}
	programs.Create(serverId, serverType, data)
}

func DeleteServer(c *gin.Context) {
	serverId := c.Param("id")
	privKey := c.Query("privkey")

	if !permissions.GetGlobal().HasPermission(privKey, "server.delete") {
		c.AbortWithStatus(403)
		return
	}

	existing, _ := programs.Get(serverId)

	if existing == nil {
		c.AbortWithStatus(404)
		return
	}

	programs.Delete(existing.Id())
}