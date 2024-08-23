package main

import (
	"log"
	"os/exec"
	"strings"
)

func main() {
	SaveInPlace("updateing...", []string{"."})
}
func SaveInPlace(comment string, files []string) {
	GitAdd(files)
	GitCommit(comment)
}

func GitAdd(files []string) error {
	var fileList string
	if len(files) == 0 {
		fileList = "."
	} else {
		fileList = strings.Join(files, " ")
	}
	_cmd := "git add " + fileList
	// + " && git commit -m " + comment
	chunks := strings.Split(_cmd, " ")
	cmd := exec.Command(chunks[0], chunks[1:]...)

	if err := cmd.Run(); err != nil {
		// fmt.Println("Error: ", err)
		return err
	}
	var out strings.Builder
	cmd.Stdout = &out

	log.Println("CMD =>", cmd)

	return nil
}

func GitCommit(comment string) error {
	_cmd := "git commit -m " + comment
	chunks := strings.Split(_cmd, " ")
	cmd := exec.Command(chunks[0], chunks[1:]...)

	if err := cmd.Run(); err != nil {
		// fmt.Println("Error: ", err)
		return err
	}
	var out strings.Builder
	cmd.Stdout = &out

	log.Println("CMD =>", cmd)

	return nil

}
