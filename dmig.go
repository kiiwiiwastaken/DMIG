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
	"strings"

	"github.com/dustin/go-humanize"
)

// WriteCounter counts the number of bytes written to it. It implements to the
// io.Writer interface and we pass this into io.TeeReader() which will report
// progress on each write cycle.
type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and
	// remove the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// Humanize package used so the amount is in a human readable way (ex: 100MB)
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

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
	fmt.Println("Downloaded: " + fileName)
}

// DownloadedFile will download a url to a local file. It's efficient because it
// write as it downloads and not load the whole file into memory
func DownloadFile(filepath string, url string) error {
	
	// Create the file, but give it a tmp file extension, this means we won't 
	// overwrite a file unless it's downloaded, the .tmp will be removed when 
	// download completes.
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create progress report and pass it to be used alongside the writer
	counter := &WriteCounter{}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}

	// The progress use the same line so print a new line once it's done.
	fmt.Print("\n")

	// Close the file without defer so it can happen before Rename()
	out.Close()

	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		return err
	}
	return nil
}