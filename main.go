package main

import (
	"encoding/json"
	"fmt"

	"github.com/sudneo/gtfodora/pkg/gtfobins"
)

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func main() {
	gtfo_location := "/tmp/gtfo"
	gtfobins.Clone(gtfo_location)
	gtfoInfo := gtfobins.ParseAll(gtfo_location)
	for _, file := range gtfoInfo {
		if file.Data.Functions.SUID != nil {
			// fmt.Println(prettyPrint(file))
			fmt.Println(file.Binary)
		}

	}
}
