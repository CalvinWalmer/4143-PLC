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
		downloadImage(urls[i], "S"+strconv.Itoa(i)+"-"+time.Now().Format("20060102150405.000000000")+".jpg")
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
			downloadImage(urls[i], "C"+strconv.Itoa(i)+"-"+time.Now().Format("20060102150405.000000000")+".jpg")
		}()
	}

	wg.Wait()
}

func main() {
	urls := []string{
		"https://unsplash.com/photos/hvdnff_bieQ/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://unsplash.com/photos/HQaZKCDaax0/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://images.unsplash.com/photo-1698778573682-346d219402b5?ixlib=rb-4.0.3&q=85&fm=jpg&crop=entropy&cs=srgb&w=640",
		"https://unsplash.com/photos/Bs2jGUWu4f8/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
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
