package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	renameFiles("*", "someFile.txt")
}


func renameFiles(input, newName string) {
	var f fs.WalkDirFunc

	reg, err := regexp.Compile(input)
	if err != nil {
		log.Println(err)
	}

	f = func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && reg.MatchString(d.Name()) {
			p := strings.Split(path, "/")
			newPath := strings.Join(p[:len(p)-1], "/")
			err := os.Rename(newPath+"/"+d.Name(), newPath+"/"+newName)
			if err != nil {
				log.Println(err)
			}
		}

		return nil
	}

	filepath.WalkDir("./sample/", f)
}