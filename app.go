package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide an argument")
		return
	}

	file := args[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)
		fileInfo, err := os.Stat(fullPath)

		if err != nil {
			fmt.Println("Error searching for this term:")
			fmt.Println(err)
			continue
		}

		mode := fileInfo.Mode()
		if mode.IsRegular() {
			if mode&011 != 0 {
				fmt.Println(fullPath)
				return
			}
		}
	}
}
