package main

import (
	"fmt"

	client "git.randomchars.net/Reviath/RemiliaScarlet/client"
)

func main() {
	err := client.ReadConfig()

	if err != nil {
		fmt.Print(err.Error())
	}

	client.Start()

	<-make(chan struct{})
}
