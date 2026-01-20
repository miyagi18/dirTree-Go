package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func getChildItem(out *bytes.Buffer, path string, prefix string, printFiles bool) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

		if !printFiles {
		for idx := len(files) - 1; idx >= 0; idx-- {
			if !files[idx].IsDir() {
				files = append(files[:idx], files[idx+1:]...)
			}
		}
	}
	
	for i := 0; i < len(files); i++ {
		isDir := files[i].IsDir()
		var size string

		if !isDir {
			if files[i].Size() > 0 {
				size = " (" + strconv.FormatInt(files[i].Size(), 10) + "b)"
			} else {
				size = " (empty)"
			}
		}

		if i < len(files)-1 {
			fmt.Fprintf(out, "%v├───%v%v\n", prefix, files[i].Name(), size)
		} else {
			fmt.Fprintf(out, "%v└───%v%v\n", prefix, files[i].Name(), size)
		}


		if isDir {
			nextDir := path + string(os.PathSeparator) + files[i].Name()
			var nextPrefix string
			if i < len(files) - 1 {
				nextPrefix = prefix + "|" + "\t"
			} else {
				nextPrefix = prefix + "\t"
			}
			err := getChildItem(out, nextDir, nextPrefix, printFiles)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func dirTree(out *bytes.Buffer, path string, printFiles bool) error {
	err := getChildItem(out, path, "", printFiles)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	out := new(bytes.Buffer)
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
