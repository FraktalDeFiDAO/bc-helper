package versioning

import (
	"flag"
	"fmt"
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
	versionShow        bool
)

func Versioning() {
	versionFile = "./version"

	versionUpdateTypes = append(versionUpdateTypes, "minor")
	versionUpdateTypes = append(versionUpdateTypes, "major")
	versionUpdateTypes = append(versionUpdateTypes, "release")

	flag.StringVar(&versionUpdateType, "update-type", "", "Type of update (minor, major or release)")
	flag.BoolVar(&versionShow, "version", false, "Type of update (minor, major or release)")
	flag.Parse()
	// log.Println(len(os.Args), os.Args, versionUpdateType)
	if versionShow {
		ShowVersion()

	}
	if _, err := os.Stat(versionFile); err != nil {
		if err := os.WriteFile(versionFile, []byte("0.0.1"), 0644); err != nil {
			log.Fatal("Error writing version file =>", err)
		}
	}
	versionBytes, err := os.ReadFile(versionFile)
	if err != nil {
		log.Fatalln("Error reading version file =>", err)
	}
	version = string(versionBytes)
	// versionChunks := strings.Split(version, ".")

	// log.Println("versionChunks =>", versionChunks, version)

	if versionUpdateType == versionUpdateTypes[0] {
		IncMinor()
	} else if versionUpdateType == versionUpdateTypes[1] {
		IncMajor()
	} else if versionUpdateType == versionUpdateTypes[2] {
		IncRelease()
	}

	if err := os.WriteFile(versionFile, []byte(version), 0644); err != nil {
		log.Fatalln("Error writing version file =>", err)
	}
}

func IncMinor() {
	log.Println("UPDATE MINOR", versionUpdateType)
	versionChunks := strings.Split(version, ".")

	u, err := strconv.Atoi(versionChunks[2])
	if err != nil {
		log.Fatal(err)
	}
	versionChunks[2] = strconv.Itoa(u + 1)

	version = strings.Join(versionChunks, ".")
}

func IncMajor() {
	log.Println("UPDATE MAJOR", versionUpdateType)

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
	log.Println("UPDATE RELEASE", versionUpdateType)

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

func ShowVersion() {
	v, err := os.ReadFile(versionFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Version =>", "v"+string(v))
}
