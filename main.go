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
	for _, file := range gtfobins {
		fmt.Printf("Unix: %v\n", file.Binary)
	}
	for _, file := range lolbas {
		fmt.Printf("Windows: %v\n", file.Name)
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

func main() {
	var functions = [...]string{"download", "upload", "execute"}
	listFunctionsPtr := flag.Bool("list-functions", false, "List the functions for the binaries")
	// unixFilterPtr := flag.Bool("unix", false, "Filter the search among only unix binaries (i.e., gtfobin)")
	// winFilterPtr := flag.Bool("win", false, "Filter the search among only windows binaries (i.e, lolbas)")
	functionPtr := flag.String("f", "", "Filter the search only for the specified function")
	listAllPtr := flag.Bool("list-all", false, "List all the binaries in the collection")
	cloneDirPtr := flag.String("clone-path", ".", "The path in which to clone the gtfobin and lolbas repos, defaults to \".\"")
	searchBinPtr := flag.String("s", "", "Search for the binary specified and prints its details")
	flag.Parse()
	if *listFunctionsPtr {
		for _, f := range functions {
			fmt.Println(f)
		}
		return
	}
	if *functionPtr != "" {
		fmt.Printf("Will search for %v\n", *functionPtr)
	}
	gtfo_location := fmt.Sprintf("%v/gtfo", *cloneDirPtr)
	gtfobins.CloneGTFO(gtfo_location)
	lolbas_location := fmt.Sprintf("%v/lolbas", *cloneDirPtr)
	lolbas.CloneLOLbas(lolbas_location)
	lolbasbin_list := lolbas.ParseAll(lolbas_location)
	gtfobin_list := gtfobins.ParseAll(gtfo_location)
	if *listAllPtr {
		listAll(gtfobin_list, lolbasbin_list)
	} else if len(*searchBinPtr) > 0 {
		search(*searchBinPtr, gtfobin_list, lolbasbin_list)
	}
}
