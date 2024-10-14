// This is the first program
// written solely for my personal use.
// It creates folders that are present in the target location.
package main

import (
	"io/fs"
	"log"
	"os"
)

func main() {
	// gets target address from first argument
	from := os.Args[1]
	// gets the filesysytem of the location
	fys := os.DirFS(from)
	// reads all
	dirs, err := fs.ReadDir(fys, ".")
	if err != nil {
		//permission problem
	}
	for _, dir := range dirs {
		//if its a dir then writes it
		if dir.IsDir() {
			if err := os.Mkdir(dir.Name(), os.ModePerm); err != nil {
				log.Fatal(err)
			}

		}
	}

}
