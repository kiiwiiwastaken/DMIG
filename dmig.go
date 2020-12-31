/*
Download Manager In Go

This is a download manager I'm making all withing go, right now all of 
func DownloadFile() is from the following url: 
https://golangcode.com/download-a-file-from-a-url/

I'm wanting to add more onto this, but this right now is just so I can have a
jumping off point for *my* code.
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Asks the user what the url to the file is.
	// Input will be below the following text.
	fmt.Println("Url you would like to download from: ")

	// var then variable name then the type.
	var fileUrl string

	// Takes input from user.
	fmt.Scanln(&fileUrl)
	fmt.Println("File name: ")
	var fileName string
	fmt.Scanln(&fileName)

	// Stuff from the older main function.
	err := DownloadFile(fileName, fileUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded: " + fileUrl)
}

// DownloadedFile will download a url to a local file. It's efficient because it
// write as it downloads and not load the whole file into memory
func DownloadFile(filepath string, url string) error {
	
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}