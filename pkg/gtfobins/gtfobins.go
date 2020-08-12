package gtfobins

import (
	"fmt"
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

func (s *Spec) SpecPrint() {
	if len(s.Description) > 0 {
		fmt.Printf("- Description:\n")
		fmt.Printf("%v\n", s.Description)
	}
	if len(s.Code) > 0 {
		fmt.Printf("- Code:\n")
		fmt.Printf("%s\n", s.Code)
	}
	if len(s.Code) > 0 || len(s.Description) > 0 {
		fmt.Printf("\n")
	}
}

func CloneGTFO(destination string) {
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
			b := transform(f)
			parsedFiles = append(parsedFiles, b)
		}
	}
	return parsedFiles
}

func transform(f FileInfo) binary.Binary {
	var bin binary.Binary
	bin.Name = f.Binary
	bin.Path = f.Filename
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

func parse(filePath string) GTFObin {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.WithFields(log.Fields{
			"File":  filePath,
			"Error": err.Error(),
		}).Error("Failed to parse file")
	}
	var bin GTFObin
	err = yaml.Unmarshal(yamlFile, &bin)
	return bin

}

func (f *FileInfo) GTFOHasFunction(a string) bool {
	switch a {
	case "FileUpload":
		return f.Data.Functions.FileUpload != nil
	case "FileDownload":
		return f.Data.Functions.FileDownload != nil
	case "FileWrite":
		return f.Data.Functions.LibraryLoad != nil
	case "Sudo":
		return f.Data.Functions.Sudo != nil
	case "NonInteractiveReverseShell":
		return f.Data.Functions.NonInteractiveReverseShell != nil
	case "Command":
		return f.Data.Functions.Command != nil
	case "BindShell":
		return f.Data.Functions.BindShell != nil
	case "SUID":
		return f.Data.Functions.SUID != nil
	case "LimitedSUID":
		return f.Data.Functions.LimitedSUID != nil
	case "ReverseShell":
		return f.Data.Functions.ReverseShell != nil
	case "NonInteractiveBindShell":
		return f.Data.Functions.NonInteractiveBindShell != nil
	case "Capabilities":
		return f.Data.Functions.Capabilities != nil
	case "Shell":
		return f.Data.Functions.Shell != nil
	default:
		return false
	}
}
func (f *FileInfo) GTFOGetFunctionDetails(a string) []Spec {
	switch a {
	case "FileUpload":
		return f.Data.Functions.FileUpload
	case "FileDownload":
		return f.Data.Functions.FileDownload
	case "FileWrite":
		return f.Data.Functions.FileWrite
	case "LibraryLoad":
		return f.Data.Functions.LibraryLoad
	case "Sudo":
		return f.Data.Functions.Sudo
	case "NonInteractiveReverseShell":
		return f.Data.Functions.NonInteractiveReverseShell
	case "Command":
		return f.Data.Functions.Command
	case "BindShell":
		return f.Data.Functions.BindShell
	case "SUID":
		return f.Data.Functions.SUID
	case "LimitedSUID":
		return f.Data.Functions.LimitedSUID
	case "ReverseShell":
		return f.Data.Functions.ReverseShell
	case "NonInteractiveBindShell":
		return f.Data.Functions.NonInteractiveBindShell
	case "Capabilities":
		return f.Data.Functions.Capabilities
	case "Shell":
		return f.Data.Functions.Shell
	default:
		return nil
	}
}

func (f *FileInfo) GTFOPrettyPrint() {
	fmt.Printf("Information about: %v\n", f.Binary)
	if f.Data.Functions.Shell != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("Shell:\n")
		specs := f.Data.Functions.Shell
		for _, spec := range specs {
			spec.SpecPrint()
		}
	}
	if f.Data.Functions.FileUpload != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("File Upload:\n")
		specs := f.Data.Functions.FileUpload
		for _, spec := range specs {
			spec.SpecPrint()
		}
	}
	if f.Data.Functions.FileDownload != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("File Download:\n")
		specs := f.Data.Functions.FileDownload
		for _, spec := range specs {
			spec.SpecPrint()
		}
	}
	if f.Data.Functions.FileWrite != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("File Write:\n")
		specs := f.Data.Functions.FileWrite
		for _, spec := range specs {
			spec.SpecPrint()
		}
	}
	if f.Data.Functions.LibraryLoad != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("Library Load:\n")
		specs := f.Data.Functions.LibraryLoad
		for _, spec := range specs {
			spec.SpecPrint()
		}
	}
	if f.Data.Functions.Sudo != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("Sudo:\n")
		specs := f.Data.Functions.Sudo
		for _, spec := range specs {
			spec.SpecPrint()
		}
	}
	if f.Data.Functions.NonInteractiveReverseShell != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("Non Interactive Reverse Shell:\n")
		specs := f.Data.Functions.NonInteractiveReverseShell
		for _, spec := range specs {
			spec.SpecPrint()
		}
	}
	if f.Data.Functions.Command != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("Command:\n")
		specs := f.Data.Functions.Command
		for _, spec := range specs {
			spec.SpecPrint()
		}
	}
	if f.Data.Functions.BindShell != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("Bind Shell:\n")
		specs := f.Data.Functions.BindShell
		for _, spec := range specs {
			spec.SpecPrint()
		}
	}
	if f.Data.Functions.SUID != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("SUID:\n")
		specs := f.Data.Functions.SUID
		for _, spec := range specs {
			spec.SpecPrint()
		}
	}
	if f.Data.Functions.LimitedSUID != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("Limited SUID:\n")
		specs := f.Data.Functions.LimitedSUID
		for _, spec := range specs {
			spec.SpecPrint()
		}
	}
	if f.Data.Functions.ReverseShell != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("Reverse Shell:\n")
		specs := f.Data.Functions.ReverseShell
		for _, spec := range specs {
			spec.SpecPrint()
		}
	}
	if f.Data.Functions.NonInteractiveBindShell != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("Non Interactive Bind Shell:\n")
		specs := f.Data.Functions.NonInteractiveBindShell
		for _, spec := range specs {
			spec.SpecPrint()
		}
	}
	if f.Data.Functions.Capabilities != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("Capabilities:\n")
		specs := f.Data.Functions.Capabilities
		for _, spec := range specs {
			spec.SpecPrint()
		}
	}
}
