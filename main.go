package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	dir    string
	dryRun bool
)

func main() {
	flag.StringVar(&dir, "dir", ".", "directory to process")
	flag.BoolVar(&dryRun, "d", false, "dry run (don't apply changes)")
	flag.Parse()

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			ext := filepath.Ext(path)
			lowerExt := strings.ToLower(ext)
			changed := false

			newPath := path

			if ext != lowerExt {
				newPath = strings.TrimSuffix(path, ext) + lowerExt
				changed = true
			}

			if lowerExt == ".jpeg" {
				newPath = strings.TrimSuffix(newPath, lowerExt) + ".jpg"
				changed = true
			}

			if changed {
				fmt.Printf("Rename: %s -> %s\n", path, newPath)
				if !dryRun {
					err := os.Rename(path, newPath)
					if err != nil {
						return err
					}
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
