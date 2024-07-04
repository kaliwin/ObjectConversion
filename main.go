package main

import (
	"github.com/kaliwin/ObjectConversion/cli"
)

func main() {

	err := cli.RootCmd.Execute()
	if err != nil {
		panic(err)
	}

}
