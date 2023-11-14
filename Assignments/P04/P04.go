// =========================================================================
// Andrew Harrison
// 4143 - PLC
// P04 - Concurrent Image Downloader
// 11/12/2023
// =========================================================================

// =========================================================================
// Assignment Description:
// Understand and implement basic concurrency in Go. Write a program that
// concurrently downloads a set of images from given URLs and saves them to
// disk. By comparing the time taken to download images sequentially vs.
// concurrently, you will observe the benefits of concurrency for I/O-bound
// tasks.
// =========================================================================

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// =========================================================================
// Sequential version of the image downloader.
func downloadImagesSequential(urls []string) {
	for i, url := range urls {
		// Format filename as image_<index>.jpg for each image
		filename := fmt.Sprintf("image_%d.jpg", i)
		// Attempt to download the image from the URL
		err := downloadImage(url, filename)
		// Handle errors in downloading, continue to next URL on failure
		if err != nil {
			fmt.Printf("Failed to download %s: %v\n", url, err)
			continue
		}
		// Confirm successful download
		fmt.Printf("Downloaded %s to %s\n", url, filename)
	}
}

// =========================================================================

// =========================================================================
// Concurrent version of the image downloader.
func downloadImagesConcurrent(urls []string) {
	// Declare a wait group for goroutines
	var wg sync.WaitGroup
	for i, url := range urls {
		// Increment wait group counter for each goroutine
		wg.Add(1)
		// Launch a goroutine for each URL
		go func(url string, i int) {
			// Decrement wait group counter when goroutine completes
			defer wg.Done()
			// Format filename as image_<index>.jpg for each image
			filename := fmt.Sprintf("image_%d.jpg", i)
			// Attempt to download the image from the URL
			err := downloadImage(url, filename)
			// Handle errors in downloading, exit goroutine on failure
			if err != nil {
				fmt.Printf("Failed to download %s: %v\n", url, err)
				return
			}
			// Confirm successful download
			fmt.Printf("Downloaded %s to %s\n", url, filename)
		}(url, i)
	}
	// Wait for all goroutines to complete
	wg.Wait()
}

// =========================================================================

// =========================================================================
// Main Function
func main() {
	// List of image URLs to download
	urls := []string{
		// Images Already on Here
		"https://unsplash.com/photos/hvdnff_bieQ/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://unsplash.com/photos/HQaZKCDaax0/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://images.unsplash.com/photo-1698778573682-346d219402b5?ixlib=rb-4.0.3&q=85&fm=jpg&crop=entropy&cs=srgb&w=640",
		"https://unsplash.com/photos/Bs2jGUWu4f8/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		// My 5 Images
		"https://unsplash.com/photos/IigBPm2VWUM/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://unsplash.com/photos/IvJBpwlF9EQ/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://www.pexels.com/photo/18705295/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://www.pexels.com/photo/18346899/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://unsplash.com/photos/M7iMdnG4R_g/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
	}

	// Sequential download
	start := time.Now()
	downloadImagesSequential(urls)
	fmt.Printf("Sequential download took: %v\n", time.Since(start))

	// Concurrent download
	start = time.Now()
	downloadImagesConcurrent(urls)
	fmt.Printf("Concurrent download took: %v\n", time.Since(start))
}

// =========================================================================

// =========================================================================
// Helper function to download and save a single image.
func downloadImage(url, filename string) error {
	// Send an HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		// Return error if GET request fails
		return err
	}
	// Ensure response body is closed after function returns
	defer resp.Body.Close()

	// Create a new file with the given filename
	file, err := os.Create(filename)
	if err != nil {
		// Return error if file creation fails
		return err
	}
	// Ensure file is closed after function returns
	defer file.Close()

	// Copy the response body to the file
	_, err = io.Copy(file, resp.Body)
	// Return any error encountered during copy
	return err
}

// =========================================================================
