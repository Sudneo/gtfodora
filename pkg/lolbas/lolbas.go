package lolbas

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	cloner "github.com/sudneo/gtfodora/pkg/repo_utils"
	"gopkg.in/yaml.v2"
)

const (
	repoURL string = "https://github.com/LOLBAS-Project/LOLBAS"
)

type LOLbasbin struct {
	Name        string        `yaml:"Name"`
	Description string        `yaml:"Description"`
	Author      interface{}   `yaml:"Author"`
	Created     string        `yaml:"Created"`
	Commands    []CommandSpec `yaml:"Commands"`
	FullPath    []struct {
		Path string `yaml:"Path"`
	} `yaml:"Full_Path"`
	CodeSample []struct {
		Code string `yaml:"Code"`
	} `yaml:"Code_Sample"`
	Detection []struct {
		IOC interface{} `yaml:"IOC"`
	} `yaml:"Detection"`
	Resources []struct {
		Link string `yaml:"Link"`
	} `yaml:"Resources"`
	Acknowledgement []struct {
		Person string `yaml:"Person"`
		Handle string `yaml:"Handle"`
	} `yaml:"Acknowledgement"`
}

type CommandSpec struct {
	Command         string `yaml:"Command"`
	Description     string `yaml:"Description"`
	UseCase         string `yaml:"UseCase"`
	Category        string `yaml:"Category"`
	Privileges      string `yaml:"Privileges"`
	MitreID         string `yaml:"MitreID"`
	MItreLink       string `yaml:"MItreLink"`
	OperatingSystem string `yaml:"OperatingSystem"`
}

type Spec struct {
	Description string
	Code        string
}

func CloneLOLbas(path string) {
	err := cloner.Clone_repo(repoURL, path)
	if err != nil {
		log.Warn("Failed to clone LOLbas repository, results will be partial")
	}
}

func pull(path string) {
	err := cloner.Pull_repo(path)
	if err != nil {
		log.Warn("Failed to pull the LOLbas repository, results might be outdated.")
	}
}

func Parse(filePath string) LOLbasbin {

	yamlFile, err := ioutil.ReadFile(filePath)
	var bin LOLbasbin
	if err != nil {
		fmt.Println("Error parsing file")
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &bin)
	return bin
}

func ParseAll(path string) []LOLbasbin {
	cloner.Pull_repo(path)
	binary_path := path + "/yml/"
	var files []string
	var parsedFiles []LOLbasbin
	err := filepath.Walk(binary_path, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.WithFields(log.Fields{
			"Path": path,
		}).Error("Failed to walk the specified path")
		return parsedFiles
	}
	for _, file := range files {
		if info, err := os.Stat(file); err == nil && !info.IsDir() {
			if filepath.Ext(file) == ".yml" {
				f := Parse(file)
				parsedFiles = append(parsedFiles, f)
			}
		}
	}
	return parsedFiles
}

func (f *LOLbasbin) LOLbasHasFunction(a string) bool {
	for _, cmd := range f.Commands {
		if cmd.Category == a {
			return true
		}
	}
	return false
}

func (f *LOLbasbin) LOLbasGetFunctionDetails(a string) CommandSpec {
	for _, cmd := range f.Commands {
		if cmd.Category == a {
			return cmd
		}
	}
	return CommandSpec{}
}

func (f *LOLbasbin) LOLbasPrettyPrint() {
	fmt.Printf("Information about: %v\n", f.Name)
	fmt.Printf("Description: %v\n", f.Description)
	for _, cmd := range f.Commands {
		fmt.Printf("--------------------------------\n")
		cmd.CmdPrettyPrint()
	}

}

func (c *CommandSpec) CmdPrettyPrint() {
	fmt.Printf("%v:\n", c.Category)
	if len(c.Description) > 0 {
		fmt.Printf("- Description:\n")
		fmt.Printf("%v\n", c.Description)
	}
	if len(c.Command) > 0 {
		fmt.Printf("- Code:\n")
		fmt.Printf("%s\n", c.Command)
	}
	if len(c.Description) > 0 || len(c.Command) > 0 {
		fmt.Printf("\n")
	}
}
