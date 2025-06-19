package handlers

import (
	"context"
	"read_product/database"
	"read_product/models"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadProducts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := database.ProductCollection.Find(ctx, map[string]interface{}{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products"})
		return
	}
	defer cursor.Close(ctx)

	var products []models.Product
	if err := cursor.All(ctx, &products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding products"})
		return
	}

	c.JSON(http.StatusOK, products)
}