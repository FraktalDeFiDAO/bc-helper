package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

type VersionUpdateType uint8

const (
	VU_Minor VersionUpdateType = iota
	VU_Major
	VU_Release
)

var (
	version            string
	versionUpdateType  string
	versionUpdateTypes []string
	versionFile        string
)

func main() {
	versionUpdateTypes[0] = "minor"
	versionUpdateTypes[1] = "major"
	versionUpdateTypes[2] = "release"

	flag.StringVar(&versionUpdateType, "update-type", "minor", "Type of update (minor, major or release)")

	versionChunks := strings.Split(version, ".")

	if versionUpdateType == versionUpdateTypes[0] {
		u, err := strconv.Atoi(versionChunks[2])
		if err != nil {
			log.Fatal(err)
		}
		versionChunks[2] = strconv.Itoa(u + 1)
	}
}

func IncMinor() {
	versionChunks := strings.Split(version, ".")

	u, err := strconv.Atoi(versionChunks[2])
	if err != nil {
		log.Fatal(err)
	}
	versionChunks[2] = strconv.Itoa(u + 1)

	version = strings.Join(versionChunks, ".")
}

func IncMajor() {
	versionChunks := strings.Split(version, ".")

	u, err := strconv.Atoi(versionChunks[1])
	versionChunks[1] = strconv.Itoa(u + 1)

	if err != nil {
		log.Fatal(err)
	}
	versionChunks[2] = strconv.Itoa(0)

	version = strings.Join(versionChunks, ".")

}

func IncRelease() {
	versionChunks := strings.Split(version, ".")

	u, err := strconv.Atoi(versionChunks[0])
	versionChunks[0] = strconv.Itoa(u + 1)

	if err != nil {
		log.Fatal(err)
	}
	versionChunks[1] = strconv.Itoa(0)
	versionChunks[2] = strconv.Itoa(0)

	version = strings.Join(versionChunks, ".")

}

func GetVersion() {
	v, err := os.ReadFile(versionFile)
	if err != nil {
		log.Fatal(err)
	}

}
