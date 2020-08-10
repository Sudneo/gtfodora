package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/sudneo/gtfodora/pkg/gtfobins"
	"github.com/sudneo/gtfodora/pkg/lolbas"
)

// Global Variables
var unixFunctions = []string{"Shell", "FileUpload", "FileDownload", "FileWrite", "FileRead", "LibraryLoad", "Sudo", "NonInteractiveReverseShell", "Command", "BindShell", "SUID", "LimitedSUID", "ReverseShell", "NonInteractiveBindShell", "Capabilities"}
var winFunctions = []string{"Execute", "AWL Bypass", "ADS", "Download", "Copy", "Encode", "Decode", "Credentials", "AwL bypass", "Compile", "AWL bypass", "Dump", "UAC bypass", "Reconnaissance"}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
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
		fmt.Println("\n>>> Windows binaries:")
		for _, file := range lolbas_list {
			fmt.Println(file.Name)
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
					fmt.Printf("The binary %v allows to perform function %v.\n", file.Binary, function)
					details := file.GTFOGetFunctionDetails(function)
					if len(details[0].Description) > 0 {
						fmt.Printf("Description:\n %v\n", details[0].Description)
					}
					if len(details[0].Code) > 0 {
						fmt.Printf("Code:\n %s\n", details[0].Code)
					}
					return true
				} else {
					fmt.Printf("The binary %v does not allow to perform function %v.\n", file.Binary, function)
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
					fmt.Printf("The binary %v allows to perform function %v.\n", file.Name, function)
					details := file.LOLbasGetFunctionDetails(function)
					if len(details.Description) > 0 {
						fmt.Printf("Description:\n %v\n", details.Description)
					}
					if len(details.Code) > 0 {
						fmt.Printf("Code:\n %s\n", details.Code)
					}
					return true
				} else {
					fmt.Printf("The binary %v does not allow to perform function %v.\n", file.Name, function)
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
		fmt.Printf("No results for binary %v\n", bin)
	}
}

func listFunctions(unix bool, win bool) []string {
	var result []string
	if unix {
		result = append(result, unixFunctions...)
	}
	if win {
		result = append(result, winFunctions...)
	}
	return result
}

func main() {
	cloneDirPtr := flag.String("clone-path", ".", "The path in which to clone the gtfobin and lolbas repos, defaults to \".\"")
	listFunctionsPtr := flag.Bool("list-functions", false, "List the functions for the binaries")
	listAllPtr := flag.Bool("list-all", false, "List all the binaries in the collection")
	unixFilterPtr := flag.Bool("unix", false, "Filter the search among only unix binaries (i.e., gtfobin)")
	winFilterPtr := flag.Bool("win", false, "Filter the search among only windows binaries (i.e, lolbas)")
	functionPtr := flag.String("f", "", "Filter the search only for the specified function")
	searchBinPtr := flag.String("s", "", "Search for the binary specified and prints its details")
	flag.Parse()
	gtfo_location := fmt.Sprintf("%v/gtfo", *cloneDirPtr)
	gtfobins.CloneGTFO(gtfo_location)
	lolbas_location := fmt.Sprintf("%v/lolbas", *cloneDirPtr)
	lolbas.CloneLOLbas(lolbas_location)
	lolbas_list := lolbas.ParseAll(lolbas_location)
	gtfo_list := gtfobins.ParseAll(gtfo_location)
	if len(lolbas_list) == 0 || len(gtfo_list) == 0 {
		fmt.Println("Error encountered getting information about binaries. Aborting")
		return
	}
	win := *winFilterPtr || (!*unixFilterPtr && !*winFilterPtr)
	unix := *unixFilterPtr || (!*unixFilterPtr && !*winFilterPtr)
	// Processing of commandLine Args
	if *listFunctionsPtr {
		functions := listFunctions(unix, win)
		fmt.Println("Functions available:")
		for _, f := range functions {
			fmt.Printf("\t%v\n", f)
		}
		return
	}
	if *listAllPtr {
		listAll(unix, win, gtfo_list, lolbas_list)
		return
	}
	if *functionPtr != "" {
		if !stringInSlice(*functionPtr, unixFunctions) && !stringInSlice(*functionPtr, winFunctions) {
			fmt.Println("The function selected does not exist.\nYou can check the existing functions by using the -list-functions switch.")
			return
		}
		if *searchBinPtr == "" {
			// We just want to list the binaries that have a certain function
			bin_list := searchFunction(unix, win, gtfo_list, lolbas_list, *functionPtr)
			if len(bin_list) > 0 {
				fmt.Printf("List of all the binaries with function %v:\n", *functionPtr)
				for _, s := range bin_list {
					fmt.Println(s)
				}
			} else {
				fmt.Printf("No binary found with function %v\n", *functionPtr)
			}
			return
		} else {
			// We want to check if a certain binary has a certain function
			search(*searchBinPtr, unix, win, *functionPtr, gtfo_list, lolbas_list)
			return
		}
	}
	if *searchBinPtr != "" {
		// Just search the binary and print its information
		search(*searchBinPtr, unix, win, "", gtfo_list, lolbas_list)
		return
	}
}
