package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Routes struct {
	Router   *gin.Engine
	Database *gorm.DB
}

func (r *Routes) SetAllRoutes() {
	r.setupHealthRoute()
	r.setupTransactionRoutes()
}
