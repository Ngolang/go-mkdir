package main

import (
	"errors"
	"flag"
	"io/fs"
	"log"
	"os"
)

func main() {
	folder := flag.String("folder", "temp", "the name of the folder you want to created")
	permission := flag.Int("permission", 0755, "permission of the folder")
	flag.Parse()

	if err := mkdir(*folder, *permission); err != nil {
		log.Println("error:", err)
	}
}

func mkdir(folder string, permission int) error {
	if os.IsExist(os.Mkdir(folder, fs.FileMode(permission))) {
		return errors.New("file exist")
	}

	return nil
}
