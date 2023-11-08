
## P04 - Concurrent Image Downloader
### Calvin Walmer
### Description:

The goal of this assignment is to understand and implement basic concurrency in Go.
This program concurrently downloads a set of images from given URLs and saves them to disk.
By comparing the time taken to download images sequentially vs. concurrently, benefits of concurrency for I/O-bound tasks are observed.

### Files

|   #   | File            | Description                                        |
| :---: | --------------- | -------------------------------------------------- |
|   1   | README.md        | Summary, function, and goals of the program    |
|   2   | go.mod  | A manifest defining my projects' dependencies        |
|   3   | main.go | Main driver of my project that downloads images sequentially and concurrently |
|4 - 14 | x... | Images downloaded with the program. C = Concurrent, S = Sequential |

### Results
Sequential download took: 1.857319 seconds
Concurrent download took: 635.8417 milliseconds
5 images were downloaded using each method.

| Links |
|:--|
| https://stocksnap.io/photo/scenic-snowy-MIDRMSZGTG |
| https://unsplash.com/photos/black-and-gray-cruiser-motorcycle-parked-beside-black-concrete-building-2LTMNCN4nEg |
|	https://unsplash.com/photos/blue-and-yellow-plastic-blocks-HpMihL323k0 |
|	https://unsplash.com/photos/a-night-sky-filled-with-lots-of-stars-fxrwJGMCz_g |
|	https://www.shopify.com/stock-photos/photos/art-gallery-interior?q=interior |

