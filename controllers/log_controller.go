package controllers

import (
	"go-crud-api/config"
	"go-crud-api/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateLog(c *gin.Context) {
	var log models.Log
	if err := c.ShouldBindJSON(&log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Timestamp = time.Now() 
	config.DB.Create(&log)
	c.JSON(http.StatusOK, log)
}

func GetLogs(c *gin.Context) {
	// var logs []models.Log
	// config.DB.Find(&logs)
	// c.JSON(http.StatusOK, logs)
	var logs []models.Log
	query := config.DB

	// Filtering by timestamp
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")
	if startTime != "" && endTime != "" {
		start, err := time.Parse(time.RFC3339, startTime)
		if err == nil {
			end, err := time.Parse(time.RFC3339, endTime)
			if err == nil {
				query = query.Where("timestamp BETWEEN ? AND ?", start, end)
			}
		}
	}

	// Filtering by severity
	severity := c.Query("severity")
	if severity != "" {
		query = query.Where("severity = ?", severity)
	}

	// Filtering by service name
	serviceName := c.Query("service_name")
	if serviceName != "" {
		query = query.Where("service_name = ?", serviceName)
	}

	// Limiting the number of results
	count := c.Query("count")
	if count != "" {
		limit, err := strconv.Atoi(count)
		if err == nil {
			query = query.Limit(limit)
		}
	}

	query.Find(&logs)
	c.JSON(http.StatusOK, logs)
}

func GetLog(c *gin.Context) {
	id := c.Param("id")
	var log models.Log
	if err := config.DB.First(&log, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Log not found!"})
		return
	}
	c.JSON(http.StatusOK, log)
}


func UpdateLog(c *gin.Context) {
	id := c.Param("id")
	var log models.Log
	if err := config.DB.First(&log, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Log not found!"})
		return
	}
	if err := c.ShouldBindJSON(&log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&log)
	c.JSON(http.StatusOK, log)
}

func DeleteLog(c *gin.Context) {
	id := c.Param("id")
	var log models.Log
	if err := config.DB.First(&log, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Log not found!"})
		return
	}
	config.DB.Delete(&log)
	c.JSON(http.StatusOK, gin.H{"message": "Log deleted"})
}