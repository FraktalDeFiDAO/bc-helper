package main

import (
	"errors"
	"log"
	"os/exec"
)

func main() {
	AddAndCommit("updateing...", []string{"."})
	err := GitPush("origin", "master")
	if err != nil {
		log.Fatal("ERROR:", err)
	}

}
func AddAndCommit(comment string, files []string) {
	err := GitAdd(files)
	if err != nil {
		log.Fatal("ERROR:", err)
	}
	GitCommit(comment)
	if err != nil {
		log.Fatal("ERROR (GitCommit):", err)
	}
}

func GitAdd(files []string) error {

	cmd := []string{
		"add",
	}
	cmd = append(cmd, files...)
	out, err := Git(cmd...)

	if err != nil {
		return errors.New("ADD =>" + err.Error())
	}

	log.Println(string(out))

	return nil
}

func GitCommit(comment string) error {

	out, err := Git("commit", "-m", comment)

	if err != nil {
		return errors.New("COMMIT =>" + err.Error())
	}
	log.Println("Git Commit =>", comment)

	log.Println(string(out))

	return nil

}

func Git(args ...string) (string, error) {
	cmd := exec.Command("git", args...)

	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
}

func GitPush(target string, branch string) error {
	out, err := Git(target, branch)
	if err != nil {
		return errors.New("PUSH =>" + err.Error())
	}
	log.Println("push", out)
	return nil
}
