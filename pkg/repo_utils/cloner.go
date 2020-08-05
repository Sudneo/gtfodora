package cloner

import (
	"fmt"
	"os"

	"github.com/go-git/go-git"
)

func Clone_repo(repoURL string, destination string) {
	_, err := git.PlainClone(destination, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	})
	if err == nil {
		fmt.Println("Successfully cloned repository.")
	} else {
		if err.Error() == "repository already exists" {
			fmt.Println("Repository exists already, nothing to do.")
		} else {
			_ = fmt.Sprintf("Failed to clone %v.", repoURL)
			fmt.Println(err)
		}
	}
}

func Pull_repo(path string) {
	r, _ := git.PlainOpen(path)
	w, err := r.Worktree()
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err == nil {
		fmt.Println("Successfully pulled repository.")
	} else {
		if err.Error() == "already up-to-date" {
			fmt.Println("Repository is already updated, nothing to do.")
		} else {
			fmt.Println("Failed to pull repository.")
			fmt.Println(err)
		}
	}
}
