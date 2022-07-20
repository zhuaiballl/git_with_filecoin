package util

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func IsGitRepo() bool {
	//_, err := os.Stat(".git")
	//return Exist(err)
	resp, err := exec.Command("git", "rev-parse", "--is-inside-work-tree").Output()
	if err != nil {
		//log.Panic(err)
	}
	return strings.HasPrefix(string(resp), "true")
}

func PathExist(loc string) bool {
	_, err := os.Stat(loc)
	return Exist(err)
}

// Exist uses err returned from os.Stat to determine if a file/folder exists
func Exist(err error) bool {
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			log.Panic(err)
		}
	}
	return true
}
