package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/sudneo/gtfodora/pkg/gtfobins"
	"github.com/sudneo/gtfodora/pkg/lolbas"
)

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

func test() {

	gtfo_location := "./gtfo"
	gtfobins.CloneGTFO(gtfo_location)
	gtfoInfo := gtfobins.ParseAll(gtfo_location)
	lolbas_location := "./lolbas"
	lolbas.CloneLOLbas(lolbas_location)
	lolbasbin_list := lolbas.ParseAll(lolbas_location)
	for _, file := range gtfoInfo {
		if file.Data.Functions.FileDownload != nil {
			fmt.Println(file.Binary)
		}
	}
	for _, file := range lolbasbin_list {
		for _, cmd := range file.Commands {
			if cmd.Category == "Download" {
				fmt.Println(file.Name)
			}
		}
	}

}
func listAll(gtfobins []gtfobins.FileInfo, lolbas []lolbas.LOLbasbin) {
	fmt.Println(" >>> Unix binaries:")
	for _, file := range gtfobins {
		if len(file.Binary) > 0 {
			fmt.Printf("%v\n", file.Binary)
		}
		if file.Binary == ".dir-locals" {
			fmt.Println(file)
		}
	}
	fmt.Println("\n >>> Windows binaries:")
	for _, file := range lolbas {
		if len(file.Name) > 0 {
			fmt.Printf("%v\n", file.Name)
		}
	}
}

func search(bin string, gtfobins []gtfobins.FileInfo, lolbas []lolbas.LOLbasbin) {
	for _, file := range gtfobins {
		if file.Binary == bin {
			fmt.Println(prettyPrint(file.Data))
			return
		}
	}
	for _, file := range lolbas {
		if file.Name == bin {
			fmt.Println(prettyPrint(file))
			return
		}
	}
	fmt.Printf("No results for binary %v\n", bin)
}

func listFunctions(lolbas []lolbas.LOLbasbin) []string {
	var functions []string
	if lolbas != nil {
		for _, bin := range lolbas {
			for _, command := range bin.Commands {
				if !stringInSlice(command.Category, functions) && len(command.Category) > 0 {
					functions = append(functions, command.Category)
				}
			}
		}
	}
	var unixFunctions = [...]string{"FileUpload", "FileDownload", "FileWrite", "FileRead", "LibraryLoad", "Sudo", "NonInteractiveReverseShell", "Command", "BindShell", "SUID", "LimitedSUID", "ReverseShell", "NonInteractiveBindShell", "Capabilities"}

	for _, f := range unixFunctions {
		functions = append(functions, f)
	}
	return functions
}

func main() {
	// var functions = [...]string{"download", "upload", "execute"}
	listFunctionsPtr := flag.Bool("list-functions", false, "List the functions for the binaries")
	// unixFilterPtr := flag.Bool("unix", false, "Filter the search among only unix binaries (i.e., gtfobin)")
	// winFilterPtr := flag.Bool("win", false, "Filter the search among only windows binaries (i.e, lolbas)")
	functionPtr := flag.String("f", "", "Filter the search only for the specified function")
	listAllPtr := flag.Bool("list-all", false, "List all the binaries in the collection")
	cloneDirPtr := flag.String("clone-path", ".", "The path in which to clone the gtfobin and lolbas repos, defaults to \".\"")
	searchBinPtr := flag.String("s", "", "Search for the binary specified and prints its details")
	flag.Parse()
	if *functionPtr != "" {
		fmt.Printf("Will search for %v\n", *functionPtr)
	}
	gtfo_location := fmt.Sprintf("%v/gtfo", *cloneDirPtr)
	gtfobins.CloneGTFO(gtfo_location)
	lolbas_location := fmt.Sprintf("%v/lolbas", *cloneDirPtr)
	lolbas.CloneLOLbas(lolbas_location)
	lolbasbin_list := lolbas.ParseAll(lolbas_location)
	gtfobin_list := gtfobins.ParseAll(gtfo_location)
	if *listFunctionsPtr {
		functions := listFunctions(lolbasbin_list)
		fmt.Println("Functions available:")
		for _, f := range functions {
			fmt.Println(f)
		}
		return
	}
	if *listAllPtr {
		listAll(gtfobin_list, lolbasbin_list)
	} else if len(*searchBinPtr) > 0 {
		search(*searchBinPtr, gtfobin_list, lolbasbin_list)
	}
}
