package routers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shokoohi/get-ethereum-transactions/controllers"
)

// LoadRouters - Prepare and load endpoint routers
func LoadRouters() *gin.Engine {
	router := gin.New()
	// Routes for '/contracts...' endpoints
	router.GET("/contracts/events/logs", controllers.ReadEventsLogs)
	// Return routers
	return router

}
