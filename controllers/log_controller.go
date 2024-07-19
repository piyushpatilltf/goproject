package controllers

import (
	"context"
	"encoding/json"
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
	cacheKey := "logs"
	val, err := config.Rdb.Get(context.Background(), cacheKey).Result()
    if err == nil {
        var cachedLogs []models.Log
        err := json.Unmarshal([]byte(val), &cachedLogs)
        if err == nil {
            c.JSON(http.StatusOK, cachedLogs)
            return
        }
    }

	var logs []models.Log
	query := config.DB

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

	severity := c.Query("severity")
	if severity != "" {
		query = query.Where("severity = ?", severity)
	}

	serviceName := c.Query("service_name")
	if serviceName != "" {
		query = query.Where("service_name = ?", serviceName)
	}

	count := c.Query("count")
	if count != "" {
		limit, err := strconv.Atoi(count)
		if err == nil {
			query = query.Limit(limit)
		}
	}

	query.Find(&logs)

	logBytes, err := json.Marshal(logs)
    if err == nil {
        err = config.Rdb.Set(context.Background(), cacheKey, logBytes, time.Minute*5).Err() 
        if err != nil {
           c.JSON(http.StatusNotFound,gin.H{"error": "Log not found!"})
		   return
        }
    }

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




