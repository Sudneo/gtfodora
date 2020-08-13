package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/sudneo/gtfodora/pkg/binary"
	"github.com/sudneo/gtfodora/pkg/gtfobins"
	"github.com/sudneo/gtfodora/pkg/lolbas"
)

// Global Variables
var unixFunctions = []string{"shell", "upload", "download", "filewrite", "fileread", "libraryload", "sudo", "noninteractiverevshell", "command", "bindshell", "suid", "limitedsuid", "revshell", "noninteractivebindshell", "capabilities"}
var winFunctions = []string{"command", "awlbypass", "ADS", "download", "copy", "encode", "decode", "credentials", "compile", "dump", "uacbypass", "reconnaissance"}

func init() {
	log.SetFormatter(&log.TextFormatter{
		PadLevelText:     true,
		DisableTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func removeDuplicateValues(s []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func listAll(b_list []binary.Binary) {
	fmt.Println(">>> Unix binaries:")
	for _, file := range b_list {
		if file.Type == "unix" {
			fmt.Println(file.Name)
		}
	}
	fmt.Println(">>> Windows binaries:")
	for _, file := range b_list {
		if file.Type == "win" {
			fmt.Println(file.Name)
		}
	}
}

func searchFunction(b_list []binary.Binary, function string) []string {
	var binaries []string
	for _, file := range b_list {
		if has, _ := file.HasFunction(function); has {
			binaries = append(binaries, file.Name)
		}
	}
	return binaries
}

func binSearch(bin string, function string, b_list []binary.Binary) bool {
	for _, file := range b_list {
		if file.Name == bin {
			if function != "" {
				if has, details := file.HasFunction(function); has {
					log.WithFields(log.Fields{
						"Function": function,
						"Binary":   file.Name,
					}).Info("The function is supported by the binary")
					fmt.Printf("[+] %v:\n", function)
					fmt.Println(strings.Repeat("-", len(function)+5))
					for _, detail := range details {
						detail.Print()
					}
					return true
				} else {
					log.WithFields(log.Fields{
						"Binary":   file.Name,
						"Function": function,
					}).Errorf("%q does not allow to perform function %q.\n", file.Name, function)
					return false
				}
			}
			file.Print()
			return true
		}
	}
	return false
}

func search(bin string, function string, b_list []binary.Binary) {
	found := binSearch(bin, function, b_list)
	if !found {
		log.WithFields(log.Fields{
			"Binary": bin,
		}).Error("No results for the specified binary")
		log.WithFields(log.Fields{
			"Details": b_list,
		}).Debug("Details")
	}
}

func listFunctions(unix bool, win bool) {
	var result []string
	if unix {
		result = append(result, unixFunctions...)
	}
	if win {
		result = append(result, winFunctions...)
	}
	result = removeDuplicateValues(result)
	fmt.Println("Functions available:")
	for _, f := range result {
		fmt.Printf("\t%v\n", f)
	}
}

func validateFunction(a string) bool {
	if !stringInSlice(a, unixFunctions) && !stringInSlice(a, winFunctions) {
		return false
	}
	return true
}

func main() {
	cloneDirPtr := flag.String("clone-path", "/tmp", "The path in which to clone the gtfobin and lolbas repos, defaults to \"/tmp.\"")
	listFunctionsPtr := flag.Bool("list-functions", false, "List the functions for the binaries")
	listAllPtr := flag.Bool("list-all", false, "List all the binaries in the collection")
	unixFilterPtr := flag.Bool("unix", false, "Filter the search among only unix binaries (i.e., gtfobin)")
	winFilterPtr := flag.Bool("win", false, "Filter the search among only windows binaries (i.e, lolbas)")
	functionPtr := flag.String("f", "", "Filter the search only for the specified function")
	searchBinPtr := flag.String("s", "", "Search for the binary specified and prints its details")
	verbosePtr := flag.Bool("v", false, "Set loglevel to DEBUG")
	flag.Parse()
	if *verbosePtr {
		log.SetLevel(log.DebugLevel)
	}
	win := *winFilterPtr || (!*unixFilterPtr && !*winFilterPtr)
	unix := *unixFilterPtr || (!*unixFilterPtr && !*winFilterPtr)
	var binary_list []binary.Binary
	if win {
		lolbas_location := fmt.Sprintf("%v/lolbas", *cloneDirPtr)
		lolbas.Clone(lolbas_location)
		binary_list = append(binary_list, lolbas.ParseAll(lolbas_location)...)
	}
	if unix {
		gtfo_location := fmt.Sprintf("%v/gtfo", *cloneDirPtr)
		gtfobins.Clone(gtfo_location)
		binary_list = append(binary_list, gtfobins.ParseAll(gtfo_location)...)
	}
	if len(binary_list) == 0 {
		log.Fatal("Error encountered getting information about binaries. Aborting")
	}
	// Processing of commandLine Args
	switch {
	case *listFunctionsPtr:
		listFunctions(unix, win)
	case *listAllPtr:
		listAll(binary_list)
	case *functionPtr != "":
		valid := validateFunction(*functionPtr)
		if !valid {
			log.WithFields(log.Fields{
				"Function":            *functionPtr,
				"Available Functions": append(unixFunctions, winFunctions...),
			}).Error("The function selected is not available.")
			return
		}
		if *searchBinPtr == "" {
			bin_list := searchFunction(binary_list, *functionPtr)
			if len(bin_list) > 0 {
				fmt.Printf("List of all the binaries with function %v:\n", *functionPtr)
				for _, s := range bin_list {
					fmt.Println(s)
				}
			} else {
				log.WithFields(log.Fields{
					"Function": *functionPtr,
				}).Error("No binary found with the specified function")
			}
		} else {
			search(*searchBinPtr, *functionPtr, binary_list)
		}
	case *searchBinPtr != "":
		search(*searchBinPtr, "", binary_list)
	default:
		log.Error("No necessary flags were specified")
	}
}
