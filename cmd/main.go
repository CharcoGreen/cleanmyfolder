package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Files []File

type File struct {
	Name string
	Date string
}

// Directory to clean
func dirToClean(path string) string {

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Errorf("Error:  %s", err)
	}

	dirToClean := home + "/" + path

	return dirToClean
}

// Check How old is the file
func isMoreOld(date time.Time, days int64) bool {

	startTime := time.Now()

	startTime = startTime.Add(time.Duration(days) * 24 * time.Hour * -1)

	if date.Format("01-02-2006") < startTime.Format("01-02-2006") {

		// Remove files
		return true
	}

	return false
}

// Add Date to Slices
func (f *Files) Add(name, date string) {

	var t File

	t.Date = date
	t.Name = name

	*f = append(*f, t)
}

// Delete files in slice With datetime more old than days by parameter
func (f *Files) Delete(days int64) {

	for _, item := range *f {

		layout := "01-02-2006"
		date, err := time.Parse(layout, item.Date)

		if err != nil {
			fmt.Errorf("Print: ", err)
		}

		if isMoreOld(date, days) {

			fmt.Println("We go to remove this File: " + item.Name + " " + item.Date)
		}
	}
}

// Get all Files and return valid Slice
func getAllFiles(startPath string) Files {

	dataFile := Files{}

	filepath.Walk(startPath, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() {
			dataFile.Add(path, info.ModTime().Format("01-02-2006"))
		}

		if err != nil {
			fmt.Println("ERROR:", err)
		}

		return nil
	})

	return dataFile
}

// Main
func main() {

	folder := flag.String("folder", dirToClean("Downloads"), "Directory to cleanup Default is Downloads")
	days := flag.Int64("days", 30, "Number of Days from this to back they will be delete")

	flag.Parse()

	fmt.Println("Directory", *folder)
	fmt.Println("Days:", *days)

	// Check Folder is not existe

	check := getAllFiles(*folder)
	check.Delete(*days)

	fmt.Printf("Dir to clean is: %s \n", dirToClean("Download"))
}
