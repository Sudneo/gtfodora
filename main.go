package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/sudneo/gtfodora/pkg/gtfobins"
	"github.com/sudneo/gtfodora/pkg/lolbas"
)

// Global Variables
var unixFunctions = []string{"Shell", "FileUpload", "FileDownload", "FileWrite", "FileRead", "LibraryLoad", "Sudo", "NonInteractiveReverseShell", "Command", "BindShell", "SUID", "LimitedSUID", "ReverseShell", "NonInteractiveBindShell", "Capabilities"}
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

func listAll(unix bool, win bool, gtfo_list []gtfobins.FileInfo, lolbas_list []lolbas.LOLbasbin) {
	if unix {
		fmt.Println(">>> Unix binaries:")
		for _, file := range gtfo_list {
			fmt.Println(file.Binary)
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

func searchFunction(unix bool, win bool, gtfo_list []gtfobins.FileInfo, lolbas_list []lolbas.LOLbasbin, function string) []string {
	var unixBinaries []string
	if unix {
		for _, file := range gtfo_list {
			if file.GTFOHasFunction(function) {
				unixBinaries = append(unixBinaries, file.Binary)
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

func unixSearch(bin string, function string, gtfo_list []gtfobins.FileInfo) bool {
	for _, file := range gtfo_list {
		if file.Binary == bin {
			if function != "" {
				if file.GTFOHasFunction(function) {
					log.WithFields(log.Fields{
						"Function": function,
						"Binary":   file.Binary,
					}).Info("The function is supported by the binary")
					details := file.GTFOGetFunctionDetails(function)
					for _, detail := range details {
						detail.SpecPrint()
					}
					return true
				} else {
					log.WithFields(log.Fields{
						"Binary":   file.Binary,
						"Function": function,
					}).Errorf("%q does not allow to perform function %q.\n", file.Binary, function)
					return false
				}
			}
			file.GTFOPrettyPrint()
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

func search(bin string, unix bool, win bool, function string, gtfo_list []gtfobins.FileInfo, lolbas_list []lolbas.LOLbasbin) {
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
		log.WithFields(log.Fields{
			"Function":            a,
			"Available Functions": append(unixFunctions, winFunctions...),
		}).Error("The function selected is not available.")
		return false
	}
	return true
}

func main() {
	// Command Line flags
	cloneDirPtr := flag.String("clone-path", "/tmp", "The path in which to clone the gtfobin and lolbas repos, defaults to \".\"")
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
	if *functionPtr != "" {
		valid := validateFunction(*functionPtr)
		if !valid {
			return
		}
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
	if *listFunctionsPtr {
		listFunctions(unix, win)
		return
	}
	if *listAllPtr {
		listAll(unix, win, gtfo_list, lolbas_list)
		return
	}
	if *functionPtr != "" {
		if *searchBinPtr == "" {
			// We just want to list the binaries that have a certain function
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
			return
		} else {
			// We want to check if a certain binary has a certain function
			search(*searchBinPtr, unix, win, *functionPtr, gtfo_list, lolbas_list)
			return
		}
	} else if *searchBinPtr != "" {
		// Just search the binary and print its information
		search(*searchBinPtr, unix, win, "", gtfo_list, lolbas_list)
		return
	}
}
