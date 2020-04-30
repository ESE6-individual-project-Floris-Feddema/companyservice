package controllers

import (
	. "companyservice/contexts"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthController struct{}

func (controller HealthController) GetHealth(c *gin.Context) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	client := GetClient(ctx);
	if client == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "unavailable"})
		return
	}

	err := client.Ping(ctx, nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "unavailable"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "available"})
}

