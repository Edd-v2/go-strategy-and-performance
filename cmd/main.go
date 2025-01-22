package main

import (
	"fmt"
	"go-strategy-and-performance/helper"
	"go-strategy-and-performance/sender"

	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	nats_url := "nats://demo.nats.io:4222"

	connManager := sender.NewConnectionManager()
	err := connManager.Connect(nats_url)
	if err != nil {
		log.Fatal("Could not connect to nats server")
	}

	defer connManager.NatsClient.Close()

	init_api(connManager)

}

func init_api(connManager *sender.ConnectionManager) {
	log.Println("Initializing Api Command")
	r := gin.Default()

	r.POST("/strategy", func(c *gin.Context) {
		log.Println("Received request to start a new task")
		var req struct {
			Strategy string `json:"strategy"`
			Workload int    `json:"workload"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Dati non validi"})
			return
		}

		// do smt here
		selected_strategy, err := helper.GetStrategy(req.Strategy)
		if err != nil {
			c.JSON(400, gin.H{"error": "Dati non validi"})
			return
		}
		// async exec
		go func() {
			startTime := time.Now()

			defer func() {
				log.Println("End of Strategy - Task Completed")
			}()

			if err := selected_strategy.Execute(req.Workload); err != nil {
				log.Printf("Error executing strategy: %v", err)
				return
			}

			elapsedTime := time.Since(startTime)

			metrics := fmt.Sprintf(
				"Strategy: %s, Workload: %d, ExecutionTime: %s",
				req.Strategy,
				req.Workload,
				elapsedTime,
			)

			connManager.Publish("metrics", metrics)

			log.Printf("Metrics sent to NATS: %s", metrics)
		}()

		log.Println("Task started successfully")
		c.JSON(http.StatusOK, gin.H{"status": "Task started"})
	})

	r.Run(":8080")

}
