package main

import (
	"fmt"
	"image-to-asciiart/asciiart"
	"image-to-asciiart/imageutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Switch to "release" mode in production
	gin.SetMode(gin.ReleaseMode)

	// Create a Gin router
	r := gin.Default()

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Define a route to display the form
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Define a route to handle image downloads
	r.POST("/download", func(c *gin.Context) {
		imageURL := c.PostForm("imageURL")
		fmt.Printf("Image URL: %s\n", imageURL) // Debug log

		err := imageutil.DownloadImage(imageURL, "downloaded_image.jpg")
		if err != nil {
			fmt.Printf("Error downloading: %s\n", err) // Debug log
			c.String(http.StatusInternalServerError, "Failed to download image")
			return
		}

		asciiArt, err := asciiart.ConvertToASCII("downloaded_image.jpg", 100, 50)
		if err != nil {
			fmt.Printf("Error converting to ASCII: %s\n", err) // Debug log
			c.String(http.StatusInternalServerError, "Failed to convert to ASCII")
			return
		}

		// Write to output.txt
		file, err := os.Create("output.txt")
		if err != nil {
			fmt.Printf("Error creating output.txt: %s\n", err)
			c.String(http.StatusInternalServerError, "Failed to create output file")
			return
		}
		defer file.Close()
		file.WriteString(asciiArt)
		file.Sync()

		fmt.Printf("ASCII Art:\n%s\n", asciiArt) // Debug log
		c.String(http.StatusOK, asciiArt)
	})

	// Run the web server on port 8080
	r.Run(":8080")
}
