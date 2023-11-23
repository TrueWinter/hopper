package main

import (
	"archive/zip"
	"log"
	"os"
	"strings"
)

func getAssetsFromJarFile(mcDir string, version string, pattern []string, exclude []string) {
	zipFile, err := os.Open(mcDir + sep() + "versions" + sep() +
		version + sep() + version + ".jar")
	if err != nil {
		log.Fatal("An error occurred while opening JAR file", err)
	}
	defer zipFile.Close()

	info, err := zipFile.Stat()
	if err != nil {
		log.Fatal("Failed to get JAR file size", err)
	}

	file, err := zip.NewReader(zipFile, info.Size())
	if err != nil {
		log.Fatal("An error occurred while reading JAR file", err)
	}

	for _, f := range file.File {
		if !strings.HasPrefix(f.Name, "assets/") {
			continue
		}

		// Language files are stored as hashed assets
		if strings.Contains(f.Name, "/lang/") {
			continue
		}

		if isExcluded(f.Name, exclude) {
			continue
		}

		for _, p := range pattern {
			if !strings.Contains(f.Name, p) {
				continue
			}

			r, err := f.Open()
			if err != nil {
				log.Fatal("An error occurred while opening asset from JAR file", err)
			}
			defer r.Close()

			copy(strings.Replace(f.Name, "assets/", "", 1), r)
		}
	}
}