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

func parseLang(lang string) []string {
	list := strings.Split(lang, ",")
	tmpList := []string{}

	for _, v := range list {
		if v != "" {
			tmpList = append(tmpList, v)
		}
	}

	return tmpList
}

func parseExclude(exclude string) []string {
	list := strings.Split(exclude, ",")
	tmpList := []string{}

	for _, v := range list {
		if v != "" {
			tmpList = append(tmpList, v)
		}
	}

	return tmpList
}

func isExcluded(name string, exclusions []string) bool {
	for _, e := range exclusions {
		if strings.Contains(name, e) {
			return true
		}
	}

	return false
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
	langTmp := flag.String("lang", "", "Comma separated list of language files name patterns to extract strings from")
	sounds := flag.Bool("sounds", false, "Copy filtered sounds.json file")
	replaceSounds := flag.Bool("replacesounds", false, "Set the replace property in sounds.json to true for all sounds")
	excludeTmp := flag.String("exclude", "", "Comma separated list of resource name patterns to exclude")
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
	lang := parseLang(*langTmp)
	for _, v := range lang {
		pattern = append(pattern, "/lang/" + v)
	}

	exclude := parseExclude(*excludeTmp)

	if *replaceSounds && !*sounds {
		flag.Usage()
		log.Fatal("The replacesounds flag cannot be used without the sounds flag")
	}

	if *sounds {
		pattern = append(pattern, "minecraft/sounds.json")
	}

	if (*index == "" || *noIndex) && len(lang) > 0 {
		flag.Usage()
		log.Fatal("Asset index is required when extracting from language files")
	}

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

	*outputDir = *outputDir + sep() + "assets"

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

	if len(lang) > 0 {
		log.Println("Using language name pattern:", strings.Join(lang, ","))
	}

	log.Println("Minecraft directory:", *mcDir)
	log.Println("Patterns:", strings.Join(pattern, ", "))
	log.Println("Output directory:", *outputDir)

	if !*noIndex {
		assetIndex := getAssetIndex(*mcDir, *index)

		for n, v := range assetIndex.Objects {
			for _, p := range pattern {
				if strings.Contains(n, p) && !isExcluded(n, exclude) {
					asset := getAssetObject(*mcDir, v.Hash)
					defer asset.Close()

					if strings.Contains(n, "/lang/") ||
						n == "minecraft/sounds.json" {
							if n == "minecraft/sounds.json" {
								filtered := filterSound(*asset, pattern, *replaceSounds)
								writeJson(n, filtered)
							}
							
							if strings.Contains(n, "/lang/") {
								if len(lang) > 0 {
									filtered := filterLang(*asset, pattern)
									writeJson(n, filtered)
								}

							}
							
							continue
						}

					copy(n, asset)
				}
			}
		}
	}

	if !*noExtract {
		getAssetsFromJarFile(*mcDir, *version, pattern, exclude)
	}
}