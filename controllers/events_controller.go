package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shokoohi/get-ethereum-transactions/utils"
)

// EventLogQuery - Contract's event logs query for binding
type EventLogQuery struct {
	Address   string `json:"address"`
	FromBlock int64  `json:"from_block"`
	ToBlock   int64  `json:"to_block"`
}

// ReadEventsLogs - Return all or a range of contract's event logs
func ReadEventsLogs(c *gin.Context) {
	var event EventLogQuery
	// Bind request query data to our event query
	if err := c.ShouldBindJSON(&event); err != nil {
		log.Println("Can't bind event logs query:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"status": "bad",
			"error":  "Can't bind event logs query: " + err.Error(),
		})
		return
	}
	// Get contract events logs
	logs, err := utils.ReadEventsLogs(context.Background(), event.Address, event.FromBlock, event.ToBlock)
	if err != nil {
		log.Println("Can't get event logs:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"status": "bad",
			"error":  "Can't get event logs: " + err.Error(),
		})
		return
	}
	// Send logs at response
	c.JSON(http.StatusOK, logs)
}
