package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Routes) setupHealthRoute() {
	transactionGroup := r.Router.Group("/health")
	{
		transactionGroup.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "OK"})
		})
	}
}
