package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type ClientJson struct {
	AssetIndex string `json:"assets"`
}

type AssetIndex struct {
	Objects map[string]AssetIndexObject `json:"objects"`
}

type AssetIndexObject struct {
	Hash string `json:"hash"`
	Size int    `json:"size"`
}

func getAssetIndexFromVersion(mcDir string, version string) AssetIndex {
	file, err := os.Open(mcDir + sep() + "versions" + sep() + version + sep() + version + ".json")
	if err != nil {
		panic(err)
	}

	clientJson := ClientJson{}
	reader, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(reader, &clientJson)
	if err != nil {
		panic(err)
	}

	log.Println("Using index: " + clientJson.AssetIndex)

	return getAssetIndex(mcDir, clientJson.AssetIndex)
}

func getAssetIndex(mcDir string, index string) AssetIndex {
	file, err := os.Open(mcDir + sep() + "assets" +
		sep() + "indexes" + sep() + index + ".json")
	if err != nil {
		panic(err)
	}

	assetIndex := AssetIndex{}
	reader, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(reader, &assetIndex)
	if err != nil {
		panic(err)
	}

	return assetIndex
}

func getAssetObject(mcDir string, hash string) *os.File {
	file, err := os.Open(mcDir + sep() + "assets" + sep() +
		"objects" + sep() + hash[:2] + sep() + hash)
	if err != nil {
		panic(err)
	}

	return file
}