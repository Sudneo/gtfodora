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
		if info, err := os.Stat(file); err == nil && !info.IsDir() && filepath.Base(file) != ".dir-locals.el" {
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
		spec := f.Data.Functions.Shell
		if len(spec[0].Description) > 0 {
			fmt.Printf("- Description:\n")
			fmt.Printf("%v\n", spec[0].Description)
		}
		if len(spec[0].Code) > 0 {
			fmt.Printf("- Code:\n")
			fmt.Printf("%s\n", spec[0].Code)
		}
	}
	if f.Data.Functions.FileUpload != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("File Upload:\n")
		spec := f.Data.Functions.FileUpload
		if len(spec[0].Description) > 0 {
			fmt.Printf("- Description:\n")
			fmt.Printf("%v\n", spec[0].Description)
		}
		if len(spec[0].Code) > 0 {
			fmt.Printf("- Code:\n")
			fmt.Printf("%s\n", spec[0].Code)
		}
	}
	if f.Data.Functions.FileDownload != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("File Download:\n")
		spec := f.Data.Functions.FileDownload
		if len(spec[0].Description) > 0 {
			fmt.Printf("- Description:\n")
			fmt.Printf("%v\n", spec[0].Description)
		}
		if len(spec[0].Code) > 0 {
			fmt.Printf("- Code:\n")
			fmt.Printf("%s\n", spec[0].Code)
		}
	}
	if f.Data.Functions.FileWrite != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("File Write:\n")
		spec := f.Data.Functions.FileWrite
		if len(spec[0].Description) > 0 {
			fmt.Printf("- Description:\n")
			fmt.Printf("%v\n", spec[0].Description)
		}
		if len(spec[0].Code) > 0 {
			fmt.Printf("- Code:\n")
			fmt.Printf("%s\n", spec[0].Code)
		}
	}
	if f.Data.Functions.LibraryLoad != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("Library Load:\n")
		spec := f.Data.Functions.LibraryLoad
		if len(spec[0].Description) > 0 {
			fmt.Printf("- Description:\n")
			fmt.Printf("%v\n", spec[0].Description)
		}
		if len(spec[0].Code) > 0 {
			fmt.Printf("- Code:\n")
			fmt.Printf("%s\n", spec[0].Code)
		}
	}
	if f.Data.Functions.Sudo != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("Sudo:\n")
		spec := f.Data.Functions.Sudo
		if len(spec[0].Description) > 0 {
			fmt.Printf("- Description:\n")
			fmt.Printf("%v\n", spec[0].Description)
		}
		if len(spec[0].Code) > 0 {
			fmt.Printf("- Code:\n")
			fmt.Printf("%s\n", spec[0].Code)
		}
	}
	if f.Data.Functions.NonInteractiveReverseShell != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("Non Interactive Reverse Shell:\n")
		spec := f.Data.Functions.NonInteractiveReverseShell
		if len(spec[0].Description) > 0 {
			fmt.Printf("- Description:\n")
			fmt.Printf("%v\n", spec[0].Description)
		}
		if len(spec[0].Code) > 0 {
			fmt.Printf("- Code:\n")
			fmt.Printf("%s\n", spec[0].Code)
		}
	}
	if f.Data.Functions.Command != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("Command:\n")
		spec := f.Data.Functions.Command
		if len(spec[0].Description) > 0 {
			fmt.Printf("- Description:\n")
			fmt.Printf("%v\n", spec[0].Description)
		}
		if len(spec[0].Code) > 0 {
			fmt.Printf("- Code:\n")
			fmt.Printf("%s\n", spec[0].Code)
		}
	}
	if f.Data.Functions.BindShell != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("Bind Shell:\n")
		spec := f.Data.Functions.BindShell
		if len(spec[0].Description) > 0 {
			fmt.Printf("- Description:\n")
			fmt.Printf("%v\n", spec[0].Description)
		}
		if len(spec[0].Code) > 0 {
			fmt.Printf("- Code:\n")
			fmt.Printf("%s\n", spec[0].Code)
		}
	}
	if f.Data.Functions.SUID != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("SUID:\n")
		spec := f.Data.Functions.SUID
		if len(spec[0].Description) > 0 {
			fmt.Printf("- Description:\n")
			fmt.Printf("%v\n", spec[0].Description)
		}
		if len(spec[0].Code) > 0 {
			fmt.Printf("- Code:\n")
			fmt.Printf("%s\n", spec[0].Code)
		}
	}
	if f.Data.Functions.LimitedSUID != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("Limited SUID:\n")
		spec := f.Data.Functions.LimitedSUID
		if len(spec[0].Description) > 0 {
			fmt.Printf("- Description:\n")
			fmt.Printf("%v\n", spec[0].Description)
		}
		if len(spec[0].Code) > 0 {
			fmt.Printf("- Code:\n")
			fmt.Printf("%s\n", spec[0].Code)
		}
	}
	if f.Data.Functions.ReverseShell != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("Reverse Shell:\n")
		spec := f.Data.Functions.ReverseShell
		if len(spec[0].Description) > 0 {
			fmt.Printf("- Description:\n")
			fmt.Printf("%v\n", spec[0].Description)
		}
		if len(spec[0].Code) > 0 {
			fmt.Printf("- Code:\n")
			fmt.Printf("%s\n", spec[0].Code)
		}
	}
	if f.Data.Functions.NonInteractiveBindShell != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("Non Interactive Bind Shell:\n")
		spec := f.Data.Functions.NonInteractiveBindShell
		if len(spec[0].Description) > 0 {
			fmt.Printf("- Description:\n")
			fmt.Printf("%v\n", spec[0].Description)
		}
		if len(spec[0].Code) > 0 {
			fmt.Printf("- Code:\n")
			fmt.Printf("%s\n", spec[0].Code)
		}
	}
	if f.Data.Functions.Capabilities != nil {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("Capabilities:\n")
		spec := f.Data.Functions.Capabilities
		if len(spec[0].Description) > 0 {
			fmt.Printf("- Description:\n")
			fmt.Printf("%v\n", spec[0].Description)
		}
		if len(spec[0].Code) > 0 {
			fmt.Printf("- Code:\n")
			fmt.Printf("%s\n", spec[0].Code)
		}
	}
}
