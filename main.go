package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aklinkert/go-stringslice"
)

var (
	fileExtRAW, fileExtJPG string
	dirNameRAW, dirNameJPG string
	yes                    bool
)

func init() {
	flag.StringVar(&fileExtRAW, "file-ext-raw", "ARW", "File extension of RAW files")
	flag.StringVar(&fileExtJPG, "file-ext-jpg", "JPG", "File extension of JPG files")
	flag.StringVar(&dirNameRAW, "dir-raw", "./RAW", "Directory of RAW files")
	flag.StringVar(&dirNameJPG, "dir-jpg", "./JPG", "Directory of JPG files")
	flag.BoolVar(&yes, "yes", false, "Execute file deletion. Will only list deleted files otherwise")
	flag.Parse()
}

func main() {
	if _, err := os.Stat(dirNameJPG); os.IsNotExist(err) {
		log.Fatalf("Directory %v does not exist", dirNameJPG)
	}
	if _, err := os.Stat(dirNameRAW); os.IsNotExist(err) {
		log.Fatalf("Directory %v does not exist", dirNameRAW)
	}

	filesJPG := listFiles(dirNameJPG, fileExtJPG)
	log.Printf("Found %v JPG files", len(filesJPG))
	filesRAW := listFiles(dirNameRAW, fileExtRAW)
	log.Printf("Found %v RAW files", len(filesRAW))

	for _, jpgFile := range filesJPG {
		rawFile := strings.ReplaceAll(jpgFile, fileExtJPG, fileExtRAW)
		if !stringslice.Contains(filesRAW, rawFile) {
			log.Fatalf("Found missing RAW for JPG %v", jpgFile)
		}
	}

		deleteFiles := make([]string, 0)
	for _, rawFile := range filesRAW {
		jpgFile := strings.ReplaceAll(rawFile, fileExtRAW, fileExtJPG)
		if stringslice.Contains(filesJPG, jpgFile) {
			continue
		}

		deleteFiles = append(deleteFiles, rawFile)
		if !yes {
			log.Printf("Would delete file %v", rawFile)
		}
	}

	log.Printf("found %v files for deletion", len(deleteFiles))
	if !yes {
		log.Print("Dry run, exiting. (run with -yes to execute deletion)")
		return
	}

	for _, file := range deleteFiles {
		fp := filepath.Join(dirNameRAW, file)
		if err := os.Remove(fp); err != nil {
			log.Fatalf("failed to delete file at path %v: %v", fp, err)
		}
	}

	log.Print("Done.")
}

func listFiles(dirName, fileExt string) []string {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatalf("failed to read directory %v: %v", dirName, err)
	}

	fileNames := make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if strings.HasPrefix(file.Name(), ".") {
			continue
		}

		if !strings.EqualFold(strings.Trim(filepath.Ext(file.Name()), "."), fileExt) {
			continue
		}

		fileNames = append(fileNames, file.Name())
	}

	return fileNames
}
