package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	ext  string    // extension to filter out
	size int64     // min file size
	list bool      // list files
	del  bool      // delete files
	wLog io.Writer // log destination writer
}

func run(root string, out io.Writer, cfg config) error {
	// Create a new instance of Logger
	delLogger := log.New(cfg.wLog, "DELETED FILE: ", log.LstdFlags)

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

			// Delete files
			if cfg.del {
				return delFile(path, delLogger)
			}

			// List is the default option if nothing else was set
			return listFile(path, out)
		})
}

func main() {
	// Parsing command line flags
	root := flag.String("root", ".", "Root directory to start")
	logFile := flag.String("log", "", "Log deletes to this file")
	// Action options
	list := flag.Bool("list", false, "List files only")
	del := flag.Bool("del", false, "Delete files")
	// Filter options
	ext := flag.String("ext", "", "File extension to filter out")
	size := flag.Int64("size", 0, "Minimum file size")
	flag.Parse()

	// Open log file for reading and writing
	// STDOUT is the default log file
	var (
		f   = os.Stdout
		err error
	)

	if *logFile != "" {
		f, err = os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
	}

	c := config{
		ext:  *ext,
		list: *list,
		size: *size,
		del:  *del,
		wLog: f,
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
