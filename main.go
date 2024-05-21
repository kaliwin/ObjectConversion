package main

import (
	"github.com/kaliwin/ObjectConversion/cli"
)

func main() {

	err := cli.RootCmd.Execute()
	if err != nil {
		panic(err)
	}

	//fmt.Println(http.BodySign([]byte("cyvk")))

	//http.Diversion("/root/tmp/data", "", "/root/tmp/dir")

}
