package main

import (
	"os"

	"github.com/MonOsmosis/osmosis/v3/app/params"
	"github.com/MonOsmosis/osmosis/v3/cmd/osmosisd/cmd"
)

func main() {
	params.SetAddressPrefixes()
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
