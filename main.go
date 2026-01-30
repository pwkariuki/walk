package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type config struct {
	ext  string // extension to filter out
	size int64  // min file size
	list bool   // list files
}

func run(root string, out io.Writer, cfg config) error {
	return filepath.Walk(root,
		func(path string, info fs.FileInfo, err error) error {
			if err != nil { // unable to walk to file/directory
				return err
			}

			if filterOut(path, cfg.ext, cfg.size, info) {
				return nil
			}

			// If list was explicitly set, no further processing
			if cfg.list {
				return listFile(path, out)
			}

			// List is the default option if nothing else was set
			return listFile(path, out)
		})
}

func main() {
	// Parsing command line flags
	root := flag.String("root", ".", "Root directory to start")
	list := flag.Bool("list", false, "List files only")
	// Filter options
	ext := flag.String("ext", "", "File extension to filter out")
	size := flag.Int64("size", 0, "Minimum file size")
	flag.Parse()

	c := config{
		ext:  *ext,
		list: *list,
		size: *size,
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
