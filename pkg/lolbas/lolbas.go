package lolbas

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	cloner "github.com/sudneo/gtfodora/pkg/repo_utils"
	"gopkg.in/yaml.v2"
)

const (
	repoURL string = "https://github.com/LOLBAS-Project/LOLBAS"
)

type LOLbasbin struct {
	Name        string      `yaml:"Name"`
	Description string      `yaml:"Description"`
	Author      interface{} `yaml:"Author"`
	Created     string      `yaml:"Created"`
	Commands    []struct {
		Command         string `yaml:"Command"`
		Description     string `yaml:"Description"`
		UseCase         string `yaml:"UseCase"`
		Category        string `yaml:"Category"`
		Privileges      string `yaml:"Privileges"`
		MitreID         string `yaml:"MitreID"`
		MItreLink       string `yaml:"MItreLink"`
		OperatingSystem string `yaml:"OperatingSystem"`
	} `yaml:"Commands"`
	FullPath []struct {
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

func CloneLOLbas(path string) {
	cloner.Clone_repo(repoURL, path)
}

func pull(path string) {
	cloner.Pull_repo(path)
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
		panic(err)
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
