package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type SoundAsset map[string]SoundAssetEvent
type SoundAssetEvent struct {
	Replace bool `json:"replace,omitempty"`
	Subtitle string `json:"subtitle,omitempty"`
	Sounds interface{} `json:"sounds,omitempty"`
}

func filterSound(file os.File, patterns []string, replace bool) SoundAsset {
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal("Failed to get sound file size", err)
	}

	data := make([]byte, fileInfo.Size())
	_, err = file.Read(data)
	if err != nil {
		log.Fatal("Failed to read sound file", err)
	}

	soundAsset := SoundAsset{}
	err = json.Unmarshal(data, &soundAsset)
	if err != nil {
		log.Fatal("Failed to parse sound file", err)
	}

	soundAssetMatches := SoundAsset{}

	for k, v := range soundAsset {
		for _, p := range patterns {
			if strings.Contains(k, p) {
				if replace {
					v.Replace = true
				}

				soundAssetMatches[k] = v
			}
		}
	}

	return soundAssetMatches
}