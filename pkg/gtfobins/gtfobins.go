package gtfobins

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	cloner "github.com/sudneo/gtfodora/pkg/repo_utils"
	"gopkg.in/yaml.v2"
)

const (
	repoURL string = "https://github.com/GTFOBins/GTFOBins.github.io"
)

type FileInfo struct {
	Filename string
	Data     GTFObin
	Binary   string
}

type GTFObin struct {
	Functions struct {
		Shell                      []spec `yaml:"shell"`
		FileUpload                 []spec `yaml:"file-upload"`
		FileDownload               []spec `yaml:"file-download"`
		FileWrite                  []spec `yaml:"file-write"`
		FileRead                   []spec `yaml:"file-read"`
		LibraryLoad                []spec `yaml:"library-load"`
		Sudo                       []spec `yaml:"sudo"`
		NonInteractiveReverseShell []spec `yaml:"non-interactive-reverse-shell"`
		Command                    []spec `yaml:"command"`
		BindShell                  []spec `yaml:"bind-shell"`
		SUID                       []spec `yaml:"suid"`
		LimitedSUID                []spec `yaml:"limited-suid"`
		ReverseShell               []spec `yaml:"reverse-shell"`
		NonInteractiveBindShell    []spec `yaml:"non-interactive-bind-shell"`
		Capabilities               []spec `yaml:"capabilities"`
	} `yaml:"functions"`
}

type spec struct {
	Code        string `yaml:"code"`
	Description string `yaml:"description"`
}

func CloneGTFO(destination string) {
	cloner.Clone_repo(repoURL, destination)
}

func ParseAll(path string) []FileInfo {
	cloner.Pull_repo(path)
	binary_path := path + "/_gtfobins/"
	var files []string
	var parsedFiles []FileInfo
	err := filepath.Walk(binary_path, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if info, err := os.Stat(file); err == nil && !info.IsDir() {
			binaryName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
			f := FileInfo{file, parse(file), binaryName}
			parsedFiles = append(parsedFiles, f)
		}
	}
	return parsedFiles
}

func pull(path string) {
	cloner.Pull_repo(path)
}

func parse(filePath string) GTFObin {

	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err.Error())
	}
	var bin GTFObin
	err = yaml.Unmarshal(yamlFile, &bin)
	return bin

}
