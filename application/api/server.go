package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jvfrodrigues/transaction-product-wex/application/api/routes"
)

type Server struct {
	router *gin.Engine
}

func NewServer(mode string) *Server {
	setServerMode(mode)
	server := &Server{}
	return server
}

func (server *Server) setupRouter(database *gorm.DB) {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers, Authorization"},
	}))
	routes := routes.Routes{Router: router, Database: database}
	routes.SetAllRoutes()
	server.router = router
}

func (server *Server) StartServer(database *gorm.DB, address string) error {
	server.setupRouter(database)
	return server.router.Run(address)
}

func setServerMode(mode string) {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else if mode == gin.DebugMode {
		gin.SetMode(mode)
	}
}
