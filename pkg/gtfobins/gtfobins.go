package gtfobins

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	binary "github.com/sudneo/gtfodora/pkg/binary"
	cloner "github.com/sudneo/gtfodora/pkg/repo_utils"
	"gopkg.in/yaml.v2"
)

const (
	repoURL string = "https://github.com/GTFOBins/GTFOBins.github.io"
)

type FileInfo struct {
	Filename string
	Data     gtfobin
	Binary   string
}

type gtfobin struct {
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

func Clone(destination string) {
	err := cloner.Clone_repo(repoURL, destination)
	if err != nil {
		log.Warn("Cloning GTFObins failed, the results will be partial")
	}
}

func pull(path string) {
	err := cloner.Pull_repo(path)
	if err != nil {
		log.Warn("Failed to pull GTFObins repo, the results might be outdated.")
	}
}
func ParseAll(path string) []binary.Binary {
	cloner.Pull_repo(path)
	binary_path := path + "/_gtfobins/"
	var files []string
	var parsedFiles []binary.Binary
	err := filepath.Walk(binary_path, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if info, err := os.Stat(file); err == nil && !info.IsDir() && filepath.Base(file) != ".dir-locals.el" {
			binaryName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
			f := FileInfo{file, parse(file), binaryName}
			b := f.transform()
			parsedFiles = append(parsedFiles, b)
		}
	}
	return parsedFiles
}

func (f *FileInfo) transform() binary.Binary {
	var bin binary.Binary
	bin.Name = f.Binary
	bin.Type = "unix"
	if f.Data.Functions.Shell != nil {
		var cmd binary.Command
		cmd.Function = "shell"
		for _, spec := range f.Data.Functions.Shell {
			cmd.Details = append(cmd.Details, binary.FunctionSpec{
				Description: spec.Description,
				Code:        spec.Code})
		}
		bin.Commands = append(bin.Commands, cmd)
	}
	if f.Data.Functions.FileUpload != nil {
		var cmd binary.Command
		cmd.Function = "upload"
		for _, spec := range f.Data.Functions.FileUpload {
			cmd.Details = append(cmd.Details, binary.FunctionSpec{
				Description: spec.Description,
				Code:        spec.Code})
		}
		bin.Commands = append(bin.Commands, cmd)
	}
	if f.Data.Functions.FileDownload != nil {
		var cmd binary.Command
		cmd.Function = "download"
		for _, spec := range f.Data.Functions.FileDownload {
			cmd.Details = append(cmd.Details, binary.FunctionSpec{
				Description: spec.Description,
				Code:        spec.Code})
		}
		bin.Commands = append(bin.Commands, cmd)
	}
	if f.Data.Functions.FileWrite != nil {
		var cmd binary.Command
		cmd.Function = "filewrite"
		for _, spec := range f.Data.Functions.FileWrite {
			cmd.Details = append(cmd.Details, binary.FunctionSpec{
				Description: spec.Description,
				Code:        spec.Code})
		}
		bin.Commands = append(bin.Commands, cmd)
	}
	if f.Data.Functions.FileRead != nil {
		var cmd binary.Command
		cmd.Function = "fileread"
		for _, spec := range f.Data.Functions.FileRead {
			cmd.Details = append(cmd.Details, binary.FunctionSpec{
				Description: spec.Description,
				Code:        spec.Code})
		}
		bin.Commands = append(bin.Commands, cmd)
	}
	if f.Data.Functions.LibraryLoad != nil {
		var cmd binary.Command
		cmd.Function = "libraryload"
		for _, spec := range f.Data.Functions.LibraryLoad {
			cmd.Details = append(cmd.Details, binary.FunctionSpec{
				Description: spec.Description,
				Code:        spec.Code})
		}
		bin.Commands = append(bin.Commands, cmd)
	}
	if f.Data.Functions.Sudo != nil {
		var cmd binary.Command
		cmd.Function = "sudo"
		for _, spec := range f.Data.Functions.Sudo {
			cmd.Details = append(cmd.Details, binary.FunctionSpec{
				Description: spec.Description,
				Code:        spec.Code})
		}
		bin.Commands = append(bin.Commands, cmd)
	}
	if f.Data.Functions.NonInteractiveReverseShell != nil {
		var cmd binary.Command
		cmd.Function = "noninteractiverevshell"
		for _, spec := range f.Data.Functions.NonInteractiveReverseShell {
			cmd.Details = append(cmd.Details, binary.FunctionSpec{
				Description: spec.Description,
				Code:        spec.Code})
		}
		bin.Commands = append(bin.Commands, cmd)
	}
	if f.Data.Functions.Command != nil {
		var cmd binary.Command
		cmd.Function = "command"
		for _, spec := range f.Data.Functions.Command {
			cmd.Details = append(cmd.Details, binary.FunctionSpec{
				Description: spec.Description,
				Code:        spec.Code})
		}
		bin.Commands = append(bin.Commands, cmd)
	}
	if f.Data.Functions.BindShell != nil {
		var cmd binary.Command
		cmd.Function = "bindshell"
		for _, spec := range f.Data.Functions.BindShell {
			cmd.Details = append(cmd.Details, binary.FunctionSpec{
				Description: spec.Description,
				Code:        spec.Code})
		}
		bin.Commands = append(bin.Commands, cmd)
	}
	if f.Data.Functions.SUID != nil {
		var cmd binary.Command
		cmd.Function = "suid"
		for _, spec := range f.Data.Functions.SUID {
			cmd.Details = append(cmd.Details, binary.FunctionSpec{
				Description: spec.Description,
				Code:        spec.Code})
		}
		bin.Commands = append(bin.Commands, cmd)
	}
	if f.Data.Functions.LimitedSUID != nil {
		var cmd binary.Command
		cmd.Function = "limitedsuid"
		for _, spec := range f.Data.Functions.LimitedSUID {
			cmd.Details = append(cmd.Details, binary.FunctionSpec{
				Description: spec.Description,
				Code:        spec.Code})
		}
		bin.Commands = append(bin.Commands, cmd)
	}
	if f.Data.Functions.ReverseShell != nil {
		var cmd binary.Command
		cmd.Function = "revshell"
		for _, spec := range f.Data.Functions.ReverseShell {
			cmd.Details = append(cmd.Details, binary.FunctionSpec{
				Description: spec.Description,
				Code:        spec.Code})
		}
		bin.Commands = append(bin.Commands, cmd)
	}
	if f.Data.Functions.NonInteractiveBindShell != nil {
		var cmd binary.Command
		cmd.Function = "noninteractivebindshell"
		for _, spec := range f.Data.Functions.NonInteractiveBindShell {
			cmd.Details = append(cmd.Details, binary.FunctionSpec{
				Description: spec.Description,
				Code:        spec.Code})
		}
		bin.Commands = append(bin.Commands, cmd)
	}
	if f.Data.Functions.Capabilities != nil {
		var cmd binary.Command
		cmd.Function = "capabilities"
		for _, spec := range f.Data.Functions.Capabilities {
			cmd.Details = append(cmd.Details, binary.FunctionSpec{
				Description: spec.Description,
				Code:        spec.Code})
		}
		bin.Commands = append(bin.Commands, cmd)
	}
	return bin
}

func parse(filePath string) gtfobin {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.WithFields(log.Fields{
			"File":  filePath,
			"Error": err.Error(),
		}).Error("Failed to parse file")
	}
	var bin gtfobin
	err = yaml.Unmarshal(yamlFile, &bin)
	return bin
}
