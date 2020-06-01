package api

import (
	"context"
	"net/http"
	pb "timer-api-service/pb"
	"timer-api-service/services/timer"

	"github.com/gin-gonic/gin"
)

type CreateTimerRequest struct {
	Callback CallbackAction `json:"callback"`
	Payload  string         `json:"payload"`
}

type CallbackAction struct {
	Type    string            `json:"type"`
	Method  string            `json:"method"`
	Uri     string            `json:"uri"`
	Headers map[string]string `json:"headers"`
	Payload string            `json:"payload"`
}

func InitTimerAPI(timer *timer.Service, r *gin.Engine) {

	// Create timer
	r.POST("/api/timers", func(c *gin.Context) {

		var request CreateTimerRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		reply, err := timer.CreateTimer(context.Background(), &pb.CreateTimerRequest{
			Payload: request.Payload,
			Callback: &pb.CallbackAction{
				Type:    request.Callback.Type,
				Method:  request.Callback.Method,
				Uri:     request.Callback.Uri,
				Headers: request.Callback.Headers,
				Payload: request.Callback.Payload,
			},
		})
		if err != nil {

			c.JSON(400, gin.H{
				"success": false,
			})

			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"timerID": reply.TimerID,
		})
	})

	// Cancel timer
	r.DELETE("/api/timer/:timerID", func(c *gin.Context) {

		in := &pb.DeleteTimerRequest{
			TimerID: c.Param("timerID"),
		}

		_, err := timer.DeleteTimer(context.Background(), in)
		if err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})
}
