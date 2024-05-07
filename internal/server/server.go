package server

import (
	"github.com/gin-gonic/gin"
	"github.com/leonlatsch/go-resolve/internal/service"
)

type Server struct {
	Service service.DnsModeService
}

func (self *Server) StartApiServer() {
	router := gin.Default()

	router.GET("/update", func(ctx *gin.Context) {
		ip := ctx.Query("ip")
		self.Service.UpdateDns(ip)
		ctx.Status(200)
	})

	router.Run("0.0.0.0:9991")
}
