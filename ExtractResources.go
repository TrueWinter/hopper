package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

func parsePattern(pattern string) []string {
	list := strings.Split(pattern, ",")

	if len(list) == 0 {
		log.Fatal("Empty pattern list")
	}

	tmpList := []string{}

	for _, v := range list {
		if v != "" {
			tmpList = append(tmpList, v)
		}
	}

	if len(tmpList) == 0 {
		log.Fatal("Invalid pattern: empty list")
	}

	return tmpList
}

var outputDir *string

func main() {
	appDataDir, appDataError := os.UserConfigDir()
	if appDataError != nil {
		panic(appDataError)
	}

	index := flag.String("index", "", "Asset index version")
	version := flag.String("version", "", "Minecraft version")
	mcDir := flag.String("mcdir", appDataDir + sep() + ".minecraft", "Minecraft directory")
	patternTmp := flag.String("pattern", "", "Comma separated list of resource names to match")
	noExtract := flag.Bool("noextract", false, "Skip extracting the JAR file, ignoring all resources it contains")
	noIndex := flag.Bool("noindex", false, "Skip using the asset index, ignoring all resources it contains")
	outputDir = flag.String("output", "", "Output directory, must be empty")
	flag.Parse()

	if *index == "" && !*noIndex {
		flag.Usage()
		log.Fatal("Asset index version is required unless noindex is true")
	}	

	if *version == "" && !*noExtract {
		flag.Usage()
		log.Fatal("Minecraft version is required unless noextract is true")
	}

	pattern := parsePattern(*patternTmp)

	if *outputDir == "" {
		flag.Usage()
		log.Fatal("Output directory is required")
	}

	if !exists(*outputDir) {
		log.Fatal("Output directory does not exist")
	}

	if !dirIsEmpty(*outputDir) {
		log.Fatal("Output directory has contents")
	}

	if *noExtract && *noIndex {
		flag.Usage()
		log.Fatal("Both noextract and noindex are true")
	}

	if !*noIndex {
		log.Println("Using index:", *index)
	}

	if !*noExtract {
		log.Println("Using version:", *version)
	}

	log.Println("Minecraft directory:", *mcDir)
	log.Println("Patterns:", strings.Join(pattern, ", "))
	log.Println("Output directory:", *outputDir)

	if !*noIndex {
		assetIndex := getAssetIndex(*mcDir, *index)

		for n, v := range assetIndex.Objects {
			for _, p := range pattern {
				if strings.Contains(n, p) {
					asset := getAssetObject(*mcDir, v.Hash)
					defer asset.Close()
					copy(n, asset)
				}
			}
		}
	}

	if !*noExtract {
		getAssetsFromJarFile(*mcDir, *version, pattern)
	}
}