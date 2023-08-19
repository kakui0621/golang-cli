package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: enhanced_du <directory>")
		return
	}

	startDir := os.Args[1]
	totalSize := int64(0)

	// ディレクトリを再帰的に走査
	err := filepath.Walk(startDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %v: %v\n", path, err)
			return err
		}

		// もしディレクトリではない場合、ファイルのサイズを合計に追加
		if !info.IsDir() {
			totalSize += info.Size()
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory %v: %v\n", startDir, err)
		return
	}

	fmt.Printf("Total size of %v: %v bytes\n", startDir, totalSize)
}
