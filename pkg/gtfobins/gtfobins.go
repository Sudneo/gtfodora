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
		Shell                      []Spec `yaml:"shell"`
		FileUpload                 []Spec `yaml:"file-upload"`
		FileDownload               []Spec `yaml:"file-download"`
		FileWrite                  []Spec `yaml:"file-write"`
		FileRead                   []Spec `yaml:"file-read"`
		LibraryLoad                []Spec `yaml:"library-load"`
		Sudo                       []Spec `yaml:"sudo"`
		NonInteractiveReverseShell []Spec `yaml:"non-interactive-reverse-shell"`
		Command                    []Spec `yaml:"command"`
		BindShell                  []Spec `yaml:"bind-shell"`
		SUID                       []Spec `yaml:"suid"`
		LimitedSUID                []Spec `yaml:"limited-suid"`
		ReverseShell               []Spec `yaml:"reverse-shell"`
		NonInteractiveBindShell    []Spec `yaml:"non-interactive-bind-shell"`
		Capabilities               []Spec `yaml:"capabilities"`
	} `yaml:"functions"`
}

type Spec struct {
	Code        string `yaml:"code"`
	Description string `yaml:"description"`
}

func Clone(destination string) {
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
		binaryName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
		f := FileInfo{file, Parse(file), binaryName}
		parsedFiles = append(parsedFiles, f)
	}
	return parsedFiles
}

func Pull(path string) {
	cloner.Pull_repo(path)
}

func Parse(filePath string) GTFObin {

	yamlFile, err := ioutil.ReadFile(filePath)
	fmt.Println(string(yamlFile))
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err.Error())
	}
	var bin GTFObin
	err = yaml.Unmarshal(yamlFile, &bin)
	return bin

}
