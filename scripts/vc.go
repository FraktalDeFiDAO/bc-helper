package main

import (
	"errors"
	"log"
	"os/exec"
)

func main() {

}
func UpdateTag(tag string) error {
	AddAndCommit("updateing...", []string{"."})
	err := GitPush("origin", "master")
	if err != nil {
		return errors.New("Error updateing tag...\n" + err.Error())
	}
	GitTag(tag)
	GitPushTag()
	return nil
}
func AddAndCommit(comment string, files []string) error {
	err := GitAdd(files)
	if err != nil {
		return errors.New("ERROR =>" + err.Error())
	}
	GitCommit(comment)
	if err != nil {
		return errors.New("ERROR (GitCommit) =>" + err.Error())
	}
	return nil
}

func GitAdd(files []string) error {
	log.Println("Git => add")
	cmd := []string{
		"add",
	}
	cmd = append(cmd, files...)
	out, err := Git(cmd...)

	if err != nil {
		return errors.New("GIT ADD =>" + err.Error())
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

	out, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	return string(out), nil
}

func GitPush(target string, branch string) error {
	out, err := Git("push", target, branch)
	if err != nil {
		return errors.New("PUSH =>" + err.Error())
	}
	log.Println("Git Push =>", target, branch)

	log.Println(string(out))
	return nil
}

func GitTag(tag string) error {
	out, err := Git("tag", tag)
	if err != nil {
		return errors.New("Git Tag Error =>" + err.Error())
	}
	log.Println("Git Tag =>", tag)

	log.Println(string(out))
	return nil

}

func GitPushTag() error {
	out, err := Git("push", "--tags")
	if err != nil {
		return errors.New("Git Tag Error =>" + err.Error())
	}
	log.Println("==> Git Tags <==")

	log.Println(string(out))
	return nil

}
