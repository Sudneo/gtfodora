package main

import (
	cloner "github.com/sudneo/gtfodora/pkg/repo_utils"
)

const (
	gtfo_repo   string = "https://github.com/GTFOBins/GTFOBins.github.io"
	lolbas_repo string = "https://github.com/LOLBAS-Project/LOLBAS"
)

func main() {
	gtfo_location := "/tmp/gtfo"
	lolbas_location := "/tmp/lolbas"
	cloner.Clone_repo(gtfo_repo, gtfo_location)
	cloner.Clone_repo(lolbas_repo, lolbas_location)
}
