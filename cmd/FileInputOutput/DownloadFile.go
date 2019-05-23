package FileInputOutput

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/dustin/go-humanize"
)

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
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	//log.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	log.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

func DownloadFile(framework string) error {

	newFile, err := os.Create(framework + ".zip")
	if err != nil {
		return err
	}
	defer newFile.Close()

	if framework == "codeigniter" {
		url := "https://codeload.github.com/bcit-ci/CodeIgniter/zip/develop"
		response, err := http.Get(url)
		defer response.Body.Close()
		counter := &WriteCounter{}
		numBytesWritten, err := io.Copy(newFile, io.TeeReader(response.Body, counter))
		if err != nil {
			return err
		}
		log.Printf("Downloaded %d byte file.\n", numBytesWritten)
	} else if framework == "laravel" {
		url := "https://codeload.github.com/laravel/laravel/zip/master"

		response, err := http.Get(url)
		defer response.Body.Close()
		counter := &WriteCounter{}
		numBytesWritten, err := io.Copy(newFile, io.TeeReader(response.Body, counter))
		if err != nil {
			return err
		}
		log.Printf("Downloaded %d byte file.\n", numBytesWritten)
	}

	return err
}

func DownloadInstaller(softwarename string) error {
	newFile, err := os.Create(softwarename + ".exe")
	if err != nil {
		return err
	}
	defer newFile.Close()
	if softwarename == "composer" {
		url := "https://getcomposer.org/Composer-Setup.exe"
		response, err := http.Get(url)
		defer response.Body.Close()
		counter := &WriteCounter{}
		numBytesWritten, err := io.Copy(newFile, io.TeeReader(response.Body, counter))
		if err != nil {
			return err
		}
		log.Printf("Downloaded %d byte file. \n", numBytesWritten)
	} else if softwarename == "npm" || softwarename == "nodejs" {
		url := "https://nodejs.org/dist/v10.15.3/node-v10.15.3-x64.msi"
		response, err := http.Get(url)
		defer response.Body.Close()
		counter := &WriteCounter{}
		numBytesWritten, err := io.Copy(newFile, io.TeeReader(response.Body, counter))
		if err != nil {
			return err
		}
		log.Printf("Downloaded %d byte file. \n", numBytesWritten)
	} else if softwarename == "php" {
		url := "https://www.apachefriends.org/xampp-files/7.1.29/xampp-windows-x64-7.1.29-1-VC14-installer.exe"
		response, err := http.Get(url)
		defer response.Body.Close()
		counter := &WriteCounter{}
		numBytesWritten, err := io.Copy(newFile, io.TeeReader(response.Body, counter))
		if err != nil {
			return err
		}
		log.Printf("Downloaded %d byte file. \n", numBytesWritten)
	} else if softwarename == "git" {
		url := "https://git-scm.com/download/win"
		response, err := http.Get(url)
		defer response.Body.Close()
		counter := &WriteCounter{}
		numBytesWritten, err := io.Copy(newFile, io.TeeReader(response.Body, counter))
		if err != nil {
			return err
		}
		log.Printf("Downloaded %d byte file. \n", numBytesWritten)
	}
	return err
}
