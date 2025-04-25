package main

import (
	"ai_project/config"
	"ai_project/internal/database"
	"ai_project/internal/models"
	"ai_project/internal/repositories"
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGODB_URI environment variable is not set")
	}

	if err := database.InitializeMongoClient(mongoURI); err != nil {
		log.Fatalf("Failed to initialize MongoDB: %v", err)
	}

	router := gin.Default()

	router.POST("/lender", func(c *gin.Context) {
		var jsonArray []map[string]interface{}

		if c.GetHeader("secret_key") != cfg.SecretKey {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Forbidden: Invalid secret key",
			})
			return
		}

		if err := c.ShouldBindJSON(&jsonArray); err != nil {
			log.Printf("JSON bind error in /lender: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON input",
			})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		collection := database.MongoClient.Database(cfg.DatabaseName).Collection("lenders")

		for _, jsonData := range jsonArray {
			jsonData["createdAt"] = time.Now()
			if _, err := collection.InsertOne(ctx, bson.M(jsonData)); err != nil {
				log.Printf("Mongo insert error for item: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to store some data",
				})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Data stored successfully",
		})
	})

	router.POST("/lender-offer", func(c *gin.Context) {
		var jsonArray []map[string]interface{}

		if c.GetHeader("secret_key") != cfg.SecretKey {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Forbidden: Invalid secret key",
			})
			return
		}

		if err := c.ShouldBindJSON(&jsonArray); err != nil {
			log.Printf("JSON bind error in /lender-offer: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON array input",
			})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		collection := database.MongoClient.Database(cfg.DatabaseName).Collection("lender_offers")

		for _, jsonData := range jsonArray {
			jsonData["createdAt"] = time.Now()
			if _, err := collection.InsertOne(ctx, bson.M(jsonData)); err != nil {
				log.Printf("Mongo insert error for item: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to store some data",
				})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "All data stored successfully",
		})
	})

	router.PUT("/lender/update", func(c *gin.Context) {
		if c.GetHeader("secret_key") != cfg.SecretKey {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Forbidden: Invalid secret key",
			})
			return
		}

		var updateData models.Lender
		if err := c.ShouldBindJSON(&updateData); err != nil {
			log.Printf("JSON bind error in /lender/update: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON input",
			})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		db := database.MongoClient.Database(cfg.DatabaseName)
		lenderRepo := repositories.NewLenderRepo(db)

		result, err := lenderRepo.UpdateByName(ctx, updateData.LenderName, updateData)
		if err != nil {
			log.Printf("Mongo update error in /lender: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update data",
			})
			return
		}

		if result.MatchedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Document not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Data updated successfully",
		})
	})

	if err := router.Run(":8009"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
