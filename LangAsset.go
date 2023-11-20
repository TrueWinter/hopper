package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type LanguageAsset map[string]string

func filterLang(file os.File, patterns []string) LanguageAsset {
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal("Failed to get language file size", err)
	}

	data := make([]byte, fileInfo.Size())
	_, err = file.Read(data)
	if err != nil {
		log.Fatal("Failed to read language file", err)
	}

	langAsset := LanguageAsset{}
	err = json.Unmarshal(data, &langAsset)
	if err != nil {
		log.Fatal("Failed to parse language file", err)
	}

	langAssetMatches := LanguageAsset{}

	for k, v := range langAsset {
		for _, p := range patterns {
			if strings.Contains(k, p) {
				langAssetMatches[k] = v
			}
		}
	}

	return langAssetMatches
}