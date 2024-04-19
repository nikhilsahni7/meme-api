package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type MemeResponse struct {
	Success bool     `json:"success"`
	Data    MemeData `json:"data"`
}

type MemeData struct {
	Memes []Meme `json:"memes"`
}

type Meme struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	BoxCount int    `json:"box_count"`
	Captions int    `json:"captions"`
}

var CreatedMemes []Meme

func main() {
	r := gin.Default()

	// check server health
	r.GET("/check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	// Get memes endpoint
	r.GET("/memes", func(c *gin.Context) {
		url := "https://api.imgflip.com/get_memes"
		res, err := http.Get(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		var response MemeResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if response.Success {
			var memes []gin.H
			for _, meme := range response.Data.Memes {
				memes = append(memes, gin.H{
					"name": meme.Name,
					"url":  meme.URL,
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"memes": memes,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch memes",
			})
		}
	})

	// post-endpoint function
	r.POST("/memes", func(c *gin.Context) {
		var meme Meme
		err := c.BindJSON(&meme)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		CreatedMemes = append(CreatedMemes, meme)

		// Save the meme to a database or file (not implemented)
		c.JSON(http.StatusCreated, gin.H{
			"message": "Meme created successfully",
		})
	})
	// get-endpoint function to get created memes

	r.GET("/created-memes", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"memes": CreatedMemes,
		})
	})

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
