## P02 - Baby Steps
### Calvin Walmer
### Description:

This program is used to show how to create a package that will be used inside of a module. 
The package draws a rectangle over an image.

### Files

|   #   | File            | Description                                        |
| :---: | --------------- | -------------------------------------------------- |
|   1   | main.go         | Main driver of my project |
|   2   | go.mod | A manifest defining my projects' dependencies,     |
|   3   | go.sum|  Contains cryptographic hashes of specific module versions defined in go.mod |
|   4   | mustangs.jpg | image before manipulation |
|   5   | mustangs.png | image afteer manipulation |
|   6   | imageManipulator/imageManipulator.go | Package that uses gg to import and then modify an image. This package is imported and used in main.go |
