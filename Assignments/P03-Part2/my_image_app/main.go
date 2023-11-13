package main

import (
	"fmt"
	"net/http"

	"github.com/yourUserName/img_mod/Colors"
	"github.com/yourUserName/img_mod/GetPic"
	"github.com/yourUserName/img_mod/Grayscale"
	"github.com/yourUserName/img_mod/Text"
)

func handler(w http.ResponseWriter, r *http.Request) {
	imageURL := r.URL.Query().Get("url")
	if imageURL == "" {
		http.Error(w, "Missing url parameter", http.StatusBadRequest)
		return
	}

	// Download the image
	GetPic.DownloadImage(imageURL)

	// Process the image
	Colors.Execute()
	Grayscale.Execute()
	Text.Execute()

	// You would typically return the processed image here.
	// For demonstration, we're just indicating success.
	fmt.Fprintf(w, "Successfully processed image from: %s", imageURL)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
