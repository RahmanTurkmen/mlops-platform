package main

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Job struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	Accuracy  int    `json:"accuracy"`
	CreatedAt string `json:"created_at"`
}

const filePath = "data/jobs.json"

func loadJobs() []Job {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return []Job{}
	}

	var jobs []Job
	json.Unmarshal(file, &jobs)
	return jobs
}

func saveJobs(jobs []Job) {
	data, _ := json.MarshalIndent(jobs, "", "  ")
	os.WriteFile(filePath, data, 0644)
}

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	r.GET("/jobs", func(c *gin.Context) {
		c.JSON(200, loadJobs())
	})

	r.POST("/jobs", func(c *gin.Context) {
		jobs := loadJobs()

		var newJob Job
		if err := c.BindJSON(&newJob); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		newJob.ID = len(jobs) + 1
		newJob.Status = "running"
		newJob.Accuracy = 0
		newJob.CreatedAt = time.Now().Format("2006-01-02")

		jobs = append(jobs, newJob)
		saveJobs(jobs)

		go simulateTraining(newJob.ID)

		c.JSON(201, newJob)
	})

	r.DELETE("/jobs/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		jobs := loadJobs()
		var updated []Job

		found := false

		for _, job := range jobs {
			if job.ID != id {
				updated = append(updated, job)
			} else {
				found = true
			}
		}

		if !found {
			c.JSON(404, gin.H{"error": "not found"})
			return
		}

		saveJobs(updated)
		c.JSON(200, gin.H{"message": "deleted"})
	})

	r.PATCH("/jobs/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		var body Job
		c.BindJSON(&body)

		jobs := loadJobs()
		found := false

		for i, job := range jobs {
			if job.ID == id {
				found = true

				if body.Status != "" {
					jobs[i].Status = body.Status
				}
				if body.Accuracy != 0 {
					jobs[i].Accuracy = body.Accuracy
				}
			}
		}

		if !found {
			c.JSON(404, gin.H{"error": "not found"})
			return
		}

		saveJobs(jobs)
		c.JSON(200, gin.H{"message": "updated"})
	})

	r.Run(":8080")
}

func simulateTraining(id int) {
	for i := 0; i <= 100; i += 10 {
		time.Sleep(1 * time.Second)

		jobs := loadJobs()

		for j, job := range jobs {
			if job.ID == id {
				jobs[j].Accuracy = i

				if i == 100 {
					jobs[j].Status = "completed"
				}
			}
		}

		saveJobs(jobs)
	}
}
