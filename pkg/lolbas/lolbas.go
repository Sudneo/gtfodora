package lolbas

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/sudneo/gtfodora/pkg/binary"
	cloner "github.com/sudneo/gtfodora/pkg/repo_utils"
	"gopkg.in/yaml.v2"
)

const (
	repoURL string = "https://github.com/LOLBAS-Project/LOLBAS"
)

type lolbasbin struct {
	Name        string        `yaml:"Name"`
	Description string        `yaml:"Description"`
	Author      interface{}   `yaml:"Author"`
	Created     string        `yaml:"Created"`
	Commands    []commandSpec `yaml:"Commands"`
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

type commandSpec struct {
	Command         string `yaml:"Command"`
	Description     string `yaml:"Description"`
	UseCase         string `yaml:"UseCase"`
	Category        string `yaml:"Category"`
	Privileges      string `yaml:"Privileges"`
	MitreID         string `yaml:"MitreID"`
	MItreLink       string `yaml:"MItreLink"`
	OperatingSystem string `yaml:"OperatingSystem"`
}

type spec struct {
	Description string
	Code        string
}

func Clone(path string) {
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

func Parse(filePath string) lolbasbin {

	yamlFile, err := ioutil.ReadFile(filePath)
	var bin lolbasbin
	if err != nil {
		fmt.Println("Error parsing file")
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &bin)
	return bin
}

func (l *lolbasbin) transform() binary.Binary {
	var bin binary.Binary
	bin.Name = l.Name
	bin.Type = "win"
	for _, c := range l.Commands {
		var cmd binary.Command
		// Check current bin commands to see if there is already one with same
		// category and in case append to it
		existing := false
		for i, _ := range bin.Commands {
			if bin.Commands[i].Function == strings.ToLower(c.Category) {
				det := bin.Commands[i].Details
				bin.Commands[i].Details = append(det, binary.FunctionSpec{
					Description: c.Description,
					Code:        c.Command})
				existing = true
			}
		}
		if !existing {
			switch strings.ToLower(c.Category) {
			case "execute":
				cmd.Function = "command"
			case "awl Bypass":
				cmd.Function = "awlbypass"
			case "uac bypass":
				cmd.Function = "uacbypass"
			default:
				cmd.Function = strings.ToLower(c.Category)
			}
			cmd.Details = append(cmd.Details, binary.FunctionSpec{
				Description: c.Description,
				Code:        c.Command})
			bin.Commands = append(bin.Commands, cmd)
		}
	}
	return bin

}

func ParseAll(path string) []binary.Binary {
	cloner.Pull_repo(path)
	binary_path := path + "/yml/"
	var files []string
	var parsedFiles []binary.Binary
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
				parsedFiles = append(parsedFiles, f.transform())
			}
		}
	}
	return parsedFiles
}
