package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"sync"
	"time"
)

// Sequential version of the image downloader.
func downloadImagesSequential(urls []string) {
	for i := 0; i < len(urls); i++ {
		downloadImage(urls[i], "xS"+strconv.Itoa(i)+"-"+time.Now().Format("20060102150405.000000000")+".jpg")
	}
}

// Concurrent version of the image downloader.
func downloadImagesConcurrent(urls []string) {
	var wg sync.WaitGroup

	for i := 0; i < len(urls); i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			downloadImage(urls[i], "xC"+strconv.Itoa(i)+"-"+time.Now().Format("20060102150405.000000000")+".jpg")
		}()
	}

	wg.Wait()
}

func main() {
	urls := []string{
		"https://stocksnap.io/photo/scenic-snowy-MIDRMSZGTG",
		"https://unsplash.com/photos/black-and-gray-cruiser-motorcycle-parked-beside-black-concrete-building-2LTMNCN4nEg",
		"https://unsplash.com/photos/blue-and-yellow-plastic-blocks-HpMihL323k0",
		"https://unsplash.com/photos/a-night-sky-filled-with-lots-of-stars-fxrwJGMCz_g",
		"https://www.shopify.com/stock-photos/photos/art-gallery-interior?q=interior",
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

// Helper function to download and save a single image.
func downloadImage(url, filename string) {
	// Create a new `http.Request` object.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a new `http.Client` object.
	client := &http.Client{}

	// Do the request and get the response.
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check the response status code.
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Response status code:", resp.StatusCode)
		return
	}

	// Create a new file to save the image to.

	// Getting path of where the program is running
	e, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		return
	}
	filename = path.Dir(e) + "/" + filename

	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	f.Close()
}
