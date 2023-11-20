package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func sep() string {
	if runtime.GOOS == "windows" {
		return "\\"
	}

	return "/"
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil;
}

func createFile(name string) *os.File {
	err := os.MkdirAll(filepath.Dir(*outputDir + "/" + name), 0644)
	if err != nil {
		log.Fatal("An error occurred while ensuring all directories exist", err)
	}

	file, err := os.Create(*outputDir + "/" + name)
	if err != nil {
		log.Fatal("Failed to create file", err)
	}

	return file
}

// https://stackoverflow.com/a/30708914
func dirIsEmpty(name string) bool {
    f, err := os.Open(name)
    if err != nil {
        return false
    }
    defer f.Close()

    _, err = f.Readdirnames(1)
    return err == io.EOF
}

func copy(name string, src io.Reader) {
	file := createFile(name)
	defer file.Close()

	_, copyErr := io.Copy(file, src)
	if copyErr != nil {
		log.Fatal("Failed to copy file", copyErr)
	}

	log.Println("Copied " + name)
}

func writeJson(name string, j interface{}) {
	file := createFile(name)
	defer file.Close()

	out, err := json.Marshal(j)
	if err != nil {
		log.Fatal("Failed to stringify JSON", err)
	}
	outFormatted := bytes.Buffer{}
	err = json.Indent(&outFormatted, out, "", "    ")
	if err != nil {
		log.Fatal("Failed to format JSON", err)
	}

	_, err = file.Write(outFormatted.Bytes())
	if err != nil {
		log.Fatal("Failed to write JSON", err)
	}

	log.Println("Copied " + name)
}