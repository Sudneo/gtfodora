package cloner

import (
	"github.com/go-git/go-git/v5"
	log "github.com/sirupsen/logrus"
)

func Clone_repo(repoURL string, destination string) error {
	_, err := git.PlainClone(destination, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: nil,
	})
	if err != nil {
		if err.Error() == "repository already exists" {
			log.Debugf("Repository %q already exists.", repoURL)
			return nil
		} else {
			log.WithFields(log.Fields{
				"Error":      err.Error(),
				"Repository": repoURL,
			}).Error("Failed to clone repository")
			return err
		}
	}
	log.Debugf("Repository %q successfully cloned into %q", repoURL, destination)
	return nil

}

func Pull_repo(path string) error {
	r, err := git.PlainOpen(path)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	w, err := r.Worktree()
	if err != nil {
		log.Error(err.Error())
		return err
	}

	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		if err.Error() == "already up-to-date" {
			log.Debugf("Repository %v is already updated, nothing to do.", path)
			return nil
		} else {
			log.WithFields(log.Fields{
				"Error":      err.Error(),
				"Repository": path,
			}).Error("Failed to pull repository")
			return err
		}
	}
	log.Debugf("Successfully pulled repository at %v", path)
	return nil
}
