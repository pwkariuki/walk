package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

// Checks if the given `path` has to be filtered out from
// the results according to the following conditions:
//
// 1. the path points to a directory
// 2. the file size is less than the minimum size provided by the user
// 3. the file extension does not math the extension provided by the user
func filterOut(path, ext string, minSize int64, info fs.FileInfo) bool {
	if info.IsDir() || info.Size() < minSize {
		return true
	}

	if ext != "" && filepath.Ext(path) != ext {
		return true
	}
	return false
}

// Print out the path of the current file to the specified `io.Writer`
func listFile(path string, out io.Writer) error {
	_, err := fmt.Fprintln(out, path)
	return err
}

// Delete the file at specified file path
// Caution: Do NOT run this function as a priviledged user
func delFile(path string, delLogger *log.Logger) error {
	if err := os.Remove(path); err != nil {
		return err
	}

	delLogger.Println(path)
	return nil
}
