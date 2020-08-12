package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/sudneo/gtfodora/pkg/binary"
	"github.com/sudneo/gtfodora/pkg/gtfobins"
	"github.com/sudneo/gtfodora/pkg/lolbas"
)

// Global Variables
var unixFunctions = []string{"shell", "upload", "download", "filewrite", "fileread", "libraryload", "sudo", "noninteractiverevshell", "command", "bindshell", "suid", "limitedsuid", "revshell", "noninteractivebindshell", "capabilities"}
var winFunctions = []string{"Execute", "AWL Bypass", "ADS", "Download", "Copy", "Encode", "Decode", "Credentials", "AwL bypass", "Compile", "AWL bypass", "Dump", "UAC bypass", "Reconnaissance"}

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

func listAll(unix bool, win bool, gtfo_list []binary.Binary, lolbas_list []lolbas.LOLbasbin) {
	if unix {
		fmt.Println(">>> Unix binaries:")
		for _, file := range gtfo_list {
			fmt.Println(file.Name)
		}
	}
	if win {
		fmt.Println(">>> Windows binaries:")
		for _, file := range lolbas_list {
			if file.Name != "" {
				fmt.Println(file.Name)
			}
		}
	}
}

func searchFunction(unix bool, win bool, gtfo_list []binary.Binary, lolbas_list []lolbas.LOLbasbin, function string) []string {
	var unixBinaries []string
	if unix {
		for _, file := range gtfo_list {
			if has, _ := file.HasFunction(function); has {
				unixBinaries = append(unixBinaries, file.Name)
			}
		}
	}
	var winBinaries []string
	if win {
		for _, file := range lolbas_list {
			if file.LOLbasHasFunction(function) {
				winBinaries = append(winBinaries, file.Name)
			}
		}
	}
	return append(unixBinaries, winBinaries...)
}

func unixSearch(bin string, function string, gtfo_list []binary.Binary) bool {
	for _, file := range gtfo_list {
		if file.Name == bin {
			if function != "" {
				if has, details := file.HasFunction(function); has {
					log.WithFields(log.Fields{
						"Function": function,
						"Binary":   file.Name,
					}).Info("The function is supported by the binary")
					for _, detail := range details {
						fmt.Print(detail)
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
			fmt.Print(file)
			return true
		}
	}
	return false
}

func winSearch(bin string, function string, lolbas_list []lolbas.LOLbasbin) bool {
	for _, file := range lolbas_list {
		if file.Name == bin {
			if function != "" {
				if file.LOLbasHasFunction(function) {
					log.WithFields(log.Fields{
						"Function": function,
						"Binary":   file.Name,
					}).Info("The function is supported by the binary")
					c := file.LOLbasGetFunctionDetails(function)
					c.CmdPrettyPrint()
					return true
				} else {
					log.WithFields(log.Fields{
						"Binary":   file.Name,
						"Function": function,
					}).Errorf("%q does not allow to perform function %q.\n", file.Name, function)
					return false
				}
			}
			file.LOLbasPrettyPrint()
			return true
		}
	}
	return false
}

func search(bin string, unix bool, win bool, function string, gtfo_list []binary.Binary, lolbas_list []lolbas.LOLbasbin) {
	var unixFound bool
	if unix {
		unixFound = unixSearch(bin, function, gtfo_list)
	}
	var winFound bool
	if win {
		winFound = winSearch(bin, function, lolbas_list)
	}
	if !unixFound && !winFound {
		log.WithFields(log.Fields{
			"Binary": bin,
		}).Error("No results for the specified binary")
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
	gtfo_location := fmt.Sprintf("%v/gtfo", *cloneDirPtr)
	gtfobins.CloneGTFO(gtfo_location)
	lolbas_location := fmt.Sprintf("%v/lolbas", *cloneDirPtr)
	lolbas.CloneLOLbas(lolbas_location)

	lolbas_list := lolbas.ParseAll(lolbas_location)
	gtfo_list := gtfobins.ParseAll(gtfo_location)

	if len(lolbas_list) == 0 || len(gtfo_list) == 0 {
		log.Error("Error encountered getting information about binaries. Aborting")
		return
	}
	win := *winFilterPtr || (!*unixFilterPtr && !*winFilterPtr)
	unix := *unixFilterPtr || (!*unixFilterPtr && !*winFilterPtr)
	// Processing of commandLine Args
	switch {
	case *listFunctionsPtr:
		listFunctions(unix, win)
	case *listAllPtr:
		listAll(unix, win, gtfo_list, lolbas_list)
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
			bin_list := searchFunction(unix, win, gtfo_list, lolbas_list, *functionPtr)
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
			search(*searchBinPtr, unix, win, *functionPtr, gtfo_list, lolbas_list)
		}
	case *searchBinPtr != "":
		search(*searchBinPtr, unix, win, "", gtfo_list, lolbas_list)
	default:
		log.Error("No necessary flags were specified")
	}
}
