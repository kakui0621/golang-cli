package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func hashFile(filePath string) ([]byte, error) {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	hash := md5.Sum(fileBytes)
	return hash[:], nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: filechecker <directory>")
		return
	}

	dirPath := os.Args[1]
	hashes := make(map[string][]string)

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			hash, err := hashFile(path)
			if err != nil {
				return err
			}
			hashStr := fmt.Sprintf("%x", hash)
			hashes[hashStr] = append(hashes[hashStr], path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for hash, files := range hashes {
		if len(files) > 1 {
			fmt.Println("Duplicate files for hash:", hash)
			for _, file := range files {
				fmt.Println("  -", file)
			}
		}
	}
}
