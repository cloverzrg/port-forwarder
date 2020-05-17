package web

import (
	_ "github.com/cloverzrg/go-portforward/docs"
	"github.com/cloverzrg/go-portforward/web/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine) {
	g := r.Group("/v1")
	setupDefaultRoute(g)
}

func setupDefaultRoute(r *gin.RouterGroup) {
	r.GET("/network/interfaces", controller.GetNetworkInterfaces)

	r.POST("/forward/", controller.AddForward)
	r.POST("/forward/:id/stop", controller.StopForward)
	r.POST("/forward/:id/start")
	r.GET("/forward/:id")
	r.DELETE("/forward/:id")
}
