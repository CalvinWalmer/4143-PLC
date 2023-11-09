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
		"https://images.pexels.com/photos/3509410/pexels-photo-3509410.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2",
		"https://images.unsplash.com/photo-1570284613060-766c33850e00?q=80&w=3270&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"https://images.unsplash.com/photo-1575470522418-b88b692b8084?q=80&w=3288&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"https://images.unsplash.com/photo-1559289431-9f12ee08f8b6?q=80&w=3387&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"https://images.unsplash.com/photo-1496354265829-17b1c7b7c363?q=80&w=1454&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
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
