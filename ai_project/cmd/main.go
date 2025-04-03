package main

import (
	"ai_project/config"
	"ai_project/internal/database"
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	// POST /lender endpoint
	router.POST("/lender", func(c *gin.Context) {
		var jsonData map[string]interface{}

		if err := c.ShouldBindJSON(&jsonData); err != nil {
			log.Printf("JSON bind error in /lender: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON input",
			})
			return
		}

		// Add timestamp
		jsonData["createdAt"] = time.Now()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		collection := database.MongoClient.Database(cfg.DatabaseName).Collection("lenders")
		_, err := collection.InsertOne(ctx, bson.M(jsonData))
		if err != nil {
			log.Printf("Mongo insert error in /lender: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to store data",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Data stored successfully",
		})
	})

	// POST /lender-offer endpoint
	router.POST("/lender-offer", func(c *gin.Context) {
		var jsonData map[string]interface{}

		if err := c.ShouldBindJSON(&jsonData); err != nil {
			log.Printf("JSON bind error in /lender-offer: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON input",
			})
			return
		}

		// Add timestamp
		jsonData["createdAt"] = time.Now()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		collection := database.MongoClient.Database(cfg.DatabaseName).Collection("lender_offers")
		_, err := collection.InsertOne(ctx, bson.M(jsonData))
		if err != nil {
			log.Printf("Mongo insert error in /lender-offer: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to store data",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Data stored successfully",
		})
	})

	// PUT /lender/update/:id endpoint
	router.PUT("/lender/update/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		objectID, err := primitive.ObjectIDFromHex(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID format",
			})
			return
		}

		var updateData map[string]interface{}
		if err := c.ShouldBindJSON(&updateData); err != nil {
			log.Printf("JSON bind error in /lender/update: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON input",
			})
			return
		}

		// Optional: update modifiedAt timestamp
		updateData["updatedAt"] = time.Now()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		collection := database.MongoClient.Database(cfg.DatabaseName).Collection("lenders")

		update := bson.M{
			"$set": updateData,
		}

		opts := options.Update().SetUpsert(false) // avoid inserting new document if not found

		result, err := collection.UpdateByID(ctx, objectID, update, opts)
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
